package ceres.geo.rtree

import ceres.geo.IBoundingBox
import kotlin.math.max
import kotlin.math.min

// Immutable RTree implementation adapted from https://github.com/mourner/rbush

class RTree<T>(val root: Node<T>) {
    fun search(bbox: IBoundingBox): Array<Node<T>> {
        var node: Node<T>? = root
        var res = emptyArray<Node<T>>()

        if(node == null || !intersects(bbox, node))
            return res

        val nodesToSearch = mutableListOf<Node<T>>()

        while(node != null) {
            val leaf = node.leaf
            node.children.forEach {
                if(intersects(bbox, it)) {
                    when {
                        leaf -> res += it
                        contains(bbox, it) -> _all(it, res)
                        else -> nodesToSearch.add(it)
                    }
                }
            }
            node = nodesToSearch.pop()
        }
        return res
    }

    private fun _all(n: Node<T>, res: Array<Node<T>>) {
        var node : Node<T>? = n
        val nodesToSearch = mutableListOf<Node<T>>()
        while(node != null) {
            if(node.leaf) res.plus(node.children)
            else nodesToSearch.add(node)
            node = nodesToSearch.pop()
        }
    }

}

private fun <T> MutableList<T>.pop(): T? {
    val sz = this.size
    if(sz > 0) {
        val x = this[0]
        this.removeAt(0)
        return x
    }
    return null
}

data class Node<T>(
    val children: Array<Node<T>>,
    val height: Int = 1,
    val leaf: Boolean = true,
    override val minX: Double = Double.POSITIVE_INFINITY,
    override val minY: Double = Double.POSITIVE_INFINITY,
    override val maxX: Double = Double.NEGATIVE_INFINITY,
    override val maxY: Double = Double.NEGATIVE_INFINITY
): IBoundingBox

fun <T> extend(a: Node<T>, b: Node<T>): Node<T> =
    a.copy(
        minX = min(a.minX, b.minX),
        minY = min(a.minY, b.minY),
        maxX = max(a.maxX, b.maxX),
        maxY = max(a.maxY, b.maxY)
    )

fun compareNodeMinX(a: IBoundingBox, b: IBoundingBox) =
    a.minX - b.minX

fun compareNodeMinY(a: IBoundingBox, b: IBoundingBox) =
    a.minY - b.minY

fun contains(a: IBoundingBox, b: IBoundingBox) =
    a.minX <= b.minX &&
            a.minY <= b.minY &&
            b.maxX <= a.maxX &&
            b.maxY <= a.maxY

fun intersects(a: IBoundingBox, b: IBoundingBox) =
    b.minX <= a.maxX &&
            b.minY <= a.maxY &&
            b.maxX >= a.minX &&
            b.maxY >= a.minY




