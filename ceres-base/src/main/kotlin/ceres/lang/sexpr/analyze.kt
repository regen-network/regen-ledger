package ceres.lang.sexpr.analyze

import ceres.lang.ast.*
import ceres.lang.sexpr.*
import ceres.parser.*
import ceres.parser.char.CharTokenSource
import ceres.parser.char.str

fun analyze(exprs: List<SExpr>): ParseResult<SExpr, Module> {
    TODO()
}

data class SExprSource(val tokens: List<SExpr>, val source: Source) : Source {
    override val uri: String?
        get() = source.uri

    val size: Int by lazy { tokens.size }
}

data class SExprTokenSource(override val loc: Int, override val source: SExprSource) : TokenSource<SExpr> {
    override val cur: SExpr? by lazy {
        if (loc < source.size) source.tokens[loc] else null
    }

    override fun next(): TokenSource<SExpr> {
        if (cur != null)
            return copy(loc = loc + 1)
        return this
    }
}

fun <R> named(parser: Parser<Char, R>): Parser<SExpr, R> = TODO()

val sym = testToken<SExpr>({it is Symbol}, {"Expected a symbol"})

fun <R> sym(parser: Parser<Char, R>): Parser<SExpr, R> =
        sym.bind { x, loc -> parser.parse(CharTokenSource((x as Symbol).name, 0, loc.source)) }

fun <R> hasElems(test: (SExpr) -> Boolean, err: (SExpr) -> String, parser: Parser<SExpr, R>): Parser<SExpr, R> =
    testToken<SExpr>({ it is Square }, { "Expected square brackets" }).bind { x, loc ->
        parser.parse(SExprTokenSource(0, SExprSource((x as HasElements).elements, loc.source)))
    }

fun <R> parens(parser: Parser<SExpr, R>): Parser<SExpr, R> =
    hasElems({ it is Parens }, { "Expected parentheses" }, parser)

fun <R> square(parser: Parser<SExpr, R>): Parser<SExpr, R> =
    hasElems({ it is Square }, { "Expected square brackets" }, parser)

fun <R> curly(parser: Parser<SExpr, R>): Parser<SExpr, R> =
    hasElems({ it is Curly }, { "Expected curly braces" }, parser)

val expr: Parser<SExpr, Expr> = TODO()

val exprs: Parser<SExpr, List<Expr>> = star(expr)

val funCall: Parser<SExpr, FunCall> =
    parens(cat(expr, exprs)).map({ x, loc -> FunCall(x.first, x.second, loc) })

val varRef = sym.map {x, loc -> VarRef((x as Symbol).name, loc)}

val strL = testToken<SExpr>({it is Str},{"Expected a string"})
    .map {x, loc -> StringL((x as Str).name, loc)}

val orExpr = parens(cat(sym(str("or")), star(expr))).map {x, loc -> OrExpr(x.second, loc)}

val andExpr = parens(cat(sym(str("and")), star(expr))).map {x, loc -> AndExpr(x.second, loc)}

//val fnArgs: Parser<SExpr, FunCall> =



