package ceres.storage

interface IKVStore<K, V> {
    suspend fun get(key: K): V?
    suspend fun set(key: K, value: V)
    suspend fun delete(key: K)
}

