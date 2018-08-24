package ceres.db

import ceres.graph.Graph
import ceres.graph.IDB
import ceres.graph.IGraph
import ceres.rocksdb.RocksDBKVStore
import ceres.rocksdb.openRocksDb
import ceres.storage.IKVStore
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.default
import io.ktor.routing.routing
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import mu.KotlinLogging
import java.io.File
import java.nio.file.Paths
import java.util.concurrent.TimeUnit
import java.util.concurrent.atomic.AtomicBoolean

private val logger = KotlinLogging.logger {}

class CeresDB(val store: IKVStore<ByteArray, ByteArray>): IDB {
    fun transact(ceresSrc: String) {
        TODO()
    }

    override fun transact(newGraph: Graph) {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun commit() {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

}

open class CeresDBServer(opts: Options) {
    interface Options {
        val dataDir: String?
        val port: Int?
    }

    open class CommandLineOptions(parser: ArgParser): Options {
        override val port by parser.storing("-p", "--port", help = "Port") { toInt() }
            .default<Int?>(null)
        override val dataDir by parser.storing("--data", help = "Data Directory")
            .default<String?>(null)
    }

    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            ArgParser(args).parseInto(::CommandLineOptions).run {
                CeresDBServer(this).start()
            }
        }
    }

    val DEFAULT_PORT = 3030
    val DEFAULT_DATA_DIR = ".ceres"
    val DATA_FILE_NAME = "data.db"

    val dataDir = opts.dataDir ?: DEFAULT_DATA_DIR

    init {
        val dataDirFile = File(dataDir)
        if(!dataDirFile.exists())
            dataDirFile.mkdirs()
    }

    val store = openRocksDb(Paths.get(dataDir, DATA_FILE_NAME).toString())

    val db = CeresDB(RocksDBKVStore(store))

    val server = embeddedServer(Netty, port=opts.port ?: DEFAULT_PORT) {
        routing {

        }
    }

    val thread = Thread (
        {
            logger.info { "Starting DB Server" }
            server.start(wait = true)
        },
        "DB Server Main Thread"
    )

    val running = AtomicBoolean(false)

    open fun start() {
        running.set(true)
        thread.isDaemon = false
        thread.start()
    }

    init {
        Runtime.getRuntime().addShutdownHook(Thread {
            onShutdown()
        })
    }

    open fun onShutdown() {
        running.set(false)
        logger.info { "Stopping DB Server" }
        server.stop(gracePeriod = 100, timeout = 1000, timeUnit = TimeUnit.MILLISECONDS)
        logger.info { "DB Server Stopped" }
    }
}

