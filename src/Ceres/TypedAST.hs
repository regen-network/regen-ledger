module Ceres.TypedAST where

import XRN.Prelude
import Text.Megaparsec (SourcePos)
import GHC.TypeLits

data Const (k :: Nat)
data (:+:) a b
data (:*:) a b
data Fun a b k

data (:.) a b

data Env classes props locals

data ExprList env t n cost where
  ENil :: ExprList env t 0 (Const 0)
  ECons :: Expr env t k -> ExprList env t n l -> ExprList env t (n + 1) (k :+: l :+: 1)

data Expr env t cost where
  PrimFnRef :: SourcePos -> PrimFn a b k -> Expr env (Fun a b k) (Const 0)
  ConstExpr :: SourcePos -> PrimVal t -> Expr env t (Const 0)
  -- Apply ordered and named function application
  -- Lambda
  -- Bind :: proxy symbol -> Expr env t k -> Expr (symbol :. env) t k -- type level unique names
  -- Ref name in env
  -- ObjExpr
  VecExpr :: ExprList env t n k -> Expr env (Vec t n) k
  -- PropRef :: SourcePos -> Expr env (Record cls props) -> proxy prop -> Expr 
  -- CaseExpr

data PrimFn a b (cost :: Type) where
  Add :: PrimFn (Double, Double) Double (Const 0)
  Map :: PrimFn (Fun a b k, Vec a n) (Vec b n) (k :*: n)
  Filter :: PrimFn (Fun a Bool k, Vec a n) (Vec a n) (k :*: n)
  Reduce :: PrimFn (a, Fun (a, b) a k, Vec b n) a (k :*: n)

data Vec t n

data PrimVal t where
  DoubleC :: Double -> PrimVal Double
  StringC :: Text -> PrimVal Text

data Record (cls :: Symbol) (prop :: [Symbol])

-- data ObjType = ObjType
--   { attrs :: [Attr]
--   }

-- data Attr = Attr
--   { name :: Text
--   , typ :: Typ
--   }

-- data Typ =
--   ObjTypeT ObjType |
--   FunTypeT ObjType Typ |
--   NumT |
--   BoolT |
--   StrT |
--   NatT

data ObjType = ObjType
  { attrs :: [Attr]
  -- constraints ::
  } deriving (Show, Eq)

data Attr = Attr
  { name :: Text
  , typ :: Typ
  } deriving (Show, Eq)

data Typ =
  ObjTypeT ObjType |
  FunTypeT ObjType Typ |
  UnionT [Typ] |
  LabelT Text |
  TypT |
  NumT |
  BoolT |
  StrT |
  NatT
  deriving (Show, Eq)
