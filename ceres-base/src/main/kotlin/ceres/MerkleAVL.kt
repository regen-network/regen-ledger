package ceres.avl

import ceres.crypto.LibSodium
import ceres.data.avl.IAVLNode
import ceres.storage.IKVStore

interface IMerkleAVLNode<K: Comparable<K>, V>: IAVLNode<K, V> {
    val hash: ByteArray
}

data class MerkleAVLContext<K, V, S>(val keySerializer: Serializer<K, S>, val valueSerializer: Serializer<V, S>)

data class MerkleAVLNode<K: Comparable<K>, V>(
        override val key: K, override val value: V, override val left: IMerkleAVLNode<K, V>?, override val right: IMerkleAVLNode<K, V>?,
        override val height: Int, override val rank: Long, val context: MerkleAVLContext<K, V, *>): IMerkleAVLNode<K, V> {

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


data class StoredMerkleAVLContext<K, V, S>(val keySerializer: Serializer<K, S>, val valueSerializer: Serializer<V, S>, val kvStore: IKVStore<S, S>)

data class StoredMerkleAVLNode<K: Comparable<K>, V, S>(val id: S, val context: StoredMerkleAVLContext<K, V, S>) {
    inner class StoredData(val data: S) {

    }

    var storedData: StoredData? = null

    suspend fun getStoredData(): StoredData? {
        if(storedData != null)
            return storedData
        val v = context.kvStore.get(id)
        if(v == null)
            throw IllegalStateException()
        storedData = StoredData(v)
        return storedData
    }

}

interface Serializer<T, S> {
    fun toBytes(t: T): ByteArray
    fun serialize(t: T): S
    fun deserialize(s: S): T
}
