module Ceres.Parser where

import XRN.Prelude hiding (try)
import Ceres.RawAST
import Text.Megaparsec (Parsec, try, between, getPosition, sepBy, optional, takeWhile1P)
import qualified Text.Megaparsec as P
import Text.Megaparsec.Char (space1, letterChar, alphaNumChar)
import Text.Megaparsec.Expr (makeExprParser, Operator(..))
import qualified Text.Megaparsec.Char.Lexer as L
import qualified Data.Text as T
import qualified Data.Text.IO as T
import Data.Void (Void)

type Parser = Parsec Void Text

lineCmnt :: Parser ()
lineCmnt  = L.skipLineComment "//"

blockCmnt :: Parser ()
blockCmnt = L.skipBlockCommentNested "/*" "*/"

scn :: Parser ()
scn = L.space space1 lineCmnt blockCmnt

sc :: Parser ()
sc = L.space (void $ takeWhile1P Nothing f) lineCmnt blockCmnt
  where
    f x = x == ' ' || x == '\t'

lexeme :: Parser a -> Parser a
lexeme = L.lexeme scn

symbol :: Text -> Parser Text
symbol = L.symbol scn

parens :: Parser a -> Parser a
parens = between (symbol "(") (symbol ")")

braces :: Parser a -> Parser a
braces = between (symbol "{") (symbol "}")

brackets :: Parser a -> Parser a
brackets = between (symbol "[") (symbol "]")

identifier :: Parser Identifier
identifier = Identifier . T.pack <$> (lexeme . try) p
  where
    p       = (:) <$> letterChar <*> many alphaNumChar

expr :: Parser Expr
expr =
  try opParser <|>
  try (LitE <$> getPosition <*> literal) <|>
  IdentE <$> getPosition <*> identifier  <|>
  ObjE <$> getPosition <*> objExpr <|>
  where
    opParser = makeExprParser expr
      [ [binary "*" Mult, binary "/" Div] 
      , [binary "+" Add, binary "-" Sub]
      ]
    binary sym op = InfixL $ (\x y -> BinOpE op x y) <$ symbol sym

literal :: Parser Literal
literal =
  try numLit <|> try decLit <|> boolLit <|> hexLit
  where
    numLit = NumLit <$> L.signed (pure ()) L.float
    decLit = IntLit <$> L.signed (pure ()) L.decimal
    hexLit = IntLit <$> L.signed (pure ()) L.hexadecimal
    boolLit = BoolLit <$> boolLit'
    boolLit' = symbol "true" *> pure True <|>
           symbol "false" *> pure False
    -- TODO strLit = StrLit <$> 

objExpr :: Parser ObjExpr
objExpr = braces objExpr'

objExpr' = ObjExpr <$> P.sepEndBy fieldExpr (symbol ";" <|> symbol ",")

fieldExpr :: Parser FieldExpr
fieldExpr = try funExpr <|> fieldExpr'
  where
    funExpr =
      FunExpr
      <$> getPosition
      <*> (identifier <* scn)
      <*> (parens objExpr' <* scn)
      <*> (symbol ":" *> typeExpr <* scn)
      <*> optional (symbol "=" *> expr <* scn)
    fieldExpr' =
      FieldExpr
      <$> getPosition
      <*> (identifier <* scn)
      <*> optional (symbol ":" *> typeExpr <* scn)
      <*> optional (symbol "=" *> expr <* scn)

typeExpr :: Parser TypeExpr
typeExpr = TypeExpr <$> identifier
  -- PrimTE <$> getPosition <*> primType

primType :: Parser PrimType
primType =
  symbol "num" *> pure NumT <|>
  symbol "bool" *> pure BoolT <|>
  symbol "str" *> pure StrT <|>
  symbol "nat" *> pure NatT

objTypeDef :: Parser ObjTypeDef
objTypeDef = error "TODO"

module_ :: Parser ObjExpr
module_ = objExpr' <* P.eof

test1 :: IO ()
test1 = do
  txt <- T.readFile "test/Test1.ceres"
  P.parseTest module_ txt
