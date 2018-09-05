package ceres.lang.uniontyped

data class UnionType(
    val canBeNone: Boolean
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
