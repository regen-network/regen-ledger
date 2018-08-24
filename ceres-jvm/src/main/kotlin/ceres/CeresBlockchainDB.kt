package ceres.db.blockchain

import ceres.db.CeresDBServer
import com.github.jtendermint.jabci.api.ABCIAPI
import com.github.jtendermint.jabci.socket.TSocket
import com.github.jtendermint.jabci.types.*
import com.google.protobuf.ByteString
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.default
import mu.KotlinLogging
import java.nio.ByteBuffer

private val logger = KotlinLogging.logger {}

open class CeresBlockchainDBServer(val blockchainOpts: Options): CeresDBServer(blockchainOpts), ABCIAPI {
    interface Options: CeresDBServer.Options {
        val tendermintPort: Int?

    }

    class CommandLineOptions(parser: ArgParser): CeresDBServer.CommandLineOptions(parser), Options {
      override val tendermintPort by parser.storing("-t", "--tendermint-port", help = "Tendermint Port") { toInt() }
        .default<Int?>(null)
    }

    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            ArgParser(args).parseInto(::CommandLineOptions).run {
                CeresBlockchainDBServer(this).start()
            }
        }

    }

    val DEFAULT_TENDERMINT_PORT: Int = 26658
    val LAST_BLOCK_HEIGHT = "LAST_BLOCK_HEIGHT".toByteArray()
    val LAST_BLOCK_APP_HASH = "LAST_BLOCK_APP_HASH".toByteArray()

    val socket = TSocket()

    init {
        socket.registerListener(this)
    }

    val blockchainThread = Thread(
        {
            logger.info { "Starting ABCI App" }
            socket.start(blockchainOpts.tendermintPort ?: DEFAULT_TENDERMINT_PORT)
        },
        "Blockchain Main Thread"
    )

    override fun start() {
        super.start()
        blockchainThread.isDaemon = false
        blockchainThread.start()
    }

    override fun onShutdown() {
        super.onShutdown()
        logger.info { "Stopping ABCI App" }
        socket.stop()
        logger.info { "ABCI App Stopped" }
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

