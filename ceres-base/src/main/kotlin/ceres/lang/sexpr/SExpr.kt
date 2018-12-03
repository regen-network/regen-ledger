package ceres.lang.sexpr

import ceres.parser.HasSourceLoc
import ceres.parser.SourceLoc
import kotlin.Number

sealed class SExpr: HasSourceLoc

fun List<SExpr>.str() = map { it.toString() }.joinToString(" ")

abstract class HasElements: SExpr() {
    abstract val elements: List<SExpr>
}

data class Parens(override val elements: List<SExpr>, override val sourceLoc: SourceLoc? = null): HasElements() {
    override fun toString(): String = "(${elements.str()})"
}

data class Square(override val elements: List<SExpr>, override val sourceLoc: SourceLoc? = null): HasElements() {
    override fun toString(): String {
        println(elements)
        return "[${elements.str()}]"
    }
}

data class Curly(override val elements: List<SExpr>, override val sourceLoc: SourceLoc? = null): HasElements() {
    override fun toString(): String = "{${elements.str()}}"
}

abstract class Named: SExpr() {
    abstract val name: String
}

data class Symbol(override val name: String, override val sourceLoc: SourceLoc? = null): Named() {
    override fun toString(): String = name
}

data class Keyword(override val name: String, override val sourceLoc: SourceLoc? = null): Named() {
    override fun toString(): String = ":${name}"
}

data class Str(override val name: String, override val sourceLoc: SourceLoc? = null): Named() {
    override fun toString() = "\"" + name.map(::escapeChar).joinToString() + "\""
}

data class Num(val x: String, override val sourceLoc: SourceLoc? = null): SExpr() {
    override fun toString() = x.toString()
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
