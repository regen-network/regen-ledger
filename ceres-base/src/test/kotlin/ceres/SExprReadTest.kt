package ceres.lang.sexpr

import ceres.lang.sexpr.analyze.analyze
import ceres.lang.sexpr.analyze.analyzeExpr
import ceres.lang.sexpr.analyze.exprParser
import ceres.lang.sexpr.read.read
import ceres.parser.char.parseString
import kotlin.test.Test

class SExprReadTest {
    @Test
    fun testRead() {
        println(read("true [true, nil] false (a b nil { x y z} #{:a b :c })"))
    }

    @Test
    fun testAnalyze() {
        println(parseString(exprParser, "(+ x y)"))
    }
}

