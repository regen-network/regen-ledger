package ceres.lang.codegen

import org.jetbrains.kotlin.cli.common.repl.KotlinJsr223JvmScriptEngineFactoryBase
import org.jetbrains.kotlin.cli.common.repl.ScriptArgsWithTypes
import org.jetbrains.kotlin.script.jsr223.KotlinJsr223JvmLocalScriptEngine
import org.jetbrains.kotlin.script.jsr223.KotlinJsr223JvmLocalScriptEngineFactory
import org.jetbrains.kotlin.script.jsr223.KotlinStandardJsr223ScriptTemplate
import org.jetbrains.kotlin.script.util.scriptCompilationClasspathFromContextOrStlib
import java.io.File
import javax.script.Bindings
import javax.script.ScriptContext
import javax.script.ScriptEngine

class KotlinScriptEngineFactory : KotlinJsr223JvmScriptEngineFactoryBase() {
    override fun getScriptEngine(): ScriptEngine =
        KotlinJsr223JvmLocalScriptEngine(
            this,
            //scriptCompilationClasspathFromContextOrStlib("kotlin-script-runtime.jar", "kotlin-script-util.jar", "kotlin-stdlib-jdk8.jar", wholeClasspath = true),
            //scriptCompilationClasspathFromContextOrStlib("kotlin-script-util.jar", "kotlin-script-runtime.jar", "kotlin-stdlib-jdk8.jar", wholeClasspath = true),
            // TODO figure out how to get proper jars in classpath
            listOf(
                "/Users/Arc/.gradle/caches/modules-2/files-2.1/org.jetbrains.kotlin/kotlin-stdlib/1.3-M2/e475d291eee6606d4512305413b0c9ed4557cae4/kotlin-stdlib-1.3-M2.jar",
                "/Users/Arc/.gradle/caches/modules-2/files-2.1/org.jetbrains.kotlin/kotlin-script-util/1.3-M2/4e57f6d4c6ce3646dceb457544c7166bb036f323/kotlin-script-util-1.3-M2.jar",
                "/Users/Arc/.gradle/caches/modules-2/files-2.1/org.jetbrains.kotlin/kotlin-script-runtime/1.3-M2/4b5d05fcf9bc14881b122fd683f4a6daff6133a5/kotlin-script-runtime-1.3-M2.jar"
            )
                .map
                { File(it) },
            KotlinStandardJsr223ScriptTemplate::
            class.qualifiedName!!,
            { ctx, types ->
                ScriptArgsWithTypes(
                    arrayOf(ctx.getBindings(ScriptContext.ENGINE_SCOPE)),
                    types ?: emptyArray()
                )
            },
            arrayOf(
                Bindings::
                class
            )
        )
}


fun test1() {
    val engine = KotlinScriptEngineFactory().scriptEngine
    println(engine.eval("1 + 1"))
}

