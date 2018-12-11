module Ceres.Compile where

import XRN.Prelude
import Language.Haskell.LSP.TH.DataTypesJSON (Uri(..), Diagnostic(..))
import qualified Language.Haskell.LSP.TH.DataTypesJSON as LSP
import qualified Text.Megaparsec as P
import Ceres.AST
import Ceres.Parser (ceresModule)
import Ceres.Resolve
import Ceres.TypeCheck
import Ceres.Diagnostics

data CompilerState = CompilerState
  {
  }

data CompilerOptions = CompilerOptions
  {
  }

data CompilationResult = CompilationResult
  { diagnostics :: [Diagnostic]
  , ast :: Maybe Module
  }

compileOne :: CompilerOptions -> CompilerState -> Uri -> Text -> CompilationResult
compileOne _ _ uri src =
  let filePath = fromMaybe "" (LSP.uriToFilePath uri)
      parseRes = P.parse ceresModule filePath src
  in case parseRes of
        Right objExpr ->
          let resolved = resolve objExpr
              typeChecked = typeCheck resolved
          in CompilationResult
          { diagnostics = objExprDiagnostics objExpr
          , ast = Just typeChecked
          }
        Left err ->
          CompilationResult
          { diagnostics = [parseErrorDiagnostic err]
          , ast =  Nothing
          }
