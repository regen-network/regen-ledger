package ceres.ast

expect class Integer

expect class Real

data class ArgDecl(val name: String, val value: Expr?)

data class FunBody(val statements: List<FunStatement>)

sealed class FunStatement {
    data class LocalBinding(val name: String, val value: Expr) : FunStatement()
    data class ExprStatement(val expr: Expr) : FunStatement()
}

sealed class NodeBinding {
    data class PropValue(val name: Expr.ID, val value: Expr)
    data class NodeID(val id: Expr.ID)
}

sealed class Expr {
    data class FunCall(val fn: Expr, val args: Map<String, Expr>) : Expr()
    data class Fun(
        val name: String,
        val args: List<ArgDecl>,
        val body: FunBody
    ) : Expr()

    data class DoubleL(val x: Double) : Expr()
    data class IntegerL(val x: Integer) : Expr()
    data class RealL(val x: Real) : Expr()
    data class StringL(val x: String) : Expr()
    data class ID(val x: String) : Expr()
    data class NodeE(val props: List<NodeBinding>) : Expr()
    data class GraphE(val nodes: List<NodeE>) : Expr()
    data class Free(val x: String)
    data class SelectE(val where: List<Expr>)
}

enum class Visibility {
    Public, Private
}

data class Decl(
    val visibility: Visibility,
    val name: String,
    val value: Expr
)

data class Module(val decls: List<Decl>)


