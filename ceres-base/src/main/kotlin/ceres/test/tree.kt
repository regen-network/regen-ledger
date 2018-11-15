package ceres.test

class Tree<A>(val f: () -> Node<A>)

class Node<A>(val value: A, val children: List<Tree<A>>)

fun <A> treeOf(a: A) = Tree({ Node(a, emptyList()) })

fun <A, B> Tree<A>.map(g: (A) -> B): Tree<B> =
        Tree ({
            val node = f()
            Node(g(node.value), node.children.map { it.map(g) })
        })

fun <A> Tree<A>.expand(f: (A) -> List<A>): Tree<A> = TODO()

