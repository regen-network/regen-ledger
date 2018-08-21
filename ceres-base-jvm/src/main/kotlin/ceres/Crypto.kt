package ceres.crypto

import com.goterl.lazycode.lazysodium.LazySodiumJava
import com.goterl.lazycode.lazysodium.SodiumJava
import com.goterl.lazycode.lazysodium.interfaces.GenericHash
import java.nio.charset.StandardCharsets

actual object LibSodium {
    val sodium = SodiumJava()
    val lazySodium = LazySodiumJava(sodium, StandardCharsets.UTF_8)

    actual fun cryptoGenericHash(bytes: ByteArray): ByteArray {
        val res = ByteArray(GenericHash.BYTES)
        sodium.crypto_generichash(res, GenericHash.BYTES, bytes, bytes.size.toLong(), null, 0)
        return res
    }

    actual fun cryptoGenericHashBuilder(): HashBuilder {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    class JavaHashBuilder(): HashBuilder {
        val state = ByteArray(GenericHash.BYTES)

        init {
            sodium.crypto_generichash_init(state, null, 0, GenericHash.BYTES)
        }

        override fun update(newBytes: ByteArray) {
            sodium.crypto_generichash_update(state, newBytes, newBytes.size.toLong())
        }

        override fun finish(): ByteArray {
            val out = ByteArray(GenericHash.BYTES)
            sodium.crypto_generichash_final(state, out, GenericHash.BYTES)
            return out
        }
    }

}