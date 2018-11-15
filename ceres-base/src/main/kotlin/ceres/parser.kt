package ceres.parser

interface Source {
    val uri: String?
}

data class SourceLoc(val source: Source, val start: Int, val end: Int)

interface HasSourceLoc {
    val sourceLoc: SourceLoc?
}

interface TokenSource<Token> {
    val cur: Token?
    val loc: Int
    val source: Source
    fun next(): TokenSource<Token>
}

data class ParseError(val error: String, val loc: SourceLoc)

sealed class ParseResult<T, out R> {
    data class Success<T, R>(val result: R, val loc: SourceLoc, val left: TokenSource<T>) : ParseResult<T, R>()
    data class Error<T, R>(val errors: List<ParseError>, val left: TokenSource<T>? = null) : ParseResult<T, R>()
}

fun <T> TokenSource<T>.sourceLoc(left: TokenSource<T>?): SourceLoc {
    return SourceLoc(source, loc, left?.loc ?: loc)
}

fun <T, R> TokenSource<T>.parsed(x: R, left: TokenSource<T>): ParseResult<T, R> {
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
            when (val bRes = b(aRes.left)) {
                is ParseResult.Success ->
                    it.parsed(aRes.result to bRes.result, bRes.left)
                is ParseResult.Error -> bRes as ParseResult<T, Pair<A, B>>
            }
        is ParseResult.Error -> aRes as ParseResult<T, Pair<A, B>>
    }
}

fun <T, A, B, C> cat(a: Parser<T, A>, b: Parser<T, B>, c: Parser<T, C>): Parser<T, Triple<A, B, C>> =
    cat(cat(a, b), c).map { x, _ ->
        Triple(x.first.first, x.first.second, x.second)
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
    var curTs: TokenSource<T> = it
    loop@ while (true) {
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
                if (res.result.size == 0)
                    it.error("Expected non-empty", it)
                else res
            is ParseResult.Error -> res
        }
    }
}

//fun <T, R> sepBy(parser: Parser<T, R>, sep: Parser<T, *>): Parser<T, List<R>> =
//    cat(parser, star(cat(sep, parser).map { x, _ -> x.second })).map { x, _ ->
//        listOf(x.first) + x.second
//    }
//
//fun <T, R> sepByEnd(parser: Parser<T, R>, sep: Parser<T, *>): Parser<T, List<R>> =
//    cat(sepBy(parser, sep), opt(sep)).map({ x, _ -> x.first })
//
//fun <T, R> sepByBefEnd(parser: Parser<T, R>, sep: Parser<T, *>): Parser<T, List<R>> =
//    cat(opt(sep), sepByEnd(parser, sep)).map { x, _ -> x.second }

fun <T, A, B> Parser<T, A>.map(f: (A, SourceLoc) -> B): Parser<T, B> = {
    when (val res = this(it)) {
        is ParseResult.Success -> it.parsed(f(res.result, res.loc), res.left)
        is ParseResult.Error -> ParseResult.Error(res.errors)
    }
}

fun <T, A> Parser<T, A>.ignore(): Parser<T, Unit> = map { _, _ -> Unit }


