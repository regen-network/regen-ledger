module Ceres.SMT where

import XRN.Prelude

data Statement

data Sort =
  IntSort |
  RealSort |
  BoolSort |
  CustomSort Text

data SMTExpr t where
  DeclareConst :: Text -> Sort -> SMTExpr Statement
  DeclareFun :: Text -> [(Text, Sort)] -> Sort -> SMTExpr Statement
  DefineFun :: Text -> [(Text, Sort)] -> Sort -> SMTExpr t -> SMTExpr Statement
  Var :: Text -> SMTExpr t
  IntegerConst :: Integer -> SMTExpr Integer
  DoubleConst :: Double -> SMTExpr Double
  Assert :: SMTExpr Bool -> SMTExpr Statement
  Or :: [SMTExpr Bool] -> SMTExpr Bool
  And :: [SMTExpr Bool] -> SMTExpr Bool
  Add :: Num a => [SMTExpr a] -> SMTExpr a
  Sub :: Num a => [SMTExpr a] -> SMTExpr a
  Mult :: Num a => [SMTExpr a] -> SMTExpr a
  Div :: Num a => [SMTExpr a] -> SMTExpr a
  Eq :: Ord a => SMTExpr a -> SMTExpr a -> SMTExpr Bool
  GT :: Ord a => SMTExpr a -> SMTExpr a -> SMTExpr Bool
  GTE :: Ord a => SMTExpr a -> SMTExpr a -> SMTExpr Bool
  LT :: Ord a => SMTExpr a -> SMTExpr a -> SMTExpr Bool
  LTE :: Ord a => SMTExpr a -> SMTExpr a -> SMTExpr Bool
  Implies :: SMTExpr Bool -> SMTExpr Bool -> SMTExpr Bool
  Forall :: Text -> Sort -> SMTExpr Bool -> SMTExpr Bool
  CheckSat :: SMTExpr Statement
  Push :: SMTExpr Statement
  Pop :: SMTExpr Statement

data SExp =
  SList [SExp] |
  SSym Text |
  SInt Integer |
  SDouble Double

instance IsString SExp where
  fromString str = error "TODO"

toSExpr :: SMTExpr t -> SExp
toSExpr (DeclareConst name sort) = SList ["declare-const", SSym name, sortToSExpr sort]

sortToSExpr :: Sort -> SExp
sortToSExpr IntSort = "Int"
sortToSExpr RealSort = "Real"
sortToSExpr BoolSort = "Bool"
sortToSExpr (CustomSort x) = SSym x
