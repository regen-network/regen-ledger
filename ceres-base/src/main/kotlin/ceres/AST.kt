package ceres.lang.ast

import ceres.lang.*

//sealed class NodeBinding {
//    data class PropValue(val name: Expr.ID, val value: Expr)
//    data class NodeID(val id: Expr.ID)
//}

data class Env(val bindings: Map<String, Type>): Map<String, Type> by bindings

sealed class TypeResult {
    data class Checked(val type: Type): TypeResult()
    data class Errors(val errors: List<TypeError>): TypeResult()

    operator fun plus(err: TypeError): TypeResult =
            when(this) {
                is Checked -> TypeResult.Errors(listOf(err))
                is Errors -> TypeResult.Errors(this.errors + err)
            }

    operator fun plus(errs: TypeResult.Errors): TypeResult =
            when(this) {
                is Checked -> errs
                is Errors -> TypeResult.Errors(this.errors + errs.errors)
            }
}

data class TypeError(val msg: String, val expr: Expr)

fun checked(type: Type): TypeResult = checked(type)

data class SourceLoc(val filename: String, val startLoc: Pair<Int, Int>, val endLoc: Pair<Int, Int>)

sealed class Expr {
    abstract val sourceLoc: SourceLoc?
    abstract fun typeCheck(env: Env): TypeResult
    protected fun error(msg: String): TypeResult.Errors =
        TypeResult.Errors(listOf(TypeError(msg, this)))
}

data class FunCall(val fn: Expr, val args: List<Expr>, override val sourceLoc: SourceLoc?) : Expr() {
    override fun typeCheck(env: Env): TypeResult {
        return when(val fnChk = fn.typeCheck(env)) {
            is TypeResult.Checked ->
                when(val fnTy = fnChk.type) {
                    is FunctionType -> {
                        val params = fnTy.params
                        val nArgs = args.size
                        val nParams = params.size
                        var res = if(nArgs == nParams) checked(fnTy) else error("Expected ${nParams} args but got ${nArgs}")
                        res = args.foldIndexed(res, { idx, res, arg ->
                            if(idx >= nParams) {
                                res
                            } else {
                                val param = params[idx]
                                when(val argTy = arg.typeCheck(env)) {
                                    is TypeResult.Checked -> {
                                        val sres = param.second.checkSubType(argTy.type)
                                        if(sres == null) res
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

data class Fun(
    val name: String,
    val args: List<ArgDecl>,
    val retType: Type,
    val body: FunBody, override val sourceLoc: SourceLoc?
) : Expr() {
    override fun typeCheck(env: Env): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class ArgDecl(val name: String, val type: Type, val value: Expr?)

data class FunBody(val statements: List<FunStatement>)

sealed class FunStatement {
    data class LocalBinding(val name: String, val value: Expr) : FunStatement()
    data class ExprStatement(val expr: Expr) : FunStatement()
}

data class Case(val clauses: List<Pair<Expr, Expr>>, val default: Expr?, override val sourceLoc: SourceLoc?): Expr() {
    override fun typeCheck(env: Env): TypeResult {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}


data class VarRef(val name: String, override val sourceLoc: SourceLoc?): Expr() {
    override fun typeCheck(env: Env): TypeResult =
            when(val ty = env[name]) {
                null -> error("Unresolved references to ${name}")
                else -> checked(ty)
            }
}

data class PropertyAccess(
    val expr: Expr,
    val prop: String, override val sourceLoc: SourceLoc?
): Expr() {
    override fun typeCheck(env: Env): TypeResult =
        when(val res = expr.typeCheck(env)) {
            is TypeResult.Checked ->
                when(val ty = res.type) {
                    is EntityType ->
                        when(val prop = ty.properties[prop]) {
                            is OneProperty<*> -> checked(prop.type)
                            is ZeroOrOneProperty<*> -> checked(NullableType(prop.type))
                            is SetProperty<*> -> checked(SetType(prop.type))
                            is ListProperty<*> -> checked(ListType(prop.type))
                            null -> TODO()
                        }
                    is DisjointEntityUnion -> TODO()
                    else -> {
                        TODO()
                    }
                }
            is TypeResult.Errors -> res
        }
}

data class DoubleL(val x: Double, override val sourceLoc: SourceLoc?) : Expr() {
    override fun typeCheck(env: Env): TypeResult = checked(DoubleType())
}

data class IntegerL(val x: Integer, override val sourceLoc: SourceLoc?) : Expr() {
    override fun typeCheck(env: Env): TypeResult = checked(IntegerType())
}

data class StringL(val x: String, override val sourceLoc: SourceLoc?) : Expr() {
    override fun typeCheck(env: Env): TypeResult = checked(StringType())
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


