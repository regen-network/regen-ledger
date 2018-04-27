{-# LANGUAGE TypeInType #-}
module Ceres.Core where

import XRN.Prelude hiding (Type)
import qualified Data.Text as T
import Text.Megaparsec (SourcePos(..))
import Ceres.Loc
import qualified Ceres.SMT as SMT

newtype IRI = IRI Text
  deriving (Eq, Generic)

instance Hashable IRI

instance IsString IRI where
  fromString str = IRI $ T.pack str

type Env = HashMap Name (Term, Type)

newtype Name where
  IRIName :: IRI -> Name
  deriving (Eq, Generic)

instance Hashable Name

instance IsString Name where
  fromString str = IRIName $ IRI $ T.pack str

ceres :: Text -> Name
ceres x = IRIName $ IRI $ "ceres:" <> x

addInt :: Name
addInt = ceres "+"

subInt :: Name
subInt = ceres "-"

multInt :: Name
multInt = ceres "*"

addDouble :: Name
addDouble = ceres "+."

subDouble :: Name
subDouble = ceres "-."

multDouble :: Name
multDouble = ceres "*."

divDouble :: Name
divDouble = ceres "/."

eq :: Name
eq = ceres "=="

notEq :: Name
notEq = ceres "!="

gtInt :: Name
gtInt = ceres ">"

gteInt :: Name
gteInt = ceres ">="

ltInt :: Name
ltInt = ceres "<"

lteInt :: Name
lteInt = ceres "<="

data Term where
  LiteralT :: Loc -> Literal -> Term
  -- TypeT :: Type -> Term
  VarT :: Loc -> Name -> Term
  ObjT :: Loc -> ObjTerm -> Term
  SetT :: Loc -> [Term] -> Term
  AppT :: Loc -> Term -> ObjTerm -> Term
  CaseT :: Loc -> [(Term, Term)] -> (Maybe Term) -> Term
  -- ForeignT :: proxy a -> Ptr a -> OpaqueType -> Term
  deriving (Eq, Generic)

instance Hashable Term

data ObjTerm = ObjTerm [ObjBinding]
  deriving (Eq, Generic)

instance Hashable ObjTerm

data ObjBinding = ObjBinding
  { loc :: Loc
  , name :: Text
  , typ :: Maybe Term
  , value :: Term
  } deriving (Eq, Generic)

instance Hashable ObjBinding

data Literal =
  None |
  IRI_C IRI |
  IntC Integer |
  DoubleC Double |
  StringC Text |
  BoolC Bool
  deriving (Eq, Generic)

instance Hashable Literal

data AllowList a = AllowAny | AllowOnly (HashSet a) | AllowNone
  deriving (Eq, Generic)

instance Hashable a => Hashable (AllowList a)

-- | A type constraint in Ceres' core type system. Types are represented
-- as a union of possible types that a term could inhabit. So
-- canBeNone would allow the term to have None as a valid value.
-- For constraints such as 'intConstraints', a value of 'Nothing'
-- would indicate that no integers would be allowed, a value of
-- 'Just []' would indicate that all integers that are allowed,
-- and 'Just xs' where xs is a non-empty list of terms constraining
-- that integer value (with the integer value bound to the implicit variable 'it')
-- indicates a refinement type of integer.
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

instance Hashable Type

checkSubtype :: MonadError SubtypeErrorReason m => Type -> Type -> m ()
checkSubtype parent child = do
  checkSimple canBeNone noneT
  checkSimple canBeBool boolT
  checkInt
  checkDouble
  error "TODO finish cases"
  where
    checkSimple f ty =
      unless (not (f child) || f parent) $ throwError $ MissingType ty
    checkInt = checkRefined (error "TODO refine int") intT (intConstraints parent) (doubleConstraints parent)
    checkDouble = checkRefined (error "TODO refine double") doubleT (doubleConstraints parent) (doubleConstraints parent)
    checkRefined f ty prt chd
      | AllowAny <- prt, _ <- chd = pure ()
      | _ <- prt, AllowNone <- chd = pure ()
      | AllowOnly _ <- prt, AllowNone <- chd = pure ()
      | AllowOnly xs <- prt, AllowAny <- chd = throwError $ MissingRefinements xs
      | AllowOnly xs <- prt, AllowOnly ys <- chd = f xs ys
      | AllowNone <- prt, _ <- chd = throwError $ MissingType ty

data TypeError =
  ObjectTypeMissingGlobalAttrs
  { expected :: ObjectType,
    actual :: ObjectType,
    missingAttrs :: HashSet IRI
  } |
  CantInferType
  { term :: Term
  } |
  SubtypeDoesntContainType

data SubtypeError = SubtypeError
  { parent :: Type
  , child :: Type
  , reason :: SubtypeErrorReason
  }
  
data SubtypeErrorReason =
  MissingType Type |
  MissingRefinements (HashSet Term)


checkSubObjectType :: ObjectType -> ObjectType -> TC ()
checkSubObjectType parent child = do
  areGlobalAttrsSubset
  where
    areGlobalAttrsSubset = do
      let diff = difference (globalAttrs parent) (globalAttrs child)
      unless (null diff) $ throwError $
        ObjectTypeMissingGlobalAttrs
        { expected = parent
        , actual = child
        , missingAttrs = diff
        }

-- | Checks whether 'x' is a subset of 'y'
isSubset :: SetContainer set => set -> set -> Bool
isSubset x y = null $ difference x y

s0 :: HashSet Integer
s0 = setFromList [1, 2, 3, 5]

s1 :: HashSet Integer
s1 = setFromList [1, 2, 3, 4]

uninhabited :: Type
uninhabited =
  Type
  { canBeNone = False
  , canBeBool = False
  , intConstraints = AllowNone
  , doubleConstraints = AllowNone
  , iriConstraints = AllowNone
  , stringConstraints = AllowNone
  , objType = Nothing
  , funTypes = AllowNone
  , setTypes = AllowNone
  , opaqueTypes = AllowNone
  , typeTypes = AllowNone
  , dependentTypes = []
  }

instance Default Type where
  def = uninhabited

noneT :: Type
noneT = def { canBeNone = True }

anyT :: Type
anyT =
  Type
  { canBeNone = True
  , canBeBool = True
  , intConstraints = AllowAny
  , doubleConstraints = AllowAny
  , iriConstraints = AllowAny
  , stringConstraints = AllowAny
  , objType = Just def
  , funTypes = AllowAny
  , setTypes = AllowAny
  , opaqueTypes = AllowAny
  , typeTypes = AllowAny
  , dependentTypes = []
  }

intT :: Type
intT = def { intConstraints = AllowAny }

doubleT :: Type
doubleT = def { doubleConstraints = AllowAny }

boolT :: Type
boolT = def { canBeBool = True }

stringT :: Type
stringT = def { stringConstraints = AllowAny }

iriT :: Type
iriT = def { iriConstraints = AllowAny }

data ObjectType = ObjectType
  { localAttrs :: HashMap Text Type
  , globalAttrs :: HashSet IRI
  , constraints :: HashSet Term
  } deriving (Eq, Generic)

instance Hashable ObjectType

instance Default ObjectType where
  def =
    ObjectType
    { localAttrs = mempty
    , globalAttrs = mempty
    , constraints = mempty
    }

data TerminationConstraint = Partial | Total | Bounded Term
  deriving (Eq, Generic)

instance Hashable TerminationConstraint

data FunctionType = FunctionType
  { args :: ObjectType
  , ret :: Type
  , allowedPrims :: AllowList IRI
  , termination :: TerminationConstraint
  } deriving (Eq, Generic)

instance Default FunctionType where
  def =
    FunctionType
    { args = def
    , ret = def
    , allowedPrims = AllowAny
    , termination = Partial
    }

instance Hashable FunctionType

-- | Represents an opaque (probably platform-specific) type
newtype OpaqueType = OpaqueType Word64
  deriving (Eq, Generic)

instance Hashable OpaqueType

data AttrRef where
  LocalAttr :: Text -> Type -> AttrRef
  GlobalAttr :: IRI -> AttrRef
  deriving (Eq, Generic)

instance Hashable AttrRef

binIntFnTy :: Type
binIntFnTy = binFnTy intT

compIntFnTy :: Type
compIntFnTy = compFnTy intT

binDoubleFnTy :: Type
binDoubleFnTy = binFnTy doubleT

compDoubleFnTy :: Type
compDoubleFnTy = compFnTy doubleT

binFnTy :: Type -> Type
binFnTy ty = binFnTy' ty ty

compFnTy :: Type -> Type
compFnTy ty = binFnTy' ty boolT

allowOne :: (Eq a, Hashable a) => a -> AllowList a
allowOne x = AllowOnly (setFromList [x])

binFnTy' :: Type -> Type -> Type
binFnTy' ty ret = def
  { funTypes =
    allowOne $ def
    { args = def { localAttrs = mapFromList [("x", ty), ("y", ty)] }
    , ret
    , termination = Bounded (LiteralT Internal (IntC 1))
    }
  }

prims :: HashMap Name Type
prims =
  mapFromList
  [(addInt, binIntFnTy)
  ,(subInt, binIntFnTy)
  ,(multInt, binIntFnTy)
  ,(addDouble, binDoubleFnTy)
  ,(subDouble, binDoubleFnTy)
  ,(multDouble, binDoubleFnTy)
  ,(divDouble, binDoubleFnTy)
  ,(eq, compFnTy anyT)
  ]

type TC = ExceptT TypeError (ReaderT () IO)

checkType :: Env -> Term -> Type -> TC (Term, Type)
checkType env term typ@Type{..}
  | LiteralT _ None <- term, canBeNone = pure (term, noneT)
  | LiteralT _ (BoolC _) <- term, canBeBool = pure (term, boolT)
  | LiteralT _ (IntC _) <- term, AllowAny <- intConstraints = pure (term, intT)
  | LiteralT _ (DoubleC _) <- term, AllowAny <- doubleConstraints = pure (term, doubleT)
  | LiteralT _ (StringC _) <- term, AllowAny <- stringConstraints = pure (term, stringT)
  | LiteralT _ (IRI_C _) <- term, AllowAny <- iriConstraints = pure (term, iriT)
  -- | VarT name <- term =
  --     case lookup name env of
  --       Just (term',typ') -> checkType env term' typ
  --       Nothing -> False
  -- | TypeT _ <- term, canBeType = True

inferType :: Env -> Term -> TC Type
inferType env term
  | LiteralT _ None <- term = pure noneT
  | LiteralT _ (BoolC _) <- term = pure boolT
  | LiteralT _ (IntC _) <- term = pure intT
  | LiteralT _ (DoubleC _) <- term = pure doubleT
  | LiteralT _ (StringC _) <- term = pure stringT
  | LiteralT _ (IRI_C _) <- term = pure iriT
  | LiteralT _ (IRI_C _) <- term = pure iriT
  | otherwise = throwError $ CantInferType { term }

-- | Tries to check the 'Term' as a valid 'Type' expression
-- in the given 'Env'.
checkTypeType :: Env -> Term -> Maybe Type
checkTypeType _ _ = Nothing

-- intConstraintToSmt :: MonadError Text m => Term -> m (SMT.SMTExpr Bool)
-- intConstraintToSmt (LiteralT (IntC x)) = pure $ SMT.IntegerConst x

-- data ObjConstraint = ObjConstraint
--   { attrs :: [Attr]
--   -- TODO assertions
--   }

-- data Typ env = Typ
--   { objType :: Maybe ObjConstraint
--   , canBeNone :: Bool
--   , canBeString :: Bool
--   -- , stringConstraint :: Maybe (StringConstraint Bool)
--   , canBeIRI :: Bool
--   -- , iriConstraint :: Maybe (IRIConstraint Bool)
--   , canBeInt :: Bool
--   -- , intConstraint :: Maybe (IntConstraint Bool)
--   , canBeDouble :: Bool
--   -- , doubleConstraint :: Maybe (DoubleConstraint Bool)
--   , setType :: Maybe Typ
--   , functionType :: Maybe FunctionType
--   -- TODO dependent terms
--   }


-- data FunctionType env = FunctionType
--   { args :: ObjConstraint
--   , ret :: Typ env
--   }


typeUnion :: Type -> Type -> Type
typeUnion x y =
  Type
  { canBeNone = canBeNone x || canBeNone y
  , canBeBool = canBeBool x || canBeBool y
  , intConstraints = constraintUnion (intConstraints x) (intConstraints y)
  , doubleConstraints = constraintUnion (doubleConstraints x) (doubleConstraints y)
  , stringConstraints = constraintUnion (stringConstraints x) (stringConstraints y)
  , iriConstraints = constraintUnion (iriConstraints x) (iriConstraints y)
  , objType = objUnion (objType x) (objType y)
  , funTypes = error "TODO"
  , setTypes = error "TODO"
  }
  where
    genUnion _ AllowAny _ = AllowAny
    genUnion _ _ AllowAny = AllowAny
    genUnion f (AllowOnly xs) (AllowOnly ys) = AllowOnly (f xs ys)
    genUnion _ (AllowOnly xs) AllowNone = AllowOnly xs
    genUnion _ AllowNone (AllowOnly xs) = AllowOnly xs
    genUnion _ AllowNone AllowNone = AllowNone
    constraintUnion = genUnion union
    objUnion = error "TODO"

typeIntersection :: Type -> Type -> Type
typeIntersection x y =
  Type
  { canBeNone = canBeNone x && canBeNone y
  , canBeBool = canBeBool x && canBeBool y
  , intConstraints = constraintIntersection (intConstraints x) (intConstraints y)
  , doubleConstraints = constraintIntersection (doubleConstraints x) (doubleConstraints y)
  , stringConstraints = constraintIntersection (stringConstraints x) (stringConstraints y)
  , iriConstraints = constraintIntersection (iriConstraints x) (iriConstraints y)
  , objType = objIntersection (objType x) (objType y)
  , funTypes = error "TODO"
  , setTypes = error "TODO"
  }
  where
    genIntersection _ AllowNone _ = AllowNone
    genIntersection _ _ AllowNone = AllowNone
    genIntersection f (AllowOnly xs) (AllowOnly ys) = AllowOnly (f xs ys)
    genIntersection _ (AllowOnly xs) AllowAny = AllowOnly xs
    genIntersection _ AllowAny (AllowOnly xs) = AllowOnly xs
    genIntersection _ AllowAny AllowAny = AllowAny
    constraintIntersection = genIntersection union
    objIntersection = error "TODO"


data Value =
  LitV Literal |
  ObjV (HashMap Text Value) |
  SetV (Set Value)
  -- FunV
  deriving (Eq, Generic)
