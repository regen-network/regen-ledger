package ceres.data

import ceres.data.avl.AVLSet
import ceres.data.avl.SimpleAVLTree

interface PersistentMap<K, V>: Map<K, V> {
    fun set(key: K, value: V): PersistentMap<K, V>

    fun setMany(pairs: Sequence<Pair<K, V>>): PersistentMap<K, V> =
        pairs.fold(this, {acc, elem -> acc.set(elem.first, elem.second)})

    fun setMany(pairs: Iterable<Pair<K, V>>): PersistentMap<K, V> =
        pairs.fold(this, {acc, elem -> acc.set(elem.first, elem.second)})

    fun delete(key: K): PersistentMap<K, V>
}

interface PersistentSet<K>: Set<K> {
    fun add(key: K): PersistentSet<K>
    fun addMany(ks: Iterable<K>): PersistentSet<K> =
        ks.fold(this, {acc, elem -> acc.add(elem)})
}

fun <K: Comparable<K>, V> emptyAvlMap(): PersistentMap<K, V> = SimpleAVLTree<K, V>()

fun <K: Comparable<K>, V> Map<K,V>.toAvlMap(): PersistentMap<K, V> =
    emptyAvlMap<K, V>().setMany(entries.map { it.key to it.value })

fun <K: Comparable<K>, V> avlMapOf(vararg pairs: Pair<K, V>): PersistentMap<K, V> =
    emptyAvlMap<K, V>().setMany(pairs.asIterable())

fun <K: Comparable<K>, V> avlMapOf(pairs: Iterable<Pair<K, V>>): PersistentMap<K, V> =
    emptyAvlMap<K, V>().setMany(pairs)

fun <K: Comparable<K>> emptyAvlSet(): PersistentSet<K> = AVLSet<K>()

fun <K: Comparable<K>> Set<K>.toAvlSet(): PersistentSet<K> =
    emptyAvlSet<K>().addMany(this)

fun <K: Comparable<K>> avlSetOf(vararg keys: K): PersistentSet<K> =
    emptyAvlSet<K>().addMany(keys.asIterable())
