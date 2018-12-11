module Ceres.Parser where

import XRN.Prelude hiding (try)
import Ceres.AST
import Text.Megaparsec (Parsec, try, between, getPosition, optional, takeWhile1P)
import qualified Text.Megaparsec as P
import Text.Megaparsec.Char (space1, letterChar, alphaNumChar)
import qualified Text.Megaparsec.Char as P
import Text.Megaparsec.Expr (makeExprParser, Operator(..))
import qualified Text.Megaparsec.Char.Lexer as L
import qualified Data.Text as T
import qualified Data.Text.IO as T

type Parser = Parsec CeresParseError Text

-- blockCmnt :: Parser [Char]
-- blockCmnt = p >> (P.manyTill e n)
--   where
--     e = blockCmnt <|> void P.anySingle
--     p = P.string "/*"
--     n = P.string "*/"

-- scn :: Parser ()
-- scn = L.space space1 lineCmnt blockCmnt

-- sc :: Parser ()
-- sc = L.space (void $ takeWhile1P Nothing f) lineCmnt blockCmnt
--   where
--     f x = x == ' ' || x == '\t'

-- lexeme :: Parser a -> Parser a
-- lexeme = L.lexeme scn

-- symbol :: Text -> Parser Text
-- symbol = L.symbol scn

-- parens :: Parser a -> Parser a
-- parens = between (symbol "(") (symbol ")")

-- braces :: Parser a -> Parser a
-- braces = between (symbol "{") (symbol "}")

-- brackets :: Parser a -> Parser a
-- brackets = between (symbol "[") (symbol "]")

identifier :: Parser Identifier
identifier = Identifier . T.pack <$> try p
  where
    p       = (:) <$> letterChar <*> many alphaNumChar

wrapRange :: Parser (Range -> a) -> Parser a
wrapRange p = do
  start <- getPosition
  f <- p
  end <- getPosition
  pure $ f Range {start, end}

commentWS :: Parser CommentWS
commentWS = wrapRange $ CommentWS . catMaybes <$> many commentWSElem
  where
    commentWSElem = P.choice [ws, lineCmnt]
    lineCmnt =
      Just . LineComment <$> (P.string "//" *> (P.takeWhileP (Just "character") (/= '\n')))
    ws = P.space1 *> pure Nothing

expr :: Parser Expr
expr = P.withRecovery (\e -> pure $ ParseErrorE e) expr'
  -- try opParser <|>
  where
    expr' =
      wrapRange (LitE <$> commentWS <*> literal) <|>
      wrapRange (ObjE <$> commentWS <*> objExpr) <|>
      wrapRange (IdentE <$> commentWS <*> identifier)
--   where
--     opParser = makeExprParser expr
--       [ [binary "*" Mult, binary "/" Div] 
--       , [binary "+" Add, binary "-" Sub]
--       ]
--     binary sym op = InfixL $ (\x y -> BinOpE op x y) <$ symbol sym

literal :: Parser Literal
literal = try numLit <|> try decLit <|> boolLit <|> hexLit
  where
    numLit = NumLit <$> L.signed (pure ()) L.float
    decLit = IntLit <$> L.signed (pure ()) L.decimal
    hexLit = IntLit <$> L.signed (pure ()) L.hexadecimal
    boolLit = BoolLit <$> boolLit'
    boolLit' = P.string "true" *> pure True <|>
               P.string "false" *> pure False
    -- TODO strLit = StrLit <$> 

objExpr :: Parser ObjExpr
objExpr = between (P.string "{") (P.string "}") objExpr'

objExpr' :: Parser ObjExpr
objExpr' = ObjExpr <$> P.sepEndBy fieldExpr (P.char ';' <|> P.char ',') <*> commentWS

fieldExpr :: Parser FieldExpr
fieldExpr = -- try funExpr <|>
  fieldExpr'
  where
    -- funExpr =
    --   FunExpr
    --   <$> getPosition
    --   <*> (identifier <* scn)
    --   <*> (parens objExpr' <* scn)
    --   <*> (symbol ":" *> typeExpr <* scn)
    --   <*> optional (symbol "=" *> expr <* scn)
    fieldExpr' = wrapRange $
      (FieldExpr
       <$> commentWS
       <*> identifier
       <*> commentWS
       <*> (optional (P.string ":") *> commentWS)
       <*> pure Nothing
       <*> commentWS
       <*> (P.string "=" *> commentWS)
       <*> expr
       <*> commentWS)

-- typeExpr :: Parser TypeExpr
-- typeExpr = TypeExpr <$> identifier
--   -- PrimTE <$> getPosition <*> primType

-- primType :: Parser PrimType
-- primType =
--   symbol "num" *> pure NumT <|>
--   symbol "bool" *> pure BoolT <|>
--   symbol "str" *> pure StrT <|>
--   symbol "nat" *> pure NatT

-- objTypeDef :: Parser ObjTypeDef
-- objTypeDef = error "TODO"

ceresModule :: Parser ObjExpr
ceresModule = objExpr' <* P.eof

test1 :: IO ()
test1 = do
  txt <- T.readFile "test/Test1.ceres"
  P.parseTest ceresModule txt

-- test2 = P.parseTest expr "1"
