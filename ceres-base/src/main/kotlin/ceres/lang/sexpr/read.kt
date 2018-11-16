package ceres.lang.sexpr.read

import ceres.lang.sexpr.*
import ceres.parser.*
import ceres.parser.char.*
import ceres.test.gen.Gen

//fun reader(tokenSource: TokenSource<Char>)
//        = reader_.parse(tokenSource)
//
//fun sexpr(tokenSource: TokenSource<Char>)
//        = sexpr.parse(tokenSource)

fun read(str: String, uri: String? = null) =
    parseString(reader, str, uri)

fun readExpr(str: String, uri: String? = null) =
    parseString(sexpr, str, uri)

object reader: Parser<Char, List<SExpr>> {
    override val gen: Gen<List<Char>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<Char>): ParseResult<Char, List<SExpr>> =
            reader_.parse(tokens)
}

object sexpr: Parser<Char, SExpr> {
    override val gen: Gen<List<Char>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<Char>): ParseResult<Char, SExpr> =
            sexpr_.parse(tokens)
}

val ws: Parser<Char, Unit> = star(testToken<Char>({
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
    cat(char('('), reader, char(')')).map { x, loc ->
        Parens(x.second, loc)
    }

val curly =
    cat(char('{'), reader, char('}')).map { x, loc ->
        Curly(x.second, loc)
    }

val square =
    cat(char('['), reader, char(']')).map { x, loc ->
        Square(x.second, loc)
    }

fun idStartChar(ch: Char): Boolean =
    ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
            || ch == '_'
            || ch == '>'
            || ch == '<'
            || ch == '='
            || ch == '+'
            || ch == '-'
            || ch == '.'
            || ch == '*'
            || ch == '!'

fun idChar(ch: Char): Boolean =
    idStartChar(ch) || ch >= '0' && ch <= '9'

val identifier: Parser<Char, String> =
    cat(
        testToken(::idStartChar, { "${it} is not a valid identifier start character" }),
        star(testToken(::idChar, { "${it} is not a valid identifier character" }))
    ).map { x, _ ->
        String(charArrayOf(x.first) + x.second)
    }

val symbol: Parser<Char, Symbol> = identifier.map { x, loc -> Symbol(x, loc) }

val kw: Parser<Char, Keyword> =
    cat(char(':'), identifier).map { x, loc -> Keyword(x.second, loc) }

//val str: Parser<Char, Str> = TODO()
//
//val num: Parser<Char, Num> = TODO()

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

val sexpr_: Parser<Char, SExpr> =
    alt<Char, SExpr>(
        parens, curly, square,
        bool, nil,
        symbol, kw,
//         TODO str, num,
        tagged
    )

val reader_: Parser<Char, List<SExpr>> =
    cat(star(cat(opt(ws), sexpr_).map { x, _ -> x.second}),
        opt(ws)).map {x,_ -> x.first}

