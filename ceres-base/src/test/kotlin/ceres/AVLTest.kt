package ceres.AVL

import kotlin.coroutines.experimental.buildSequence
import kotlin.math.abs
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
    return buildSequence {
        yieldAll(This.left.allValues())
        assertEquals(This.key, This.value)
        yield(This.key)
        yieldAll(This.right.allValues())
    }
}

fun <T: Comparable<T>> assertAllValuesInOrder(values: Iterable<T>, tree: IAVLTree<T, T>) {
    assertEquals(values.toList(), tree.root.allValues().toList())
}

fun <T: Comparable<T>> assertIsBalanced(node: IAVLNode<T, T>) {
    assertEquals(node.height, node.calcHeight())
    val diff = node.left.nodeHeight - node.right.nodeHeight
    assertTrue(abs(diff) <= 1)
}

fun <T: Comparable<T>> assertWellBehaved(expected: Iterable<T>, tree: IAVLTree<T, T>) {
    assertAllValuesPresent(expected, tree)
    assertAllValuesInOrder(expected, tree)
    val root = tree.root
    if(root != null)
        assertIsBalanced(root)
}


class AVLTest {
    @Test fun test() {
        for(i in 1..100) {
            Random.next()
        }
    }
}
