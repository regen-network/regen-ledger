package ceres.lang.sexpr

import ceres.lang.sexpr.read.read
import ceres.parser.char.parseString
import kotlin.test.Test

class SExprReadTest {
    @Test
    fun test() {
        println(read("true [true, nil] false (a b nil { x y z} #{:a b :c })"))
    }
}

