module XRN.Prelude
  (module ClassyPrelude
  ,module X)
where

import ClassyPrelude
import Data.Kind as X (Type)
import Data.Proxy as X (Proxy(..))
import GHC.TypeLits as X (Nat, Symbol)
import Data.Default as X
import Control.Monad.Except as X (MonadError (..), ExceptT)
