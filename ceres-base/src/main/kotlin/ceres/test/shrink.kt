package ceres.test.shrink

fun towardsLong(dest: Long, x: Long): List<Long> =
    if (dest == x) emptyList()
    else {
        val diff = x / 2L - dest / 2L
        consNub(dest, halvesLong(diff).map { x - it })
    }

fun towardsInt(dest: Int, x: Int): List<Int> =
    if (dest == x) emptyList()
    else {
        val diff = x / 2 - dest / 2
        consNub(dest, halvesInt(diff).map { x - it })
    }

fun towardsDouble(dest: Double, x: Double): List<Double> =
    if (dest == x)
        emptyList()
    else {
        val diff = x - dest
        generateSequence(diff, { it / 2.0 })
            .takeWhile({it != x && it.isNaN() &&  it.isInfinite() })
            .map { x - it }
            .toList()
    }

fun halvesLong(x: Long): List<Long> =
        generateSequence(x, {it / 2L}).takeWhile { it != 0L }.toList()

fun halvesInt(x: Int): List<Int> =
    generateSequence(x, {it / 2}).takeWhile { it != 0 }.toList()

fun <A> consNub(x: A, ys: List<A>): List<A> =
    when {
        ys.isEmpty() -> listOf(x)
        x == ys.last() -> ys
        else -> ys.plus(x)
    }

