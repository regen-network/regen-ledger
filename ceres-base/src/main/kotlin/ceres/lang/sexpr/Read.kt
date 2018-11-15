package ceres.lang.sexpr

import ceres.lang.util.Source
import ceres.lang.util.SourceLoc

fun read(str: String, uri: String? = null): ParseResult<Char, List<SExpr>> {
    val ts = CharTokenSource(str, 0, StringSource(str, uri))
    if (str.length <= 0)
        return ts.unexpectedEof()
    return reader(ts)
}

val ws: Parser<Char, Unit> = star(testChar({
    when (it) {
        ' ' -> true
        ',' -> true
        '\t' -> true
        '\r' -> true
        '\n' -> true
        else -> false
    }
}, { "Expected whitespace, found ${it}" })).ignore()

val parens =
    cat(char('('), sepByEnd(sexpr, ws), char(')')).map { x, loc ->
        Parens(x.second, loc)
    }

val curly =
    cat(char('{'), sepByEnd(sexpr, ws), char('}')).map { x, loc ->
        Curly(x.second, loc)
    }

val square =
    cat(char('['), sepByEnd(sexpr, ws), char(']')).map { x, loc ->
        Square(x.second, loc)
    }

fun idStartChar(ch: Char): Boolean =
    ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
            || ch == '_'
            || ch == '>'
            || ch == '<'
            || ch == '='

fun idChar(ch: Char): Boolean =
    idStartChar(ch) || ch >= '0' && ch <= '9'

val identifier: Parser<Char, String> =
    cat(testChar(::idStartChar, {"${it} is not a valid identifier start character"}),
        star(testChar(::idChar, {"${it} is not a valid identifier character"}))).map {
            x, _ -> String(charArrayOf(x.first) + x.second)
        }

val symbol: Parser<Char, Symbol> = identifier.map { x, loc -> Symbol(x, loc) }

val kw: Parser<Char, Keyword> =
    cat(char(':'), identifier).map { x, loc -> Keyword(x.second, loc) }

val str: Parser<Char, Str> = TODO()

val num: Parser<Char, Num> = TODO()

val bool: Parser<Char, Bool> =
    alt(
        str("true").map({ _, loc -> Bool(true, loc) }),
        str("false").map({ _, loc -> Bool(false, loc) })
    )

val nil: Parser<Char, Nil> =
    str("nil").map({ _, loc -> Nil(loc) })

val tagged: Parser<Char, Tagged> =
    cat(char('#'), opt(identifier), sexpr).map({ x, loc ->
        Tagged(x.second, x.third, loc)
    })

val sexpr: Parser<Char, SExpr> =
    alt<Char, SExpr>(parens, curly, square, symbol, kw, str, num, bool, nil, tagged)

val reader: Parser<Char, List<SExpr>> = sepByEnd(sexpr, ws)

interface TokenSource<Token> {
    val cur: Token
    val loc: Int
    val source: Source
    fun next(): TokenSource<Token>?
}

data class CharTokenSource(val charSequence: CharSequence, override val loc: Int, override val source: Source) :
    TokenSource<Char> {
    override val cur: Char
        get() = charSequence[loc]

    override fun next(): TokenSource<Char>? {
        if (loc + 1 < charSequence.length)
            return copy(loc = loc + 1)
        return null
    }
}

data class StringSource(val text: CharSequence, override val uri: String?) : Source

data class ParseError(val error: String, val loc: SourceLoc)

sealed class ParseResult<T, out R> {
    data class Success<T, R>(val result: R, val loc: SourceLoc, val left: TokenSource<T>?) : ParseResult<T, R>()
    data class Error<T, R>(val errors: List<ParseError>, val left: TokenSource<T>? = null) : ParseResult<T, R>()
}

fun <T> TokenSource<T>.sourceLoc(left: TokenSource<T>?): SourceLoc {
    return SourceLoc(source, loc, left?.loc ?: loc)
}

fun <T, R> TokenSource<T>.parsed(x: R, left: TokenSource<T>?): ParseResult<T, R> {
    return ParseResult.Success(x, sourceLoc(left), left)
}

fun <T, R> TokenSource<T>.error(error: String, left: TokenSource<T>?): ParseResult<T, R> {
    return ParseResult.Error(listOf(ParseError(error, sourceLoc(left))), left)
}

fun <T, R> TokenSource<T>.unexpectedEof(): ParseResult<T, R> {
    return ParseResult.Error(listOf(ParseError("Unexpected EOF", sourceLoc(null))))
}

typealias Parser<Token, Result> = (TokenSource<Token>) -> ParseResult<Token, Result>

