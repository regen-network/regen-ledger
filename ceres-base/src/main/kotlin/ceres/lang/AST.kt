package ceres.lang.ast

import ceres.data.PersistentMap
import ceres.data.avlMapOf
import ceres.lang.*
import ceres.parser.HasSourceLoc
import ceres.parser.SourceLoc

//sealed class NodeBinding {
//    data class PropValue(val name: Expr.ID, val value: Expr)
//    data class NodeID(val id: Expr.ID)
//}

data class TypeCheckEnv(val bindings: PersistentMap<String, Type>) : PersistentMap<String, Type> by bindings

data class EvalEnv(val bindings: PersistentMap<String, Any?>) : PersistentMap<String, Any?> by bindings {
    fun with(vararg pairs: Pair<String, Any?>): EvalEnv = EvalEnv(bindings.setMany(pairs.asIterable()))
    fun with(pairs: Iterable<Pair<String, Any?>>): EvalEnv = EvalEnv(bindings.setMany(pairs))
}

sealed class TypeResult {
    data class Checked(val type: Type) : TypeResult()
    data class Errors(val errors: List<TypeError>) : TypeResult()

    operator fun plus(err: TypeError): TypeResult =
        when (this) {
            is Checked -> TypeResult.Errors(listOf(err))
            is Errors -> TypeResult.Errors(this.errors + err)
        }

    operator fun plus(errs: TypeResult.Errors): TypeResult =
        when (this) {
            is Checked -> errs
            is Errors -> TypeResult.Errors(this.errors + errs.errors)
        }
}

data class TypeError(val msg: String, val expr: Expr)

fun checked(type: Type): TypeResult = checked(type)

sealed class Expr : HasSourceLoc {
    abstract fun typeCheck(env: TypeCheckEnv): TypeResult
    abstract fun eval(env: EvalEnv): Any?
    protected fun error(msg: String, expr: Expr? = null): TypeResult.Errors =
        TypeResult.Errors(listOf(TypeError(msg, if(expr != null) expr else this)))
}

data class FunCall(val fn: Expr, val args: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any? {
        val fn = fn.eval(env) as AbstractFun
        val argVals = args.map { it.eval(env) }
        return fn.invoke(argVals)
    }
    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        return when (val fnChk = fn.typeCheck(env)) {
            is TypeResult.Checked ->
                when (val fnTy = fnChk.type) {
                    is FunctionType -> {
                        val params = fnTy.params
                        val nArgs = args.size
                        val nParams = params.size
                        var res =
                            if (nArgs == nParams) checked(fnTy) else error("Expected ${nParams} args but got ${nArgs}")
                        res = args.foldIndexed(res, { idx, res, arg ->
                            if (idx >= nParams) {
                                res
                            } else {
                                val param = params[idx]
                                when (val argTy = arg.typeCheck(env)) {
                                    is TypeResult.Checked -> {
                                        val sres = param.second.checkSubType(argTy.type)
                                        if (sres == null) res
                                        else res + error(sres)
                                    }
                                    is TypeResult.Errors -> res + argTy
                                }
                            }
                        })
                        return res
                    }
                    else -> error("Can't call non-function type ${fnTy}")
                }
            is TypeResult.Errors -> fnChk
        }
    }
}

data class FunExpr(
    val name: String,
    val args: List<ArgDecl>,
    val retType: Type,
    val body: Expr, override val sourceLoc: SourceLoc? = null
) : Expr() {
    override fun eval(env: EvalEnv): Any = EvalWrapper(env)

    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    inner class EvalWrapper(val env: EvalEnv) : AbstractFun {

        override fun invoke(): Any? {
            if (args.size != 0)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 0")
            return body.eval(env)
        }

        override fun invoke(a: Any?): Any? {
            if (args.size != 1)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 1")
            return body.eval(env.with(args[0].name to a))
        }

        override fun invoke(a: Any?, b: Any?): Any? {
            if (args.size != 2)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 2")
            return body.eval(
                env.with(
                    args[0].name to a,
                    args[1].name to b
                )
            )
        }

        override fun invoke(a: Any?, b: Any?, c: Any?): Any? {
            if (args.size != 3)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 3")
            return body.eval(
                env.with(
                    args[0].name to a,
                    args[1].name to b,
                    args[2].name to c
                )
            )
        }

        override fun invoke(a: Any?, b: Any?, c: Any?, d: Any?): Any? {
            if (args.size != 4)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 4")
            return body.eval(
                env.with(
                    args[0].name to a,
                    args[1].name to b,
                    args[2].name to c,
                    args[3].name to d
                )
            )
        }

        override fun invoke(a: Any?, b: Any?, c: Any?, d: Any?, e: Any?): Any? {
            if (args.size != 5)
                throw IllegalArgumentException("Expected ${args.size} arguments, got 5")
            return body.eval(
                env.with(
                    args[0].name to a,
                    args[1].name to b,
                    args[2].name to c,
                    args[3].name to d,
                    args[4].name to e
                )
            )
        }

        override fun invoke(vararg params: Any?): Any? {
            if(args.size != params.size)
                throw IllegalArgumentException("Expected ${args.size} arguments, got ${params.size}")
            return body.eval(env.with(args.mapIndexed { index, argDecl ->  argDecl.name to params[index] }))
        }
    }
}

