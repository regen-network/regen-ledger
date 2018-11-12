package ceres.lang.uniontyped

import ceres.graph.IRI
import ceres.lang.Refinement

sealed class Allow<in T> {
    object Any: Allow<Any>()
    data class Some<T>(val param: T): Allow<T>()
    object None: Allow<Any>()
}

typealias AllowList<T> = Allow<Set<T>>

data class UnionType(
    // This covers None and Bool types
    val exactIRIs: Set<IRI>,
    val funTypes: AllowList<FunctionType>,
    val intConstraints: Allow<Refinement>,
    val doubleConstraints: Allow<Refinement>,
    val stringConstraints: Boolean
)

data class FunctionType(
  val params: List<Pair<String, UnionType>>,
  val ret: UnionType,
  val suspend: Boolean = false,
  val terminationConstraint: TerminationConstraint = TerminationConstraint.Partial
// TODO allowedPrimitives
// TODO contextParams
)

sealed class TerminationConstraint {
    object Total: TerminationConstraint()
    //TODO data class Bounded(): TerminationConstraint()
    object Partial: TerminationConstraint()
}

/*
data Type = Type
{ canBeNone :: Bool
    , canBeBool :: Bool
    -- , canBeType :: Bool -- want to allow more constraining on type terms
    -- For these type constraints
    , intConstraints :: AllowList Term
    , doubleConstraints :: AllowList Term
    , iriConstraints :: AllowList Term
    , stringConstraints :: AllowList Term
    , objType :: Maybe ObjectType
    , funTypes :: AllowList FunctionType
    , setTypes :: AllowList Type
        -- | The set of allowed opaque (non-constrainable) types
    , opaqueTypes :: AllowList OpaqueType
    , typeTypes :: AllowList Type
    , dependentTypes :: [Term]
} deriving (Eq, Generic)
*/
