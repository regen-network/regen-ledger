package ceres.lang

interface AbstractFun {
    fun invoke(): Any? = IllegalStateException("fun doesn't take arity 0")
    fun invoke(a: Any?): Any? = IllegalStateException("fun doesn't take arity 1")
    fun invoke(a: Any?, b: Any?): Any? = IllegalStateException("fun doesn't take arity 2")
    fun invoke(a: Any?, b: Any?, c: Any?): Any? = IllegalStateException("fun doesn't take arity 3")
    fun invoke(a: Any?, b: Any?, c: Any?, d: Any?): Any? = IllegalStateException("fun doesn't take arity 4")
    fun invoke(a: Any?, b: Any?, c: Any?, d: Any?, e: Any?): Any? = IllegalStateException("fun doesn't take arity 5")
    fun invoke(vararg params: Any?): Any? = IllegalStateException("fun doesn't take arity ${params.size}")
}

fun AbstractFun.call(params: List<Any?>): Any? =
    when(params.size) {
        0 -> invoke()
        1 -> invoke(params[0])
        2 -> invoke(params[0], params[1])
        3 -> invoke(params[0], params[1], params[2])
        4 -> invoke(params[0], params[1], params[2], params[3])
        5 -> invoke(params[0], params[1], params[2], params[3], params[4])
        else -> invoke(*params.toTypedArray())
    }

typealias Fun0<R> = () -> R
typealias Fun1<R, A> = (A) -> R
typealias Fun2<R, A, B> = (A, B) -> R
typealias Fun3<R, A, B, C> = (A, B, C) -> R
typealias Fun4<R, A, B, C, D> = (A, B, C, D) -> R
typealias Fun5<R, A, B, C, D, E> = (A, B, C, D, E) -> R

fun <R> wrap(f: Fun0<R>): AbstractFun =
        object : AbstractFun {
            override fun invoke(): Any? = f()
        }

fun <R, A> wrap(f: Fun1<R, A>): AbstractFun =
    object : AbstractFun {
        override fun invoke(a: Any?): Any? = f(a as A)
    }

fun <R, A, B> wrap(f: Fun2<R, A, B>): AbstractFun =
    object : AbstractFun {
        override fun invoke(a: Any?, b: Any?): Any? = f(a as A, b as B)
    }

fun <R, A, B, C> wrap(f: Fun3<R, A, B, C>): AbstractFun =
    object : AbstractFun {
        override fun invoke(a: Any?, b: Any?, c: Any?): Any? = f(a as A, b as B, c as C)
    }
