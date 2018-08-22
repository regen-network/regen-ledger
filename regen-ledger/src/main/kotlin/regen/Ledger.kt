package regen.ledger

import ceres.rocksdb.*
import com.github.jtendermint.jabci.api.*
import com.github.jtendermint.jabci.socket.TSocket
import com.github.jtendermint.jabci.types.*
import com.google.protobuf.ByteString
import java.nio.ByteBuffer
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.default
import java.io.File

interface Options {
    val port: Int
    val dataDir: File
}

class CommandLineOptions(parser: ArgParser): Options {
    override val port by parser.storing("-p", "--port", help = "Port") { toInt() }
        .default(RegenLedger.DEFAULT_PORT)
    override val dataDir by parser.storing("--data", help = "Data Directory") { File(this) }
        .default(File(".xrn/regen_ledger.db"))
}

class RegenLedger(val opts: Options) : ABCIAPI  {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            ArgParser(args).parseInto(::CommandLineOptions).run {
                RegenLedger(this).start()
            }
        }

        const val DEFAULT_PORT: Int = 26658
    }

    val LAST_BLOCK_HEIGHT = "xrn/LAST_BLOCK_HEIGHT".toByteArray()
    val LAST_BLOCK_APP_HASH = "xrn/LAST_BLOCK_APP_HASH".toByteArray()

    val socket = TSocket()
    init {
        socket.registerListener(this)
    }

    init {
        if(!opts.dataDir.exists())
            opts.dataDir.mkdirs()
    }
    val store = openRocksDb(opts.dataDir.absolutePath)


    fun start() {
        println("Starting Regen Ledger")
        socket.start(opts.port)
    }

    override fun requestInitChain(req: RequestInitChain?): ResponseInitChain =
        ResponseInitChain.getDefaultInstance()

    override fun requestInfo(req: RequestInfo?): ResponseInfo {
        val builder = ResponseInfo.newBuilder()
        val lastBlockHeight = store.get(LAST_BLOCK_HEIGHT)
        builder.setLastBlockHeight(lastBlockHeight?.toLong() ?: 0)
        val lastBlockHash = store.get(LAST_BLOCK_APP_HASH)
        if(lastBlockHash != null)
            builder.setLastBlockAppHash(ByteString.copyFrom(lastBlockHash))
        return builder.build()
    }

    override fun requestBeginBlock(req: RequestBeginBlock?): ResponseBeginBlock =
        ResponseBeginBlock.getDefaultInstance()

    override fun requestCheckTx(req: RequestCheckTx?): ResponseCheckTx =
        ResponseCheckTx.getDefaultInstance()

    override fun receivedDeliverTx(req: RequestDeliverTx?): ResponseDeliverTx =
        ResponseDeliverTx.getDefaultInstance()

    override fun requestCommit(requestCommit: RequestCommit?): ResponseCommit =
        ResponseCommit.getDefaultInstance()

    override fun requestEndBlock(req: RequestEndBlock?): ResponseEndBlock =
        ResponseEndBlock.getDefaultInstance()

    override fun requestQuery(req: RequestQuery?): ResponseQuery =
        ResponseQuery.getDefaultInstance()

    override fun requestSetOption(req: RequestSetOption?): ResponseSetOption =
        ResponseSetOption.getDefaultInstance()

    override fun requestFlush(reqfl: RequestFlush?): ResponseFlush =
        ResponseFlush.getDefaultInstance()

    override fun requestEcho(req: RequestEcho?): ResponseEcho =
        ResponseEcho.newBuilder().setMessage(req?.message).build()
}

fun ByteArray.toLong(): Long =
    ByteBuffer.wrap(this).getLong()

fun Long.toBytes(): ByteArray {
    val buf = ByteBuffer.allocate(Long.SIZE_BYTES)
    buf.putLong(this)
    return buf.array()
}

