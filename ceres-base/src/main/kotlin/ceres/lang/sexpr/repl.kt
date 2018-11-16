package ceres.lang.sexpr.repl

import ceres.lang.DoubleType
import ceres.lang.ast.EvalEnv
import ceres.lang.ast.checked
import ceres.lang.sexpr.analyze.exprParser
import ceres.parser.ParseResult
import ceres.parser.char.parseString

interface InputReader {
    suspend fun readLine(prompt: String): String?
}

interface OutputWriter {
    fun write(str: String)
    fun writeLn(str: String)
    fun writeLn()
    fun reportEx(ex: Throwable)
    fun flush()
}

interface Completer {
    fun complete(): Iterable<String>
}

interface IOStreams {
    fun input(completer: Completer): InputReader
    val output: OutputWriter
    val err: OutputWriter
}

suspend fun repl(io: IOStreams, env: EvalEnv) {
    var env = env.with(
        "x" to checked(DoubleType.default, 0.0),
        "y" to checked(DoubleType.default, 1.0)
    )
    val completer = object: Completer {
        override fun complete() = env.keys
    }
    val input = io.input(completer)
    val out = io.output
    val err = io.err
    while(true) {
        var line = input.readLine("ceres=> ")
        if(line == null)
            return
        if(line.isBlank())
            continue
        val parseRes = parseString(exprParser, line)
        when(parseRes) {
            is ParseResult.Success -> {
                try {
                    val expr = parseRes.result
                    val tc = expr.typeCheck(env, true)
//                    val res = expr.eval(env)
//                    out.writeLn(res.toString())
                    out.writeLn(tc.toString())
                } catch (ex: Throwable) {
                    err.reportEx(ex)
                }
            }
            is ParseResult.Error -> {
                err.writeLn(parseRes.toString())
            }
        }
    }
}
