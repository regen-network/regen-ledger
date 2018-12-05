package ceres.lang

import ceres.data.Failure
import ceres.data.Success

abstract class FunWrapper(override val type: FunctionType, val smtEncoder: SmtEncoder? = null) : TypedFun {
    open fun invoke(): Any? = IllegalStateException("fun doesn't take arity 0")
    open fun invoke(a: Any?): Any? = IllegalStateException("fun doesn't take arity 1")
    open fun invoke(a: Any?, b: Any?): Any? = IllegalStateException("fun doesn't take arity 2")
    open fun invoke(a: Any?, b: Any?, c: Any?): Any? = IllegalStateException("fun doesn't take arity 3")
    open fun invoke(a: Any?, b: Any?, c: Any?, d: Any?): Any? = IllegalStateException("fun doesn't take arity 4")
    open fun invoke(a: Any?, b: Any?, c: Any?, d: Any?, e: Any?): Any? =
        IllegalStateException("fun doesn't take arity 5")

    open fun invoke(vararg params: Any?): Any? = IllegalStateException("fun doesn't take arity ${params.size}")
    override fun evalChecked(params: List<TypeResult.Checked>): TypeResult {
        if (!params.all { it.hasValue })
            return TypeResult.Errors(
                listOf(
                    TypeError(
                        "Cannot evaluate wrapped function with unevaluated parameters.",
                        null
                    )
                )
            )
        val res = when (params.size) {
            0 -> invoke()
            1 -> invoke(params[0].value)
            2 -> invoke(params[0].value, params[1].value)
            3 -> invoke(params[0].value, params[1].value, params[2].value)
            4 -> invoke(params[0].value, params[1].value, params[2].value, params[3].value)
            5 -> invoke(params[0].value, params[1].value, params[2], params[3].value, params[4].value)
            else -> invoke(*params.map { it.value }.toTypedArray())
        }
        val retTc = checkFnCall(type, Env(), params)
        when (retTc) {
            is TypeResult.Checked ->
                return checked(retTc.type, res, true)
            is TypeResult.Errors -> return retTc
        }
    }

    override fun smtEncode(params: List<TypeResult.Checked>): TypeResult {
        if (smtEncoder != null) {
            val paramEncodings = params.map { it.smtEncoding }.filterNotNull()
            if (paramEncodings.size == params.size) {
                val res = smtEncoder.invoke(paramEncodings.toTypedArray())
                when(res) {
                    is Success -> {
                        val retTc = checkFnCall(type, Env(), params)
                        when (retTc) {
                            is TypeResult.Checked ->
                                return checked(retTc.type, smtEncoding = res.result)
                            is TypeResult.Errors -> return retTc
                        }
                    }
                    is Failure -> {
                        return TypeResult.error("SMT encoder returned error: ${res.error}")
                    }
                }
            }
            else return TypeResult.error("Have an SMT encoder, but parameters themselves aren't SMT encoded properly")
        } else return TypeResult.error("No SMT encoder for ${this}")
    }
}

typealias Fun0<R> = () -> R
typealias Fun1<R, A> = (A) -> R
typealias Fun2<R, A, B> = (A, B) -> R
typealias Fun3<R, A, B, C> = (A, B, C) -> R
typealias Fun4<R, A, B, C, D> = (A, B, C, D) -> R
typealias Fun5<R, A, B, C, D, E> = (A, B, C, D, E) -> R

fun <R> wrap0(ty: FunctionType, f: Fun0<R>, smtEncoder: SmtEncoder? = null) =
    checked(ty,
        object : FunWrapper(ty, smtEncoder) {
            override fun invoke(): Any? = f()
        }
    )

fun <R, A> wrap1(ty: FunctionType, f: Fun1<R, A>, smtEncoder: ((String) -> String)? = null): TypeResult.Checked {
    val enc: SmtEncoder? = if (smtEncoder == null) null else { vars ->
        if (vars.size != 1)
            Failure<String, String>("Expected 1 arg, got ${vars.size}")
        else
            Success<String, String>(smtEncoder(vars[0]))
    }
    return checked(ty,
        object : FunWrapper(ty, enc) {
            override fun invoke(a: Any?): Any? = f(a as A)
        }
    )
}

fun <R, A, B> wrap2(
    ty: FunctionType,
    f: Fun2<R, A, B>,
    smtEncoder: ((String, String) -> String)? = null
): TypeResult.Checked {
    val enc: SmtEncoder? = if (smtEncoder == null) null else { vars ->
        if (vars.size < 2)
            Failure<String, String>("Expected 2 args, got ${vars.size}")
        else
            Success<String, String>(smtEncoder(vars[0], vars[1]))
    }
    return checked(ty,
        object : FunWrapper(ty, enc) {
            override fun invoke(a: Any?, b: Any?): Any? = f(a as A, b as B)
        }
    )
}

fun <R, A, B, C> wrap3(ty: FunctionType, f: Fun3<R, A, B, C>, smtEncoder: SmtEncoder? = null) =
    checked(ty,
        object : FunWrapper(ty, smtEncoder) {
            override fun invoke(a: Any?, b: Any?, c: Any?): Any? = f(a as A, b as B, c as C)
        }
    )
