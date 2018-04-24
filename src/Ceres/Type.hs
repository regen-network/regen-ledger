module Ceres.Type where

import XRN.Prelude

data Attr = Attr
  { name :: Text
  , typ :: Typ
  }

data ObjConstraint = ObjConstraint
  { attrs :: [Attr]
  -- TODO assertions
  }

data Typ = Typ
  { objType :: Maybe ObjConstraint
  , canBeNone :: Bool
  , canBeString :: Bool
  , stringConstraint Maybe (StringConstraint Bool)
  , canBeIRI :: Bool
  , iriConstraint Maybe (IRIConstraint Bool)
  , canBeInt :: Bool
  , intConstraint :: Maybe (IntConstraint Bool)
  , canBeDouble :: Bool
  , doubleConstraint :: Maybe (DoubleConstraint Bool)
  , setType :: Maybe Typ
  -- TODO dependent terms
  }

data IntConstraint a where
  VarI :: Text -> IntConstraint Integer
  ConstI :: Integer -> IntConstraint Integer
  AddI :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Integer
  SubI :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Integer
  MultI :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Integer
  GT_I :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Bool
  GTE_I :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Bool
  LT_I :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Bool
  LTE_I :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Bool
  Eq_I :: IntConstraint Integer -> IntConstraint Integer -> IntConstraint Bool
  And_I :: IntConstraint Bool -> IntConstraint Bool -> IntConstraint Bool
  Or_I :: IntConstraint Bool -> IntConstraint Bool -> IntConstraint Bool

data DoubleConstraint a where
  VarD :: Text -> DoubleConstraint Double
  ConstD :: Double -> DoubleConstraint Double
  AddD :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Double
  SubD :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Double
  MultD :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Double
  GT_D :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Bool
  GTE_D :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Bool
  LT_D :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Bool
  LTE_D :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Bool
  Eq_D :: DoubleConstraint Double -> DoubleConstraint Double -> DoubleConstraint Bool
  And_D :: DoubleConstraint Bool -> DoubleConstraint Bool -> DoubleConstraint Bool
  Or_D :: DoubleConstraint Bool -> DoubleConstraint Bool -> DoubleConstraint Bool

data IRIConstraint a where
  And_I :: IRIConstraint Bool -> IRIConstraint Bool -> IRIConstraint Bool
  Or_I :: IRIConstraint Bool -> IRIConstraint Bool -> IRIConstraint Bool
  Eq_I :: Text -> IRIConstraint Bool
  RefersToType :: Typ -> IRIConstraint Bool

data StringConstraint a where
  Eq_T :: Text -> StringConstraint Bool
  MinLen_T :: Int -> StringConstraint Bool
  MaxLen_T :: Int -> StringConstraint Bool
  -- Regex_T :: Text -> StringConstraint Bool
  And_T :: StringConstraint Bool -> StringConstraint Bool -> StringConstraint Bool
  Or_T :: StringConstraint Bool -> StringConstraint Bool -> StringConstraint Bool
