package ceres.lang.server

import org.eclipse.lsp4j.*
import org.eclipse.lsp4j.jsonrpc.messages.Either
import org.eclipse.lsp4j.services.LanguageServer
import org.eclipse.lsp4j.services.TextDocumentService
import org.eclipse.lsp4j.services.WorkspaceService
import java.util.concurrent.CompletableFuture

class CeresLanguageServer: LanguageServer, TextDocumentService, WorkspaceService {
    override fun didOpen(params: DidOpenTextDocumentParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun didSave(params: DidSaveTextDocumentParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun didClose(params: DidCloseTextDocumentParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun didChange(params: DidChangeTextDocumentParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun didChangeWatchedFiles(params: DidChangeWatchedFilesParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun didChangeConfiguration(params: DidChangeConfigurationParams?) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun shutdown(): CompletableFuture<Any> {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun getTextDocumentService(): TextDocumentService {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun exit() {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun initialize(params: InitializeParams?): CompletableFuture<InitializeResult> {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun getWorkspaceService(): WorkspaceService {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun codeAction(params: CodeActionParams?): CompletableFuture<MutableList<Either<Command, CodeAction>>> {
        return super.codeAction(params)
    }
}


