package ceres.util

fun <A, T> Iterable<T>.reduceAcc(acc: A, operation: (A, T) -> A): A {
    var res = acc
    this.forEach {
        res = operation(res, it)
    }
    return res
}

fun <A, T> Sequence<T>.reduceAcc(acc: A, operation: (A, T) -> A): A {
    var res = acc
    this.forEach {
        res = operation(res, it)
    }
    return res
}
