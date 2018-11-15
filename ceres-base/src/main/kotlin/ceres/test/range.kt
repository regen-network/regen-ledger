package ceres.test.range

inline class Size(val size: Int)

data class Range<A>(val origin: A, val f: (Size) -> Pair<A, A>)

fun <A> Range<A>.bounds(size: Size) = f(size)

fun <A> constantRange(lower: A, upper: A) =
        constantRangeFrom(lower, lower, upper)

fun <A> constantRangeFrom(origin: A, lower: A, upper: A) =
        Range(origin, { lower to upper })


