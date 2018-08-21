package ceres.util

expect object Base64 {
    fun encode(bytes: ByteArray): String
    fun decode(str: String): ByteArray
}

