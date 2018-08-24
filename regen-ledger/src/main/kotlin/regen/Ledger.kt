package regen.ledger

import ceres.db.blockchain.CeresBlockchainDBServer
import ceres.rocksdb.*
import com.github.jtendermint.jabci.api.*
import com.github.jtendermint.jabci.socket.TSocket
import com.github.jtendermint.jabci.types.*
import com.google.protobuf.ByteString
import java.nio.ByteBuffer
import com.xenomachina.argparser.ArgParser
import com.xenomachina.argparser.default
import java.io.File

class RegenLedger(opts: Options) : CeresBlockchainDBServer(opts)  {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            ArgParser(args).parseInto(::CommandLineOptions).run {
                RegenLedger(this).start()
            }
        }
    }

    //override val DEFAULT_DATA_DIR = ".xrn"
}
