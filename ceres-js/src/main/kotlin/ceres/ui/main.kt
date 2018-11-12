package ceres.ui

import react.RBuilder
import react.dom.div
import react.dom.h1
import react.dom.h2
import react.dom.render
import kotlin.browser.document

fun RBuilder.test() {

}

fun main(args: Array<String>) {
    val state: dynamic = module.hot?.let { hot ->
        hot.accept()

        hot.dispose { data ->
            //data.appState = application?.dispose()
            //application = null
        }

        hot.data
    }

    render(document.getElementById("root")) {
        div {
            h1 {
                +"Hello World!"
            }
            h2 {
                +"Let's see how long this takes!!!"
            }
        }
    }
}

external val module: Module

external interface Module {
    val hot: Hot?
}

external interface Hot {
    val data: dynamic

    fun accept()
    fun accept(dependency: String, callback: () -> Unit)
    fun accept(dependencies: Array<String>, callback: (updated: Array<String>) -> Unit)

    fun dispose(callback: (data: dynamic) -> Unit)
}

external fun require(name: String): dynamic