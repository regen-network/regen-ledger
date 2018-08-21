package ceres.util

import ceres.crypto.LibSodium
import com.goterl.lazycode.lazysodium.Sodium

actual object Base64 {
    actual fun encode(bytes: ByteArray): String {
        val sodium = LibSodium.sodium
        val len = sodium.sodium_base64_encoded_len(bytes.size, TODO())
        //sodium.sodium_bin2base64()
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    actual fun decode(str: String): ByteArray {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}