package ceres.lang.sexpr

import ceres.parser.HasSourceLoc
import ceres.parser.SourceLoc
import kotlin.Number

sealed class SExpr: HasSourceLoc

fun List<SExpr>.str() = map { it.toString() }.joinToString(" ")

data class Parens(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString(): String = "(${xs.str()})"
}

data class Square(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString(): String {
        println(xs)
        return "[${xs.str()}]"
    }
}

data class Curly(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString(): String = "{${xs.str()}}"
}

data class Symbol(val name: String, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString(): String = name
}

data class Keyword(val name: String, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString(): String = ":${name}"
}

data class Num(val x: Number, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = x.toString()
}

data class Str(val x: String, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = "\"" + x.map(::escapeChar).joinToString() + "\""
}

fun escapeChar(ch: Char) = when(ch) {
    '\"' -> "\\\""
    '\\' -> "\\\\"
    else -> ch.toString()
}

data class Bool(val x: Boolean, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = x.toString()
}

data class Nil(override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = "nil"
}

data class Tagged(val tag: String?, val expr: SExpr, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = if(tag == null) "#${expr}" else "#${tag} ${expr}"
}
