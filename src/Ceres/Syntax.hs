module Ceres.Syntax where

import XRN.Prelude
import Text.Megaparsec (SourcePos)
import qualified Text.Megaparsec as P

newtype Identifier = Identifier Text
  deriving (Eq, Show, Generic)

data Range = Range { start :: SourcePos, end :: SourcePos }
  deriving (Eq, Show, Generic)

data CommentWS = CommentWS [CommentWSElem] Range
  deriving (Eq, Show, Generic)

data CommentWSElem =
  LineComment Text |
  BlockComment Text |
  Comma |
  Semi
  deriving (Eq, Show, Generic)

type ParseError = P.ParseError Char CeresParseError

data CeresParseError = CPE_TODO
  deriving (Eq, Ord, Show, Generic)

instance P.ShowErrorComponent CeresParseError where
  showErrorComponent = show

type Module = ObjExpr

data Expr =
  LitE CommentWS Literal Range |
  IdentE CommentWS Identifier Range |
  ObjE CommentWS ObjExpr Range |
  TypeE CommentWS TypeExpr Range |
  ParseErrorE ParseError
  deriving (Eq, Show, Generic)

data TypeExpr = TE_TODO
  deriving (Eq, Show, Generic)

data Literal =
  NumLit Double |
  IntLit Integer |
  StrLit Text |
  BoolLit Bool |
  NoneLit
  deriving (Eq, Show, Generic)

data ObjExpr = ObjExpr
  { fields :: [FieldExpr]
  , trailingCommentWS :: CommentWS
  -- , objCopyId :: Maybe Identifier
  }
  deriving (Eq, Show, Generic)

-- data FieldExpr =
--   FieldExpr
--     CommentWS
--     Identifier
--     CommentWS CommentWS
--     (Maybe TypeExpr)
--     CommentWS CommentWS
--     (Maybe Expr)
--     Range |
--   FunExpr CommentWS Identifier ObjExpr TypeExpr (Maybe Expr) Range
--   deriving (Eq, Show, Generic)

data FieldExpr = FieldExpr
  { wsA :: CommentWS
  , name :: Identifier
  , wsB :: CommentWS
  , wsC :: CommentWS
  , typ :: Maybe TypeExpr
  , wsD :: CommentWS
  , wsE :: CommentWS
  , value :: Expr
  , wsF :: CommentWS
  , range :: Range
  }
  deriving (Eq, Show, Generic)
