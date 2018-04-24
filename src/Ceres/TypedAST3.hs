{-# LANGUAGE TypeInType #-}
module Ceres.TypedAST2 where

import XRN.Prelude

data Expr =
  LitE Literal |
  VarE Text |
  AppE Expr |
  ObjE ObjExpr |
  FunE Expr Expr Expr |
  TypeE Expr

data ObjExpr = ObjExpr
  { bindings :: [(Text, Expr)]
  }

data Literal =
  IntL Integer |
  DoubleL Double |
  StringL Text |
  BoolL Bool |
  None

data ObjType = ObjType
  { attrs :: [(Text, Expr)]
  , constraints :: [Expr]
  }

data Typ =
  NoneT |
  BoolT |
  IntT |
  DoubleT |
  StringT |
  Fun ObjType Type |
  ObjT ObjType |
  UnionT [Typ] |
  TypeT