fun <T, A, B> cat(a: Parser<T, A>, b: Parser<T, B>): Parser<T, Pair<A, B>> = {
    when (val aRes = a(it)) {
        is ParseResult.Success ->
            when (val left = aRes.left) {
                null -> it.unexpectedEof<T, Pair<A, B>>()
                else -> when (val bRes = b(left)) {
                    is ParseResult.Success ->
                        it.parsed(aRes.result to bRes.result, bRes.left)
                    is ParseResult.Error -> bRes as ParseResult<T, Pair<A, B>>
                }
            }
        is ParseResult.Error -> aRes as ParseResult<T, Pair<A, B>>
    }
}

fun <T, A, B, C> cat(a: Parser<T, A>, b: Parser<T, B>, c: Parser<T, C>): Parser<T, Triple<A, B, C>> = {
    when (val aRes = a(it)) {
        is ParseResult.Success ->
            when (val aLeft = aRes.left) {
                null -> it.unexpectedEof<T, Triple<A, B, C>>()
                else -> when (val bRes = b(aLeft)) {
                    is ParseResult.Success ->
                        when (val bLeft = bRes.left) {
                            null -> aLeft.unexpectedEof<T, Triple<A, B, C>>()
                            else -> when (val cRes = c(bLeft)) {
                                is ParseResult.Success ->
                                    it.parsed(Triple(aRes.result, bRes.result, cRes.result), cRes.left)
                                is ParseResult.Error -> cRes as ParseResult<T, Triple<A, B, C>>
                            }
                        }
                    is ParseResult.Error -> bRes as ParseResult<T, Triple<A, B, C>>
                }
            }
        is ParseResult.Error -> aRes as ParseResult<T, Triple<A, B, C>>
    }
}

//data class Tuple4<A, B, C, D>(val a: A, val b: B, val c: C, val d: D)
//
//data class Tuple5<A, B, C, D, E>(val a: A, val b: B, val c: C, val d: D, val e: E)

fun <T, R> alt(vararg parsers: Parser<T, R>): Parser<T, R> = fun(ts: TokenSource<T>): ParseResult<T, R> {
    for (parser in parsers) {
        val res = parser(ts)
        when (res) {
            is ParseResult.Success -> return res
            else -> {
            }
        }
    }
    return ts.error("Expected one of the alternatives", ts)
}

fun <T, R> opt(parser: Parser<T, R>): Parser<T, R?> = {
    when (val res = parser(it)) {
        is ParseResult.Success -> res
        is ParseResult.Error -> it.parsed(null, it)
    }
}

fun <T, R> star(parser: Parser<T, R>): Parser<T, List<R>> = {
    val res = mutableListOf<R>()
    var curTs: TokenSource<T>? = it
    loop@ while (curTs != null) {
        val x = parser(curTs)
        when (x) {
            is ParseResult.Success -> {
                res.add(x.result)
                curTs = x.left
            }
            is ParseResult.Error -> break@loop
        }
    }
    it.parsed(res, curTs)
}

fun <T, R> plus(parser: Parser<T, R>): Parser<T, List<R>> {
    val p = star(parser)
    return {
        when (val res = p(it)) {
            is ParseResult.Success ->
                if(res.result.size == 0)
                    it.error("Expected non-empty", it)
                else res
            is ParseResult.Error -> res
        }
    }
}

fun <T, R> sepBy(parser: Parser<T, R>, sep: Parser<T, *>): Parser<T, List<R>> =
    star(cat(sep, parser).map({ x, _ -> x.second }))

fun <T, R> sepByEnd(parser: Parser<T, R>, sep: Parser<T, *>): Parser<T, List<R>> =
    cat(sepBy(parser, sep), sep).map({ x, _ -> x.first })

fun <T, A, B> Parser<T, A>.map(f: (A, SourceLoc) -> B): Parser<T, B> = {
    when (val res = this(it)) {
        is ParseResult.Success -> it.parsed(f(res.result, res.loc), res.left)
        is ParseResult.Error -> ParseResult.Error(res.errors)
    }
}

fun <T, A> Parser<T, A>.ignore(): Parser<T, Unit> = map { _, _ -> Unit }

fun testChar(f: (Char) -> Boolean, err: (Char) -> String): Parser<Char, Char> = {
    val cur = it.cur
    if (f(cur)) {
        it.parsed(cur, it.next())
    } else {
        it.error(err(cur), it)
    }
}

fun char(x: Char): Parser<Char, Char> = testChar({ it == x }, { "Expected ${x}, got ${it}" })

fun str(x: String): Parser<Char, String> = fun(it: TokenSource<Char>): ParseResult<Char, String> {
    var curTs: TokenSource<Char>? = it
    for (ch in x) {
        if (curTs == null) return it.unexpectedEof()
        else if (curTs.cur == ch) curTs = curTs.next()
        else return it.error("Expected ${x}", curTs)
    }
    return it.parsed(x, curTs)
}
