package ceres.lang

import ceres.lang.smtlib.RExpr

sealed class Type

data class FunctionType(
    val params: List<Pair<String, Type>>,
    val ret: Type,
    val suspend: Boolean = false,
    val terminationConstraint: TerminationConstraint = TerminationConstraint.Partial
// TODO allowedPrimitives
// TODO contextParams
): Type()

sealed class TerminationConstraint {
    object Total: TerminationConstraint()
    //TODO data class Bounded(): TerminationConstraint()
    object Partial: TerminationConstraint()
}

typealias Refinement = (varname: String) -> RExpr<Boolean>

data class IntegerT(val refinement: Refinement): Type()

data class DoubleT(val refinement: Refinement): Type()

object StringT: Type()

object BoolT: Type()

object NoneT: Type()

data class OpaquePlatformType(val id: String): Type ()

data class SetT(val elemType: Type): Type ()

data class ListT(val elemType: Type): Type ()

data class NodeT(
    val paramTypes: List<NodeParamType>
  //TODO refinements
): Type ()

data class NodeParamType(
    val name: String,
    val type: Type,
    val cardinality: Cardinality
)

data class Cardinality (
    val min: Int,
    val max: Int? = null
)

//TODO data class GraphT(): Type ()
