module XRN.Core where

import XRN.Prelude
import GHC.TypeLits

data Cardinality = One | Many

class IsAttr (name :: Symbol) where
  type AttrType name :: Type
  type AttrCardinality name :: Cardinality

instance IsAttr "acct:balance" where
  type AttrType "acct:balance" = Integer
  type AttrCardinality "acct:balance" = 'One

class IsMethod (name :: Symbol) where
  type MethodInput name :: Type
  type MethodOutput name :: Type
  -- methodImpl

class IsQuery (name :: Symbol) where

data Schema (methods :: [Symbol]) (queries :: [Symbol])

type XRN_API = Schema ["transfer", "track"] '[]

data ID

data Results t

data Query t where
  EAQuery :: (IsAttr attr) => ID -> proxy attr -> Query (AttrType attr)
  AVQuery :: (IsAttr attr, t ~ AttrType attr, Eq t) => proxy attr -> t -> Query (Results ID)
  AVRangeQuery :: (IsAttr attr, t ~ AttrType attr, Ord t) => proxy attr -> t -> t -> Query (Results ID)


class Monad m => MonadTx m where

main :: IO ()
main = pure ()
