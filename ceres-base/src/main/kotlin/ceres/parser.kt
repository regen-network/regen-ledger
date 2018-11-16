package ceres.parser

import ceres.test.gen.Gen
import ceres.test.gen.HasGen
import ceres.test.gen.choice

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
    //TODO custom error messages
}

fun <T> TokenSource<T>.sourceLoc(left: TokenSource<T>?): SourceLoc {
    // TODO allow the token source to override this (ex. SExpr tokens should still point to the original string source if possible, or to a code source if not)
    // would also be great if the token source could point to both the relevant parsed expression
    // and preceding whitespace/comments for re-formatting and docstring extraction
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

interface Parser<Token, out Result>: HasGen<List<Token>> {
    fun parse(tokens: TokenSource<Token>): ParseResult<Token, Result>
    // TODO unparse(result: Result): List<Token>
}

class cat2<T, A, B>(val a: Parser<T, A>, val b: Parser<T, B>) : Parser<T, Pair<A, B>> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(ts: TokenSource<T>): ParseResult<T, Pair<A, B>> =
        when (val aRes = a.parse(ts)) {
            is ParseResult.Success ->
                when (val bRes = b.parse(aRes.left)) {
                    is ParseResult.Success ->
                        ts.parsed(aRes.result to bRes.result, bRes.left)
                    is ParseResult.Error -> bRes as ParseResult<T, Pair<A, B>>
                }
            is ParseResult.Error -> aRes as ParseResult<T, Pair<A, B>>
        }
}

fun <T, A, B> cat(a: Parser<T, A>, b: Parser<T, B>): Parser<T, Pair<A, B>> = cat2(a, b)

fun <T, A, B, C> cat(a: Parser<T, A>, b: Parser<T, B>, c: Parser<T, C>): Parser<T, Triple<A, B, C>> =
    cat(cat(a, b), c).map { x, _ ->
        Triple(x.first.first, x.first.second, x.second)
    }

//data class Tuple4<A, B, C, D>(val a: A, val b: B, val c: C, val d: D)
//
//data class Tuple5<A, B, C, D, E>(val a: A, val b: B, val c: C, val d: D, val e: E)

class alt<T, R>(vararg val parsers: Parser<T, R>) : Parser<T, R> {
    override val gen: Gen<List<T>> by lazy {
        choice(parsers.toList())
    }

    override fun parse(ts: TokenSource<T>): ParseResult<T, R> {
        for (parser in parsers) {
            val res = parser.parse(ts)
            when (res) {
                is ParseResult.Success -> return res
                else -> {
                }
            }
        }
        return ts.error("Expected one of the alternatives", ts)
    }
}

class opt<T, R>(val parser: Parser<T, R>) : Parser<T, R?> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<T>): ParseResult<T, R?> =
        when (val res = parser.parse(tokens)) {
            is ParseResult.Success -> res
            is ParseResult.Error -> tokens.parsed(null, tokens)
        }
}

class star<T, R>(val parser: Parser<T, R>) : Parser<T, List<R>> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<T>): ParseResult<T, List<R>> {
        val res = mutableListOf<R>()
        var curTs: TokenSource<T> = tokens
        loop@ while (true) {
            val x = parser.parse(curTs)
            when (x) {
                is ParseResult.Success -> {
                    res.add(x.result)
                    curTs = x.left
                }
                is ParseResult.Error -> break@loop
            }
        }
        return tokens.parsed(res, curTs)
    }
}

class plus<T, R>(parser: Parser<T, R>) : Parser<T, List<R>> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    val pstar = star(parser)
    override fun parse(tokens: TokenSource<T>): ParseResult<T, List<R>> {
        when (val res = pstar.parse(tokens)) {
            is ParseResult.Success ->
                if (res.result.size == 0)
                    return tokens.error("Expected non-empty", tokens)
                else return res
            is ParseResult.Error -> return res
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

class Mapper<T, A, B>(val parser: Parser<T, A>, val f: (A, SourceLoc) -> B): Parser<T, B> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<T>): ParseResult<T, B> =
    when (val res = parser.parse(tokens)) {
        is ParseResult.Success -> tokens.parsed(f(res.result, res.loc), res.left)
        is ParseResult.Error -> ParseResult.Error(res.errors)
    }
}

fun <T, A, B> Parser<T, A>.map(f: (A, SourceLoc) -> B): Parser<T, B> = Mapper(this, f)

fun <T, A> Parser<T, A>.ignore(): Parser<T, Unit> = map { _, _ -> Unit }

class Binder<S, T, A, B>(val parser: Parser<T, A>, val f: (A, SourceLoc) -> ParseResult<S, B>): Parser<T, B> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<T>): ParseResult<T, B> =
        when (val res = parser.parse(tokens)) {
            is ParseResult.Success ->
                when(val subRes = f(res.result, res.loc)) {
                    is ParseResult.Success -> ParseResult.Success(subRes.result, res.loc, res.left)
                    is ParseResult.Error -> ParseResult.Error(subRes.errors, res.left)
                }
            is ParseResult.Error -> ParseResult.Error(res.errors)
        }
}

fun <S, T, A, B> Parser<T, A>.bind(f: (A, SourceLoc) -> ParseResult<S, B>): Parser<T, B> =
    Binder(this, f)

class testToken<T>(val f: (T) -> Boolean, val err: (T) -> String) : Parser<T, T> {
    override val gen: Gen<List<T>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<T>): ParseResult<T, T> {
        val cur = tokens.cur
        return if (cur == null) tokens.unexpectedEof()
        else if (f(cur)) {
            tokens.parsed(cur, tokens.next())
        } else {
            tokens.error(err(cur), tokens)
        }
    }
}
