package ceres.lang.sexpr

import ceres.lang.util.HasSourceLoc
import ceres.lang.util.SourceLoc
import kotlin.Number

sealed class SExpr: HasSourceLoc

data class Parens(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr()

data class Square(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr()

data class Curly(val xs: List<SExpr>, override val sourceLoc: SourceLoc? = null): SExpr()

data class Symbol(val name: String, override val sourceLoc: SourceLoc? = null): SExpr()

data class Keyword(val name: String, override val sourceLoc: SourceLoc? = null): SExpr()

data class Num(val x: Number, override val sourceLoc: SourceLoc? = null): SExpr()

data class Str(val x: String, override val sourceLoc: SourceLoc? = null): SExpr()

data class Bool(val x: Boolean, override val sourceLoc: SourceLoc? = null): SExpr()

data class Nil(override val sourceLoc: SourceLoc? = null): SExpr()

data class Tagged(val tag: String?, val expr: SExpr, override val sourceLoc: SourceLoc? = null): SExpr()

