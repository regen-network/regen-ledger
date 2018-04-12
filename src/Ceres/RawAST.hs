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

data Expr =
  LitE SourcePos Literal |
  IdentE SourcePos Identifier |
  TypeE SourcePos TypeExpr |
  ObjE SourcePos ObjExpr

data Literal =
  NumLit Double |
  BoolLit Bool |
  StrLit Text

data ObjExpr = ObjExpr
  { fields :: [FieldExpr]
  , objCopyId :: Maybe Identifier
  }

data FieldExpr =
  FieldAssignExpr SourcePos Identifier (Maybe TypeExpr) Expr |
  FieldCopyExpr SourcePos Identifier

data TypeExpr =
  PrimTE SourcePos PrimType

data PrimType = NumT | BoolT | StrT | NatT

data ObjTypeDef = ObjTypeDef
  { fields :: [FieldDef] }

data FieldDef =
  FieldDef Identifier (Maybe TypeExpr) (Maybe Expr)
