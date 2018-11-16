package ceres.lang

import ceres.lang.ast.*


abstract class FunWrapper(override val type: FunctionType): TypedFun {
    open fun invoke(): Any? = IllegalStateException("fun doesn't take arity 0")
    open fun invoke(a: Any?): Any? = IllegalStateException("fun doesn't take arity 1")
    open fun invoke(a: Any?, b: Any?): Any? = IllegalStateException("fun doesn't take arity 2")
    open fun invoke(a: Any?, b: Any?, c: Any?): Any? = IllegalStateException("fun doesn't take arity 3")
    open fun invoke(a: Any?, b: Any?, c: Any?, d: Any?): Any? = IllegalStateException("fun doesn't take arity 4")
    open fun invoke(a: Any?, b: Any?, c: Any?, d: Any?, e: Any?): Any? = IllegalStateException("fun doesn't take arity 5")
    open fun invoke(vararg params: Any?): Any? = IllegalStateException("fun doesn't take arity ${params.size}")
    override fun evalChecked(params: List<TypeResult.Checked>): TypeResult {
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
        when(retTc) {
            is TypeResult.Checked -> return checked(retTc.type, res, true)
            is TypeResult.Errors -> return retTc
        }
    }
}

typealias Fun0<R> = () -> R
typealias Fun1<R, A> = (A) -> R
typealias Fun2<R, A, B> = (A, B) -> R
typealias Fun3<R, A, B, C> = (A, B, C) -> R
typealias Fun4<R, A, B, C, D> = (A, B, C, D) -> R
typealias Fun5<R, A, B, C, D, E> = (A, B, C, D, E) -> R

fun <R> wrap(ty: FunctionType, f: Fun0<R>) =
    checked(ty,
        object : FunWrapper(ty) {
            override fun invoke(): Any? = f()
        }
    )

fun <R, A> wrap(ty: FunctionType, f: Fun1<R, A>) =
    checked(ty,
        object : FunWrapper(ty) {
            override fun invoke(a: Any?): Any? = f(a as A)
        }
    )

fun <R, A, B> wrap(ty: FunctionType, f: Fun2<R, A, B>) =
    checked(ty,
        object : FunWrapper(ty) {
            override fun invoke(a: Any?, b: Any?): Any? = f(a as A, b as B)
        }
    )

fun <R, A, B, C> wrap(ty: FunctionType, f: Fun3<R, A, B, C>) =
    checked(ty,
        object : FunWrapper(ty) {
            override fun invoke(a: Any?, b: Any?, c: Any?): Any? = f(a as A, b as B, c as C)
        }
    )
