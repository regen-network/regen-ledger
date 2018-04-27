{-# OPTIONS_GHC -fno-warn-orphans #-}
module Ceres.Loc
  (Loc(..)
  ,SourcePos(..))
where

import XRN.Prelude
import Text.Megaparsec (SourcePos(..), unPos)

data Loc =
  SourceLoc SourcePos |
  Internal
  deriving (Eq, Generic)

instance Hashable SourcePos where
  hashWithSalt salt SourcePos{..} =
   salt `hashWithSalt`
   sourceName `hashWithSalt`
   (unPos sourceLine) `hashWithSalt`
   (unPos sourceColumn)

instance Hashable Loc
