module Ceres.Diagnostics where

import XRN.Prelude
import Ceres.Syntax
import Language.Haskell.LSP.TH.DataTypesJSON (Diagnostic(..), DiagnosticSeverity(..))
import qualified Language.Haskell.LSP.TH.DataTypesJSON as LSP
import qualified Text.Megaparsec as P
import qualified Data.List.NonEmpty as NonEmpty

exprDiagnostics :: Expr -> [Diagnostic]
exprDiagnostics (LitE _ _ _) = []
exprDiagnostics (IdentE _ _ _) = []
exprDiagnostics (ObjE _ objExpr _) = objExprDiagnostics objExpr
exprDiagnostics (TypeE _ _ _) = []
exprDiagnostics (ParseErrorE err) = [parseErrorDiagnostic err]

parseErrorDiagnostic :: ParseError -> Diagnostic
parseErrorDiagnostic err@(P.TrivialError sp _ _) = parseErrorDiagnostic' err sp
parseErrorDiagnostic err@(P.FancyError sp _) = parseErrorDiagnostic' err sp

parseErrorDiagnostic' :: ParseError -> NonEmpty.NonEmpty P.SourcePos -> Diagnostic
parseErrorDiagnostic' err sp =
  let pos = sourcePosToPosition $ NonEmpty.head sp
  in Diagnostic
  { _range = LSP.Range pos pos
  , _severity = Just DsError
  , _code = Nothing
  , _source = Nothing
  , _message = cs $ P.parseErrorPretty err
  }

sourcePosToPosition :: P.SourcePos -> LSP.Position
sourcePosToPosition sourcePos =
  LSP.Position { LSP._line = P.unPos $ P.sourceLine sourcePos
               , LSP._character = P.unPos $ P.sourceColumn sourcePos }

objExprDiagnostics :: ObjExpr -> [Diagnostic]
objExprDiagnostics ObjExpr {..} =
  concatMap fieldExprDiagnostics fields

fieldExprDiagnostics :: FieldExpr -> [Diagnostic]
fieldExprDiagnostics FieldExpr {..} =
  exprDiagnostics value

