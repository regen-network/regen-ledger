module Ceres.Value where

import XRN.Prelude

data Value =
  None |
  IntV Integer |
  DoubleV Double |
  BoolV Bool |
  StringV Text |
  IRI_V Text |
  ObjV (HashMap Text Value) |
  SetV (Set Value)
  deriving (Show, Eq)
