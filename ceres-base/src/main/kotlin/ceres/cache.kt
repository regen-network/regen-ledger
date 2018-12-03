package ceres.cache

interface Cache<K, V>: MutableMap<K, V> {
    fun invalidate(key: K)
}

