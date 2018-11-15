package ceres.parser.char

import ceres.parser.*

fun <T> parseString(parser: Parser<Char, T>, str: String, uri: String? = null): ParseResult<Char, T> {
    val ts = CharTokenSource(str, 0, StringSource(str, uri))
    if (str.length <= 0)
        return ts.unexpectedEof()
    return parser(ts)
}


data class CharTokenSource(val charSequence: CharSequence, override val loc: Int, override val source: Source) :
    TokenSource<Char> {
    override val cur: Char?
        get() = if(loc < charSequence.length) charSequence[loc] else null

    override fun next(): TokenSource<Char> {
        if(loc < charSequence.length)
            return copy(loc = loc + 1)
        return this
    }

    override fun toString(): String = "CharTokenSource(loc = ${loc}, source = ${source})"
}

data class StringSource(val text: CharSequence, override val uri: String?) : Source {
    override fun toString(): String = "StringSource(${uri})"
}

fun testChar(f: (Char) -> Boolean, err: (Char) -> String): Parser<Char, Char> = {
    val cur = it.cur
    if (cur == null) it.unexpectedEof()
    else if (f(cur)) {
        it.parsed(cur, it.next())
    } else {
        it.error(err(cur), it)
    }
}

fun char(x: Char): Parser<Char, Char> = testChar({ it == x }, { "Expected ${x}, got ${it}" })

fun str(x: String): Parser<Char, String> = fun(it: TokenSource<Char>): ParseResult<Char, String> {
    var curTs: TokenSource<Char> = it
    for (ch in x) {
        val cur = curTs.cur
        if (cur == null) return it.unexpectedEof()
        else if (cur == ch) curTs = curTs.next()
        else return it.error("Expected ${x}", curTs)
    }
    return it.parsed(x, curTs)
}