data class ArgDecl(val name: String, val type: Type, val value: Expr?, override val sourceLoc: SourceLoc? = null) :
    HasSourceLoc

data class FunBody(val statements: List<FunStatement>)

sealed class FunStatement {
    data class LocalBinding(val name: String, val value: Expr) : FunStatement()
    data class ExprStatement(val expr: Expr) : FunStatement()
}

data class CondExpr(
    val clauses: List<Pair<Expr, Expr>>,
    val default: Expr?,
    override val sourceLoc: SourceLoc? = null
) : Expr() {
    override fun eval(env: EvalEnv): Any? {
        for(clause in clauses) {
            if(clause.first.eval(env) == true)
                return clause.second.eval(env)
        }
        if(default != null)
            return default.eval(env)
        throw IllegalStateException("Unexpected conditional expression with missing cases")
    }

    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        var ret = clauses.fold(checked(EmptyType), { ret, clause ->
            when(val ra = clause.first.typeCheck(env)) {
                is TypeResult.Checked ->
                    when(val condErr = boolType.checkSubType(ra.type)) {
                        null -> TODO()
                        else -> {
                            ret + error(condErr, clause.first)
                        }
                    }
                is TypeResult.Errors -> return ret + ra
            }
        })
        return ret
    }
}

data class TypeCaseExpr(
    val clauses: List<Pair<Expr, Expr>>,
    val default: Expr?,
    override val sourceLoc: SourceLoc? = null
) : Expr() {
    override fun eval(env: EvalEnv): Any? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class ObjExpr(val pairs: List<Pair<String, Expr>>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any? {
        val kvs: MutableList<Pair<String, Any>> = mutableListOf()
        for(kvp in pairs) {
            var v = kvp.second.eval(env)
            if(v != null) {
                kvs.add(kvp.first to v)
            }
        }
        return EntityImpl(avlMapOf<String, Any>(kvs))
    }

    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

sealed class PropertyTypeExpr

data class ObjTypeExpr(val pairs: List<Pair<String, PropertyTypeExpr>>, override val sourceLoc: SourceLoc? = null) :
    Expr() {
    override fun eval(env: EvalEnv): Any? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class VarRef(val name: String, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any? = env[name]

    override fun typeCheck(env: TypeCheckEnv): TypeResult =
        when (val ty = env[name]) {
            null -> error("Unresolved references to ${name}")
            else -> checked(ty)
        }
}

data class PropertyAccess(
    val expr: Expr,
    val prop: String, override val sourceLoc: SourceLoc? = null
) : Expr() {
    override fun eval(env: EvalEnv): Any? {
        val obj = expr.eval(env) as Entity
        return obj.get(prop)
    }

    override fun typeCheck(env: TypeCheckEnv): TypeResult =
        when (val res = expr.typeCheck(env)) {
            is TypeResult.Checked ->
                when (val ty = res.type) {
                    is EntityType ->
                        when (val prop = ty.properties[prop]) {
                            is OneProperty<*> -> checked(prop.type)
                            is ZeroOrOneProperty<*> -> checked(NullableType(prop.type))
                            is SetProperty<*> -> checked(SetType(prop.type))
                            is ListProperty<*> -> checked(ListType(prop.type))
                            null -> error("Property ${prop} not found in entity with type ${ty}")
                        }
                    is DisjointEntityUnion -> TODO()
                    else -> error("Can't access property on non-entity type ${ty}")
                }
            is TypeResult.Errors -> res
        }
}

data class OrExpr(val exprs: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun eval(env: EvalEnv): Any? {
        for(expr in exprs) {
            if(expr.eval(env) == true)
                return true
        }
        return false
    }
}

data class AndExpr(val exprs: List<Expr>, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun typeCheck(env: TypeCheckEnv): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun eval(env: EvalEnv): Any? {
        for(expr in exprs) {
            if(expr.eval(env) == false)
                return false
        }
        return true
    }
}

data class DoubleL(val x: Double, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any = x

    override fun typeCheck(env: TypeCheckEnv): TypeResult = checked(DoubleType())
}

data class IntegerL(val x: Integer, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any = x

    override fun typeCheck(env: TypeCheckEnv): TypeResult = checked(IntegerType())
}

data class StringL(val x: String, override val sourceLoc: SourceLoc? = null) : Expr() {
    override fun eval(env: EvalEnv): Any = x

    override fun typeCheck(env: TypeCheckEnv): TypeResult = checked(StringType())
}
//data class RealL(val x: Real) : Expr()
//data class ID(val x: String) : Expr()
//data class NodeE(val props: List<NodeBinding>) : Expr()
//data class GraphE(val nodes: List<NodeE>) : Expr()
//data class Free(val x: String)
//data class SelectE(val where: List<Expr>)

enum class Visibility {
    Public, Private
}

data class Decl(
    val visibility: Visibility,
    val name: String,
    val value: Expr
)

data class Module(val decls: List<Decl>)


