module Ceres.Codegen.Javascript where

import XRN.Prelude hiding ((<>))
import Control.Monad.State (State, get, modify)
import Data.Text.Prettyprint.Doc
import Text.Megaparsec (unPos)
import SourceMap.Types
import Ceres.Core
import Ceres.Loc

-- term :: Monad m => m Doc
-- term

type PP = Identity

data JSAnnotation = JSAnnotation
  { loc :: Loc
  }

type JSDoc = Doc JSAnnotation

ppTerm :: Term -> PP JSDoc
ppTerm (LiteralT loc lit) = annotateLoc loc <$> ppLiteral lit
-- -- ppTerm (TypeT _) = error "TODO pp Type"
ppTerm (VarT loc (IRIName (IRI name))) = pure $ annotateLoc loc $ pretty name
ppTerm (ObjT loc (ObjTerm bindings)) = do
  bindings' <- mapM ppObjBinding bindings
  pure $ annotateLoc loc $ encloseSep lbracket rbracket comma bindings'
-- ppTerm (TypeT _) = error "TODO pp Type"

ppObjBinding :: ObjBinding -> PP JSDoc
ppObjBinding (ObjBinding loc name _ value) = do
  value' <- ppTerm value
  pure $ annotateLoc loc $ pretty name <> colon <+> value'

ppLiteral :: Literal -> PP JSDoc
ppLiteral None = "none"
ppLiteral (IRI_C (IRI x)) = pure $ angles (pretty x)
ppLiteral (IntC x) =
  if x >= -9007199254740991 || x <= 9007199254740991
  then pure $ "Int.fromDouble" <> parens (pretty x)
  else pure $ "Int.fromStrihg" <> parens (dquotes (pretty x))
ppLiteral (DoubleC x) = pure $ pretty x
ppLiteral (StringC x) = pure $ dquotes $ pretty x
  -- where
  --   escape "\n" = "\\n"
  --   escape c = c
ppLiteral (BoolC True) = "true"
ppLiteral (BoolC False) = "false"

annotateLoc :: Loc -> JSDoc -> JSDoc
annotateLoc loc doc = annotate JSAnnotation {loc} doc

genSourceMappings :: SimpleDocStream JSAnnotation -> State Pos [Mapping]
genSourceMappings (SChar _ stream) = do
  modify $ \pos -> pos {posColumn = posColumn pos + 1}
  genSourceMappings stream
genSourceMappings (SText len _ stream) = do
  modify $ \pos -> pos {posColumn = posColumn pos + (fromInteger (toInteger len))}
  genSourceMappings stream
genSourceMappings (SLine idnt stream) = do
  modify $ \pos -> pos {posLine = posLine pos + 1, posColumn = (fromInteger (toInteger idnt))}
  genSourceMappings stream
genSourceMappings (SAnnPush JSAnnotation{..} stream) = do
  pos <- get
  mappings <- genSourceMappings stream
  case locToMapping loc pos of
    Just m -> pure $ m : mappings
    _ -> pure mappings
  where
    locToMapping (SourceLoc SourcePos{..}) genPos =
      Just Mapping
      { mapGenerated = genPos
      , mapOriginal = Just Pos {posLine = intConv sourceLine, posColumn = intConv sourceColumn}
      , mapSourceFile = Just sourceName
      , mapName = Nothing
      }
    locToMapping _ _ = Nothing
    intConv = fromInteger . toInteger . unPos
genSourceMappings _ = pure []
