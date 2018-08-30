package ceres

import ceres.xtext.CeresRuntimeModule
import ceres.xtext.CeresStandaloneSetup
import ceres.xtext.ceres.Model
import com.google.inject.Guice
import org.eclipse.xtext.parser.IParser
import java.io.Reader
import java.io.StringReader
import javax.inject.Inject

class CeresLanguage {
    @Inject
    lateinit var parser: IParser

    init {
        val injector = CeresStandaloneSetup().createInjectorAndDoEMFRegistration()
        injector.injectMembers(this)
    }

    fun parse(reader: Reader): Model? {
        val res = parser.parse(reader)
        if(res.hasSyntaxErrors()) {
            println(res.syntaxErrors)
            return null
        } else {
            return res.rootASTElement as Model
        }
    }

    fun test1() =
        parse(StringReader("x = 1;"))
}