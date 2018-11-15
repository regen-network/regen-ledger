package ceres.parser.char

import ceres.parser.*
import ceres.test.gen.Gen

fun <T> parseString(parser: Parser<Char, T>, str: String, uri: String? = null): ParseResult<Char, T> {
    val ts = CharTokenSource(str, 0, StringSource(str, uri))
    if (str.length <= 0)
        return ts.unexpectedEof()
    return parser.parse(ts)
}


data class CharTokenSource(val charSequence: CharSequence, override val loc: Int, override val source: Source) :
    TokenSource<Char> {
    override val cur: Char? by lazy {
        if (loc < charSequence.length) charSequence[loc] else null
    }

    override fun next(): TokenSource<Char> {
        if (cur != null)
            return copy(loc = loc + 1)
        return this
    }

    override fun toString(): String = "CharTokenSource(loc = ${loc}, source = ${source})"
}

data class StringSource(val text: CharSequence, override val uri: String?) : Source {
    override fun toString(): String = "StringSource(${uri})"
}

fun char(x: Char): Parser<Char, Char> = testToken({ it == x }, { "Expected ${x}, got ${it}" })

class str(val x: String) : Parser<Char, String> {
    override val gen: Gen<Sequence<Char>>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun parse(tokens: TokenSource<Char>): ParseResult<Char, String> {
        var curTs: TokenSource<Char> = tokens
        for (ch in x) {
            val cur = curTs.cur
            if (cur == null) return tokens.unexpectedEof()
            else if (cur == ch) curTs = curTs.next()
            else return tokens.error("Expected ${x}", curTs)
        }
        return tokens.parsed(x, curTs)
    }
}
