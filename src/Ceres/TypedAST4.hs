{-# LANGUAGE TypeInType #-}
module Ceres.TypedAST3 where

import XRN.Prelude hiding (Type)
import qualified Data.Text as T
import qualified Ceres.SMT as SMT

newtype IRI = IRI Text
  deriving (Eq)

instance IsString IRI where
  fromString str = IRI $ T.pack str

type Env = HashMap Name (Term, Type)

newtype Name where
  IRIName :: IRI -> Name
  deriving (Eq)

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
  LiteralT :: Literal -> Term
  TypeT :: Type -> Term
  VarT :: Name -> Term
  ObjT :: ObjTerm -> Term
  SetT :: [Term] -> Term
  AppT :: Term -> ObjTerm -> Term
  CaseT :: [(Term, Term)] -> (Maybe Term) -> Term
  -- ForeignT :: proxy a -> Ptr a -> OpaqueType -> Term

data ObjTerm = ObjTerm [ObjBinding]

data ObjBinding = ObjBinding
  { name :: Text
  , typ :: Maybe Term
  , value :: Term
  }

data Literal =
  None |
  IRI_C IRI |
  IntC Integer |
  DoubleC Double |
  StringC Text |
  BoolC Bool


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
  { functionType :: Maybe FunctionType
  , canBeNone :: Bool
  , canBeBool :: Bool
  -- , canBeType :: Bool -- want to allow more constraining on type terms
  -- For these type constraints
  , intConstraints :: Maybe [Term]
  , doubleConstraints :: Maybe [Term]
  , iriConstraints :: Maybe [Term]
  , stringConstraints :: Maybe [Term]
  , objType :: Maybe ObjectType
  , funType :: Maybe FunctionType
  , setType :: Maybe Type
  -- | The set of allowed opaque (non-constrainable) types
  , opaqueTypes :: Set OpaqueType
  , dependentTypes :: [Term]
  }

uninhabited :: Type
uninhabited =
  Type
  { functionType = Nothing
  , canBeNone = True
  , canBeBool = True
  , intConstraints = Nothing
  , doubleConstraints = Nothing
  , iriConstraints = Nothing
  , stringConstraints = Nothing
  , objType = Nothing
  , funType = Nothing
  , setType = Nothing
  , opaqueTypes = mempty
  , dependentTypes = []
  }

instance Default Type where
  def = uninhabited

intT :: Type
intT = def { intConstraints = Just [] }

doubleT :: Type
doubleT = def { doubleConstraints = Just [] }

data ObjectType = ObjectType
  { attrs :: [Attr]
  , constraints :: [Term]
  }

instance Default ObjectType where
  def =
    ObjectType
    { attrs = []
    , constraints = []
    }

data FunctionType = FunctionType
  { args :: ObjectType
  , ret :: Type
  }

-- | Represents an opaque (probably platform-specific) type
newtype OpaqueType = OpaqueType Word64
  deriving (Eq, Ord)

data Attr where
  LocalAttr :: Text -> Type -> Attr
  GlobalAttr :: IRI -> Type -> Attr

binIntFnTy :: Type
binIntFnTy = def
  { funType =
    Just FunctionType
    { args = def { attrs = [LocalAttr "x" intT, LocalAttr "y" intT] }
    , ret = intT
    }
  }

checkType :: Env -> Term -> Type -> Bool
checkType env term typ@Type{..}
  | LiteralT None <- term, canBeNone = True
  | LiteralT (BoolC _) <- term, canBeBool = True
  | LiteralT (IntC _) <- term, Just [] <- intConstraints = True
  | LiteralT (DoubleC _) <- term, Just [] <- doubleConstraints = True
  | LiteralT (StringC _) <- term, Just [] <- stringConstraints = True
  | LiteralT (IRI_C _) <- term, Just [] <- iriConstraints = True
  -- | VarT name <- term =
  --     case lookup name env of
  --       Just (term',typ') -> checkType env term' typ
  --       Nothing -> False
      
  -- | TypeT _ <- term, canBeType = True

intConstraintToSmt :: MonadError Text m => Term -> m (SMT.SMTExpr Bool)
intConstraintToSmt (LiteralT (IntC x)) = pure $ SMT.IntegerConst x

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
