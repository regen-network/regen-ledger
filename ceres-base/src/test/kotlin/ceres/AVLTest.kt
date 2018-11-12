package ceres.AVL

import kotlin.math.abs
import kotlin.random.Random
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull
import kotlin.test.assertTrue

fun <T: Comparable<T>> assertAllValuesPresent(values: Iterable<T>, tree: IAVLTree<T, T>) {
    values.forEach {
        assertEquals(it, tree.get(it))
    }
}

fun <T: Comparable<T>> IAVLNode<T, T>?.allValues(): Sequence<T> {
    if(this == null)
        return emptySequence()
    val This = this
    return sequence {
        yieldAll(This.left.allValues())
        assertEquals(This.key, This.value)
        yield(This.key)
        yieldAll(This.right.allValues())
    }
}

fun <T: Comparable<T>> assertAllValuesInOrder(values: Iterable<T>, tree: IAVLTree<T, T>) {
    val treeValues = tree.root.allValues().toList()
    //println("treeValues:${treeValues}")
    assertEquals(values.toList(), treeValues)
}

fun <T: Comparable<T>, V> IAVLNode<T, V>.assertIsBalanced() {
    assertEquals(height, calcHeight())
    val diff = left.nodeHeight - right.nodeHeight
    assertTrue(abs(diff) <= 1)
    left?.assertIsBalanced()
    right?.assertIsBalanced()
}

fun <T: Comparable<T>> assertWellBehaved(expected: Iterable<T>, tree: IAVLTree<T, T>) {
    assertAllValuesPresent(expected, tree)
    assertAllValuesInOrder(expected, tree)
    tree.root?.assertIsBalanced()
}

fun <T: Comparable<T>> makeTreeSet(values: Iterable<T>) = SimpleAVLTree<T, T>().setMany(values.map { it to it })

fun <T: Comparable<T>> testTree(values: Iterable<T>) = assertWellBehaved(values, makeTreeSet(values))

class AVLTest {
    @Test fun test() {
        for(i in 1..100) {
            val values = generateSequence { Random.nextDouble() }.take(100).sorted().toList()
            //println("values:${values}")
            testTree(values.toList())
        }
    }
}
