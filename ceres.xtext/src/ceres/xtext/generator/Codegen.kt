package ceres.xtext.generator

import ceres.xtext.ceres.*

fun Decl.emitDecl(): String =
    when (this) {
        is FunDecl -> this.emitFunDecl()
        is ValDecl -> this.emitValDecl()
        else -> {
            ""
        }
    }

fun ValDecl.emitValDecl(): String =
    "val $name = ${expr.emitExpr()}"

fun FunDecl.emitFunDecl(): String =
    "fun $name() {}"

fun Expr.emitExpr(): String =
    when (this) {
        is IntLit -> this.emitIntLit()
        is StrLit -> this.emitStrLit()
        is VarRef -> this.emitVarRef()
        else -> {
            ""
        }
    }

fun IntLit.emitIntLit(): String = this.value.toString()

fun StrLit.emitStrLit(): String = this.value

fun VarRef.emitVarRef(): String =
        when(value) {
            is ValDecl -> (value as ValDecl).name
            else -> ""
        }
