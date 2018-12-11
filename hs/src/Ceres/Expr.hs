module Ceres.Expr where

import XRN.Prelude
import Text.Megaparsec (SourcePos)

data SourceLoc = SourceLoc { commentWsStart:: SourcePos, start :: SourcePos, end :: SourcePos }
  deriving (Eq, Show, Generic)

data Expr where
  LitE :: SourceLoc -> Expr
  FunCall :: Expr -> [Expr] -> SourceLoc -> Expr
  FunExpr :: SourceLoc -> Expr
  CaseExpr :: SourceLoc -> Expr
  VarExpr :: Text -> Expr
  PropAccessExpr :: Expr -> Text -> SourceLoc -> Expr
  AndExpr :: Expr -> Expr -> SourceLoc -> Expr
  OrExpr :: Expr -> Expr -> SourceLoc -> Expr
  NotExpr :: Expr -> SourceLoc -> Expr

data TypeT where
  NilType :: TypeT
  StringType :: TypeT
  BoolType :: TypeT
  IntegerType :: TypeT
  DoubleType :: TypeT
  ObjType :: (HashMap Text Type) -> Expr -> TypeT
  SetType :: TypeT -> TypeT
  ListType :: TypeT -> TypeT
  VectorType :: TypeT -> TypeT
  FunType :: [(Text, Expr)] -> Expr -> TypeT
  RefinementType :: TypeT -> Expr -> Text -> TypeT
  UnionType :: [TypeT] -> TypeT
  TypeType :: TypeT

data SMTExpr =
  Sym Text |
  List [SMTExpr]

data Env = Env
  { bindings :: HashMap Text Checked
  -- , smtEngine
  , smtAssertions :: [SMTExpr]
  }

data Checked = Checked
  { typ :: TypeT
  , value :: Maybe ()
  -- , cost
  , smtEncoding :: Maybe SMTExpr
  }

type TypeErrors = [(Text, Maybe Expr)]

type CheckResult = Either TypeErrors Checked

data EvalLevel = None | Eval | SMTEncode

class Monad m => MonadMeter m where
  consumeGas :: Word32 -> m ()

class Monad m => MonadTrackUsage m where
  trackUsage :: Text -> Expr -> m ()

typeCheck :: (MonadReader Env m, MonadMeter m, MonadTrackUsage m) => Expr -> EvalLevel -> m CheckResult
typeCheck = error "TODO"
