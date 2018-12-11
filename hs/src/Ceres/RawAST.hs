module Ceres.RawAST where

import XRN.Prelude
import Text.Megaparsec (SourcePos)

-- data Module = Module [Decl]
--   deriving (Show)

-- data Decl =
--   DataDecl Text [AttrRef] |
--   FunDecl
--   deriving (Show)

-- data Expr =
--   Add Expr Expr |
--   Sub Expr Expr |
--   Mul Expr Expr |
--   Div Expr Expr

-- data AttrRef = AttrRef
--   {
--   } deriving (Show)

newtype Identifier = Identifier Text
  deriving (Show)

data Expr =
  LitE SourcePos Literal |
  IdentE SourcePos Identifier |
  -- TypeE SourcePos TypeExpr |
  ObjE SourcePos ObjExpr |
  BinOpE BinOp Expr Expr
  deriving (Show)

data BinOp = Add | Sub | Mult | Div | Or | And | Eq | Gte | Gt | Lte | Lt
  deriving (Show)


data Literal =
  NumLit Double |
  IntLit Integer |
  BoolLit Bool |
  StrLit Text
  deriving (Show)

data ObjExpr = ObjExpr
  { fields :: [FieldExpr]
  -- , objCopyId :: Maybe Identifier
  }
  deriving (Show)

data FieldExpr =
  FieldExpr SourcePos Identifier (Maybe TypeExpr) (Maybe Expr) |
  FunExpr SourcePos Identifier ObjExpr TypeExpr (Maybe Expr)
  deriving (Show)
  -- FieldAssignExpr SourcePos Identifier (Maybe TypeExpr) Expr |
  -- FieldCopyExpr SourcePos Identifier
  -- deriving (Show)

data TypeExpr = TypeExpr Identifier
  -- PrimTE SourcePos PrimType
  deriving (Show)

data PrimType = NumT | BoolT | StrT | NatT
  deriving (Show)

data ObjTypeDef = ObjTypeDef
  { fields :: [FieldDef] }

data FieldDef =
  FieldDef Identifier (Maybe TypeExpr) (Maybe Expr)
