package ceres.main

import ceres.lang.BaseEvalEnv
import ceres.lang.sexpr.repl.*
import ceres.lang.sexpr.repl.Completer
import org.jline.reader.*
import org.jline.reader.impl.history.DefaultHistory
import org.jline.terminal.TerminalBuilder
import org.jline.utils.AttributedString
import java.io.PrintWriter
import java.io.Writer

class WriterOutput(w: Writer): OutputWriter {
    val writer = PrintWriter(w)

    override fun reportEx(ex: Throwable) {
        ex.printStackTrace(writer)
        writer.flush()
    }

    override fun write(str: String) {
        writer.write(str)
    }

    override fun writeLn(str: String) {
        writer.write(str)
        writeLn()
    }

    override fun writeLn() {
        writer.write("\n")
        writer.flush()
    }

    override fun flush() {
        writer.flush()
    }

}

object ConsoleIO: IOStreams {
    override fun input(completer: Completer): InputReader =
        object: InputReader {
            val history = DefaultHistory()

            init {
                Runtime.getRuntime().addShutdownHook(Thread({
                    history.save()
                }))
            }

            val reader =
                LineReaderBuilder.builder()
                    .appName("Ceres")
                    .variable(LineReader.HISTORY_FILE, ".ceres-repl-history")
                    .history(history)
                    .terminal(TerminalBuilder.builder()
                        .build())
                    .completer(object: org.jline.reader.Completer {
                        override fun complete(
                            reader: LineReader?,
                            line: ParsedLine?,
                            candidates: MutableList<Candidate>?
                        ) {
                            fun strToCandidate(it: String) =
                                Candidate(AttributedString.stripAnsi(it), it, null, null, null, null, true)
                            fun addAll(xs:Iterable<String>) = candidates?.addAll(xs.map(::strToCandidate))
                            addAll(listOf("(", ")", "[", "]", "{", "}"))
                            addAll(completer.complete())
                        }
                    })
                    .build()

            override suspend fun readLine(prompt: String): String? = reader.readLine(prompt)

        }
    override val output: OutputWriter
        get() = WriterOutput(System.out.writer())
    override val err: OutputWriter
        get() = WriterOutput(System.err.writer())
}

suspend fun main() {
    println("Ceres")
    while(true) {
        try {
            repl(ConsoleIO, BaseEvalEnv)
        } catch (ex: UserInterruptException) {
            println("Clearing REPL state (use Ctrl-D to exit)")
            // cancel current operation and restart repl
        } catch (ex: EndOfFileException) {
            println("Goodbye!")
            return
        }
    }
}