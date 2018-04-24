{-# LANGUAGE TypeInType #-}
module Ceres.TypedAST2 where

import XRN.Prelude
import qualified Ceres.RawAST as Raw
import Text.Megaparsec (SourcePos)
import GHC.TypeLits

data SymTyp = (:>) Symbol Typ

type Env = [SymTyp]

data Expr (env :: Env) (t :: Typ) where
  ConstExpr :: SourcePos -> PrimVal t -> Expr env t

data ArgList (env :: Env) where
  ALNil :: ArgList '[]
  ALCons :: ArgList env -> proxy (name :: Symbol) -> Expr env 'TypT -> ArgList ((name ':> 'TypT) ': env)

data PrimVal t where
  DoubleC :: Double -> PrimVal 'NumT
  StringC :: Text -> PrimVal 'StrT

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
