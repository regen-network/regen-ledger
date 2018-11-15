package ceres.test.gen

import ceres.test.*
import ceres.test.range.*
import ceres.test.seed.*
import ceres.test.shrink.*
import kotlin.random.*

inline class Gen<A>(val runGen: (Size, Seed) -> Tree<A>)

fun <A> generate(f: (Size, Seed) -> A): Gen<A> =
    Gen({ size, seed -> treeOf(f(size, seed)) })

fun <A, B> Gen<A>.mapGenT(f: (Tree<A>) -> Tree<B>): Gen<B> =
    Gen({ size, seed ->
        f(runGen(size, seed))
    })

fun <A, B> Gen<A>.map(f: (A) -> B): Gen<B> =
        mapGenT { it.map(f) }

fun <A, B> Gen<A>.bind(f: (A) -> Gen<B>): Gen<B> =
        Gen({ size, seed ->
            val (sk, sm) = seed.split()
//            Tree({
//                val treeA = runGen(size, sm)
//                val nodeA = treeA.f()
//                val genB = f(nodeA.value)
//                nodeA.children.map {
//
//                }
//                // TODO correct mapping
//                genB.runGen(size, sk)
//                TODO()
//            })
            TODO()
        })

fun <A> Gen<A>.shrink(f: (A) -> List<A>): Gen<A> =
    mapGenT({
        it.expand(f)
    })

fun long(range: Range<Long>): Gen<Long> =
    long_(range).shrink { towardsLong(range.origin, it) }

fun long_(range: Range<Long>): Gen<Long> =
    generate { size, seed ->
        val (x, y) = range.bounds(size)
        seed.random.nextLong(x, y)
    }

fun int(range: Range<Int>): Gen<Int> =
    int_(range).shrink { towardsInt(range.origin, it) }

fun int_(range: Range<Int>): Gen<Int> =
    generate { size, seed ->
        val (x, y) = range.bounds(size)
        seed.random.nextInt(x, y)
    }

fun double(range: Range<Double>) =
    double_(range).shrink { towardsDouble(range.origin, it) }

fun double_(range: Range<Double>) =
    generate { size, seed ->
        val (x, y) = range.bounds(size)
        seed.random.nextDouble(x, y)
    }


// Combinators

fun <A> constant(a: A): Gen<A> = generate({ _, _ -> a })

fun <A> element(xs: List<A>): Gen<A> =
    if (xs.isEmpty())
        throw IllegalArgumentException("element used with empty list")
    else
        int(constantRange(0, xs.size)).map { xs[it] }

fun <A> choice(xs: List<Gen<A>>): Gen<A> =
    if (xs.isEmpty())
        throw IllegalArgumentException("element used with empty list")
    else
        int(constantRange(0, xs.size)).bind { xs[it] }

fun <A> Gen<A>.filter(pred: (A) -> Boolean): Gen<A> =
        Gen({ size, seed ->
            val tree = runGen(size, seed)
            TODO()
        })
