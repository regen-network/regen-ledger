package ceres.avl

import ceres.AVL.IAVLNode
import ceres.crypto.LibSodium

interface IMerkleAVLNode<K: Comparable<K>, V>: IAVLNode<K, V> {
    val hash: ByteArray
}

data class MerkleAVLNode<K: Comparable<K>, V>(
        override val key: K, override val value: V, override val left: IMerkleAVLNode<K, V>?, override val right: IMerkleAVLNode<K, V>?,
        override val height: Int, override val rank: Long, val context: MerkleAVLContext<K, V>): IMerkleAVLNode<K, V> {

    override fun setValue(value: V): IAVLNode<K, V> {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    val keyBytes: ByteArray by lazy {
        context.keySerializer.toBytes(key)
    }

    val valueBytes: ByteArray by lazy {
        context.valueSerializer.toBytes(value)
    }

    override val hash: ByteArray by lazy {
        val hashBuilder = LibSodium.cryptoGenericHashBuilder()
        hashBuilder.update(keyBytes)
        hashBuilder.update(valueBytes)
        val leftHash = left?.hash
        if(leftHash != null) hashBuilder.update(leftHash)
        val rightHash = right?.hash
        if(rightHash != null) hashBuilder.update(rightHash)
        hashBuilder.finish()
    }
}

interface IKVStore {
    suspend fun get(key: ByteArray): ByteArray?
    suspend fun set(key: ByteArray, value: ByteArray)
    suspend fun delete(key: ByteArray)
}

data class StoredMerkleAVLContext<K, V>(val keySerializer: Serializer<K>, val valueSerializer: Serializer<V>, val kvStore: IKVStore)

data class StoredMerkleAVLNode<K: Comparable<K>, V>(val hash: ByteArray, val context: StoredMerkleAVLContext<K, V>) {
    data class StoredData(val bytes: ByteArray) {

    }

    var storedData: StoredData? = null

    suspend fun getStoredData(): StoredData? {
        if(storedData != null)
            return storedData
        val v = context.kvStore.get(hash)
        if(v == null)
            throw IllegalStateException()
        storedData = StoredData(v)
        return storedData
    }

}

interface Serializer<T> {
    fun toBytes(t: T): ByteArray
    fun fromBytes(bytes: ByteArray): T
}

data class MerkleAVLContext<K, V>(val keySerializer: Serializer<K>, val valueSerializer: Serializer<V>)
