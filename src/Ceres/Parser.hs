module Ceres.Parser where

import XRN.Prelude
import Ceres.RawAST
import Text.Megaparsec (Parsec, between, getPosition, sepBy, optional)
import Text.Megaparsec.Char (space1, letterChar, alphaNumChar)
import Text.Megaparsec.Expr ()
import qualified Text.Megaparsec.Char.Lexer as L

type Parser = Parsec () Text

sc :: Parser ()
sc = L.space space1 lineCmnt blockCmnt
  where
    lineCmnt  = L.skipLineComment "//"
    blockCmnt = L.skipBlockCommentNested "/*" "*/"

lexeme :: Parser a -> Parser a
lexeme = L.lexeme sc

symbol :: Text -> Parser Text
symbol = L.symbol sc

parens :: Parser a -> Parser a
parens = between (symbol "(") (symbol ")")

braces :: Parser a -> Parser a
braces = between (symbol "{") (symbol "}")

brackets :: Parser a -> Parser a
brackets = between (symbol "[") (symbol "]")

-- identifier :: Parser Text
-- identifier = (lexeme . try) (p >>= check)
--   where
--     p       = (:) <$> letterChar <*> many alphaNumChar
--     check x = if x `elem` rws
--                 then fail $ "keyword " ++ show x ++ " cannot be an identifier"
--                 else return x

identifier :: Parser Identifier
identifier = error "TODO"

expr :: Parser Expr
expr =
  LitE <$> getPosition <*> literal <|>
  IdentE <$> getPosition <*> identifier <|>
  ObjE <$> getPosition <*> objExpr

literal :: Parser Literal
literal =
  numLit <|> boolLit -- <|> strLit
  where
    numLit = NumLit <$> L.signed (pure ()) L.float
    boolLit = BoolLit <$> bool
    bool = symbol "true" *> pure True <|>
           symbol "false" *> pure False
    -- TODO strLit = StrLit <$> 

objExpr :: Parser ObjExpr
objExpr = braces $ objExpr'
  where
    objExpr' = ObjExpr <$> sepBy fieldExpr (symbol ",") <*> pure Nothing
    fieldExpr :: Parser FieldExpr
    fieldExpr =
      FieldAssignExpr
        <$> getPosition
        <*> identifier
        <*> optional (symbol ":" *> typeExpr)
        <*> (symbol "=" *> expr)

typeExpr :: Parser TypeExpr
typeExpr =
  PrimTE <$> getPosition <*> primType

primType :: Parser PrimType
primType =
  symbol "num" *> pure NumT <|>
  symbol "bool" *> pure BoolT <|>
  symbol "str" *> pure StrT <|>
  symbol "nat" *> pure NatT

objTypeDef :: Parser ObjTypeDef
objTypeDef = error "TODO"
