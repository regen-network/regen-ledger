package ceres.crypto

interface HashBuilder {
    fun update(newBytes: ByteArray)
    fun finish(): ByteArray
}

expect object LibSodium {
    fun cryptoGenericHash(bytes: ByteArray): ByteArray
    fun cryptoGenericHashBuilder(): HashBuilder
}

