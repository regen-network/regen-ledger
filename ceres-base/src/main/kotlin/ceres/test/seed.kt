package ceres.test.seed

import kotlin.random.Random

@ExperimentalUnsignedTypes
data class Seed(val value: ULong, val gamma: ULong) {
    val random = Random(value.toLong())
}

@ExperimentalUnsignedTypes
fun Seed.next(): Seed = Seed(value + gamma, gamma)

@ExperimentalUnsignedTypes
fun Seed.split(): Pair<Seed, Seed> {
    val s1 = next()
    val s2= s1.next()
    return s2 to Seed(mix64(s1.value), mixGamma(s2.value))
}

@ExperimentalUnsignedTypes
fun mix64(x: ULong): ULong {
    val y = (x xor (x shr  33)) * 0xff51afd7ed558ccdUL
    val z = (y xor (y shr 33)) * 0xc4ceb9fe1a85ec53UL
    return z xor (z shr 33)
}

@ExperimentalUnsignedTypes
fun mix64variant13(x: ULong): ULong {
    val y = (x xor (x shr  30)) * 0xc4ceb9fe1a85ec53UL
    val z = (y xor (y shr 27)) * 0xc4ceb9fe1a85ec53UL
    return z xor (z shr 31)
}


@ExperimentalUnsignedTypes
fun mixGamma(x: ULong): ULong {
    val y = mix64variant13(x) or 1UL
    val n = popCount(y xor (y shr 1))
    return if(n < 24) y xor 0xaaaaaaaaaaaaaaaaUL
    else y
}

@ExperimentalUnsignedTypes
fun popCount(uLong: ULong): Int {
    TODO("number of set bits count")

}
