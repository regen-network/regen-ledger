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
import java.math.BigInteger

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

sealed class ID

abstract class RootID(): ID()

data class ED25519PubKeyHashID(val bytes: ByteArray): RootID()

data class UUID_ID(val bytes: ByteArray): RootID()

data class ChildKey(val owner: RootID, val name: String): ID()

sealed class Type<T>

object NodeType: Type<Node>()

object StringType: Type<String>()

object IntegerType: Type<BigInteger>()

object DoubleType: Type<Double>()

data class Property<T>(val name: ChildKey, val type: Type<T>)
data class SetProperty<T>(val name: ChildKey, val type: Type<T>)

data class ClassProperty<T>(val property: ClassProperty<T>, val required: Boolean)

data class Class<T>(
    val id: ChildKey,
    val properties: List<ClassProperty<Any>>
)

interface Node {
    val id: ID
    fun <T> get(property: Property<T>): T?
    fun <T> getMany(property: Property<T>): Set<T>
}

interface DB {
    fun get(id: ID): Node
}

typealias  OperationHandler = (state: DB) -> DB

