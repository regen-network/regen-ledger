package ceres.AVL

import kotlin.math.max

interface IAVLNode<K: Comparable<K>, V> {
    val key: K
    val value: V
    val left: IAVLNode<K, V>?
    val right: IAVLNode<K, V>?
    val height: Int
    val rank: Long
    fun setValue(value: V) : IAVLNode<K, V>
}

typealias AVLNodeFactory<K, V> = (key: K, value: V, left: IAVLNode<K, V>?, right: IAVLNode<K, V>?) -> IAVLNode<K, V>

data class SimpleAVLNode<K: Comparable<K>, V>(override val key: K, override val value: V, override val left: IAVLNode<K, V>?, override val right: IAVLNode<K, V>?, override val height: Int, override val rank: Long) : IAVLNode<K, V> {
    override fun setValue(value: V): IAVLNode<K, V> = this.copy(value = value)

}

fun <K:Comparable<K>, V> makeSimpleAVLNode(key: K, value: V, left: IAVLNode<K, V>?, right: IAVLNode<K, V>?): IAVLNode<K, V> =
    SimpleAVLNode(key, value, left, right, 1 + max(left?.height ?: 0, right?.height ?: 0), 0)


interface IAVLTree<K: Comparable<K>, V> /*: Map<K, V>*/ {
    fun set(key: K, value: V): IAVLTree<K, V>
    fun get(key: K): V?
    fun delete(key: K): IAVLTree<K, V>
    fun containsKey(key: K): Boolean
    val size: Long
    fun isEmpty(): Boolean
}

class SimpleAVLTree<K: Comparable<K>, V>(val root: IAVLNode<K, V>? = null): IAVLTree<K, V> {
    override val size: Long
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun containsKey(key: K) = root?.findNode(key) != null

    override fun get(key: K) = root?.findNode(key)?.value

    override fun isEmpty() = root == null

    override fun set(key: K, value: V) =
        SimpleAVLTree(root.insert(key, value, ::makeSimpleAVLNode))

    override fun delete(key: K) =
            SimpleAVLTree(root.delete(key, ::makeSimpleAVLNode))

}

fun <K: Comparable<K>, V> IAVLNode<K, V>.findNode(key: K): IAVLNode<K, V>? {
    val order = key.compareTo(this.key);
    return when {
        order < 0 -> this.left?.findNode(key)
        order > 0 -> this.right?.findNode(key)
        else -> this;
    }
}

fun <K: Comparable<K>, V> IAVLNode<K, V>?.insert(key: K, value: V, makeNode: AVLNodeFactory<K, V>): IAVLNode<K, V> {
    if(this == null)
        return makeNode(key, value, null,null)
    val thisKey = this.key;
    val order = key.compareTo(thisKey);
    return when {
        order < 0 -> balance(thisKey, this.value, this.left.insert(key, value, makeNode), this.right, makeNode)
        order > 0 -> balance(thisKey, this.value, this.left, this.right.insert(key, value, makeNode), makeNode)
        else -> this.setValue(value);
    }
}

fun <K: Comparable<K>, V> IAVLNode<K, V>?.delete(key: K, makeNode: AVLNodeFactory<K, V>): IAVLNode<K, V> {
    TODO()
}

inline val <K: Comparable<K>, V> IAVLNode<K, V>?.nodeHeight: Int
    get() = this?.height ?: 0

inline val <K: Comparable<K>, V> IAVLNode<K, V>.balanceFactor: Int
    get() = this.left.nodeHeight - this.right.nodeHeight


fun <K: Comparable<K>, V> balance(key: K, value: V, left: IAVLNode<K,V>?, right: IAVLNode<K, V>?, makeNode: AVLNodeFactory<K, V>): IAVLNode<K, V> {
    val diff = left.nodeHeight - right.nodeHeight
    return when {
        // Left Big
        diff == 2 -> {
            if(left == null) throw IllegalStateException()
            val balFactor = left.balanceFactor
            when {
                // Left Heavy
                balFactor >= 0  ->
                    makeNode(left.key, left.value, left.left, makeNode(key, value, left.right, right))
                // Right Heavy
                else -> {
                    val lr = left.right
                    if(lr == null) throw IllegalStateException()
                    makeNode(lr.key, lr.value,
                            makeNode(left.key, left.value, left.left, lr.left),
                            makeNode(key, value, lr.right, right))
                }
            }
        }
        // Right Big
        diff == -2 -> {
            if(right == null) throw IllegalStateException()
            val balFactor = right.balanceFactor
            when {
                // Left Heavy
                balFactor > 0  -> {
                    val rl = right.left
                    if(rl == null) throw IllegalStateException()
                    makeNode(rl.key, rl.value,
                            makeNode(key,value, left, rl.left),
                            makeNode(right.key, right.value, rl.right, right.right)
                    )
                }
                // Right Heavy
                else  -> makeNode(right.key, right.value, makeNode(key, value, left, right.left), right.right)
            }
        }
        else -> makeNode(key, value, left, right)
    }
}

fun <K: Comparable<K>, V> IAVLNode<K, V>?.calcHeight() : Int =
        if(this == null) 0 else max(this.left.calcHeight(), this.right.calcHeight())
