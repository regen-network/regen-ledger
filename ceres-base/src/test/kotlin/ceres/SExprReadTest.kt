package ceres.lang.sexpr

import ceres.parser.char.parseString
import kotlin.test.Test

class SExprReadTest {
    @Test
    fun test() {
        println(parseString(::reader, "true [true, nil] false (a b nil { x y z} #{:a b :c })"))
    }
}

