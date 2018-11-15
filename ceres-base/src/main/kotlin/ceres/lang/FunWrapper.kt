package ceres.lang

interface AbstractFun {
    fun invoke(): Any? = NotImplementedError()
    fun invoke(a: Any?): Any? = NotImplementedError()
    fun invoke(a: Any?, b: Any?): Any? = NotImplementedError()
    fun invoke(a: Any?, b: Any?, c: Any?): Any? = NotImplementedError()
    fun invoke(a: Any?, b: Any?, c: Any?, d: Any?): Any? = NotImplementedError()
    fun invoke(a: Any?, b: Any?, c: Any?, d: Any?, e: Any?): Any? = NotImplementedError()
    fun invoke(vararg params: Any?): Any? = NotImplementedError()
}

fun AbstractFun.invoke(params: List<Any?>): Any? =
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
