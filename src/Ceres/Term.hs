module Ceres.Term where

import XRN.Prelude

data Term =
  ConstT Const |
  ObjT ObjTerm |
  SetT [Term] |
  AppT Term ObjTerm |
  CaseT [(Term, Term)] (Maybe Term) |
  OpaqueT Word64
  -- LabelT Text ?

data Const =
  None |
  IRI_C IRI |
  IntC Integer |
  DoubleC Double |
  StringC Text |
  BoolC Bool

newtype IRI = IRI Text

data ObjTerm = ObjTerm [ObjBinding]

data ObjBinding = ObjBinding
  { name :: Text
  , typ :: Maybe Term
  , value :: Term
  }

type Env = HashMap Text Term

-- hasType :: Env -> Term -> Term -> Maybe Text
-- hasType _ 
