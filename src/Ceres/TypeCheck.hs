-- https://github.com/tomprimozic/type-systems
module Ceres.TypeCheck where

import XRN.Prelude
import Ceres.Term

type TypedEnv = HashMap Text (Term, Term)

data Attr = Attr
  { name :: Text
  , typ :: Typ
  }

data ObjType = ObjType
  { attrs :: [Attr]
  -- TODO assertions
  }

data Typ = Typ
  { objType :: Maybe ObjType
  , canBeNone :: Bool
  , canBeString :: Bool
  , canBeIRI :: Bool
  , canBeInt :: Bool
  , intRefinement :: Maybe (IntRefinement Bool)
  , canBeDouble :: Bool
  , doubleRefinement :: Maybe (DoubleRefinement Bool)
  , setType :: Maybe Typ
  -- TODO dependent terms
  }

typeUnion :: Typ -> Typ -> Typ
typeUnion x y =
  Typ { objType = error "TODO"
      , canBeNone = canBeNone x || canBeNone y
      , canBeString = canBeString x || canBeString y
      , canBeIRI = canBeIRI x || canBeIRI y
      , canBeInt = canBeInt x || canBeInt y
      , canBeDouble = canBeDouble x || canBeDouble y
      , intRefinement = error "TODO"
      , doubleRefinement = error "TODO"
      , setType = error "TODO"
      }

typeIntersection :: Typ -> Typ -> Typ
typeIntersection x y =
  Typ { objType = error "TODO"
      , canBeNone = canBeNone x && canBeNone y
      , canBeString = canBeString x && canBeString y
      , canBeIRI = canBeIRI x && canBeIRI y
      , canBeInt = canBeInt x && canBeInt y
      , canBeDouble = canBeDouble x && canBeDouble y
      , intRefinement = error "TODO"
      , doubleRefinement = error "TODO"
      , setType = error "TODO"
      }

isInhabited :: Typ -> Bool
isInhabited x = canBeNone x || canBeDouble x || canBeInt x || canBeIRI x || canBeIRI x || isJust (objType x) || isJust (setType x)

isUninhabited :: Typ -> Bool
isUninhabited = not . isInhabited

isSubtype :: Typ -> Typ -> Bool
isSubtype child parent =
  checkSimple canBeNone &&
  checkSimple canBeInt &&
  checkSimple canBeDouble &&
  checkSimple canBeIRI &&
  checkSimple canBeString
  where
    checkSimple f = not (f child) || f parent

isSubObjType :: ObjType -> ObjType -> Bool
isSubObjType child parent = False

data IntRefinement a where
  VarI :: Text -> IntRefinement Integer
  ConstI :: Integer -> IntRefinement Integer
  AddI :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Integer
  SubI :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Integer
  MultI :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Integer
  GT_I :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Bool
  GTE_I :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Bool
  LT_I :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Bool
  LTE_I :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Bool
  Eq_I :: IntRefinement Integer -> IntRefinement Integer -> IntRefinement Bool
  And_I :: IntRefinement Bool -> IntRefinement Bool -> IntRefinement Bool
  Or_I :: IntRefinement Bool -> IntRefinement Bool -> IntRefinement Bool

data DoubleRefinement a where
  VarD :: Text -> DoubleRefinement Double
  ConstD :: Double -> DoubleRefinement Double
  AddD :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Double
  SubD :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Double
  MultD :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Double
  GT_D :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Bool
  GTE_D :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Bool
  LT_D :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Bool
  LTE_D :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Bool
  Eq_D :: DoubleRefinement Double -> DoubleRefinement Double -> DoubleRefinement Bool
  And_D :: DoubleRefinement Bool -> DoubleRefinement Bool -> DoubleRefinement Bool
  Or_D :: DoubleRefinement Bool -> DoubleRefinement Bool -> DoubleRefinement Bool

data Type =
  IntT (Maybe (IntRefinement Bool)) |
  DoubleT (Maybe (DoubleRefinement Bool)) |
  BoolT |
  StringT 

-- checkTypeTerm :: Env -> Term -> Bool
-- checkTypeTerm _ (ConstT (IRI_C iri)) = _
-- checkTypeTerm _ (ConstT _) = False

{-
dsc = decentralized schema

dsc:ceres/Int
dsc:ceres/Double
dsc:ceres/Bool
dsc:ceres/None
dsc:ceres/String

dsc:ceres/Attr
dsc:ceres/name
dsc:ceres/type


functions
dsc:ceres:1.0/args
dsc:ceres:1.0/ret
dsc:ceres:1.0/exp

object types
dsc:ceres:1.0/properties 
-}

{-
module Ceres.Core

label Int
label Double
label Bool
label None

Any{it:any} = True

-- type Type = {it:Any} => Bool

attr name : String
attr type : Type

type Attr = {name, type}

attr attrs : Set Attr

type ObjType = {attrs}

type Type = ObjType | 'Int | 'Double | 'Bool | 'Set

Entity(t:ObjType): ObjType = {
  attrs = t.attrs.map(a =>
    match a with
    | ObjType -> a {type = Ref a.type}
    | _ -> a
  )
}

type Store = {
  saveEntity(t:ObjType, obj: Entity t): IRI
}

-}

{-
type Bounds = {
  low: Num
  high: Num
  assert(low < high)
}

attr soilPh : Num & it >= 0 & it <= 14

type SoilTestResults = { soilPh }


-}

