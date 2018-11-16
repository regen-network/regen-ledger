package ceres.lang

import ceres.data.avlMapOf
import ceres.lang.ast.EvalEnv
import ceres.lang.ast.Literal

private fun fnTy(retTy: Type, vararg params: Pair<String, Type>, smtIntrinsic: Boolean = false) =
    FunctionType(params.map { it.first to Literal(it.second, TypeType)}, Literal(retTy, TypeType))

private val dddFn =
    fnTy(DoubleType.default, "x" to DoubleType.default, "y" to DoubleType.default, smtIntrinsic = true)

private val ddbFn =
    fnTy(BoolType.default, "x" to DoubleType.default, "y" to DoubleType.default, smtIntrinsic = true)

private val bbbFn =
    fnTy(BoolType.default, "x" to BoolType.default, smtIntrinsic = true)

val BaseEvalEnv = EvalEnv(avlMapOf(
    "+." to wrap(dddFn, { x:Double, y:Double -> x + y}),
    "-." to wrap(dddFn, { x:Double, y:Double -> x - y}),
    "*." to wrap(dddFn, { x:Double, y:Double -> x * y}),
    "div" to wrap(dddFn, { x:Double, y:Double -> x / y}),
    "==" to wrap(bbbFn, { x:Any?, y:Any? -> x == y}),
    "!=" to wrap(bbbFn, { x:Any?, y:Any? -> x != y}),
    "<." to wrap(ddbFn, { x:Double, y:Double -> x < y}),
    ">." to wrap(ddbFn, { x:Double, y:Double -> x > y}),
    "<=." to wrap(ddbFn, { x:Double, y:Double -> x <= y}),
    ">=." to wrap(ddbFn, { x:Double, y:Double -> x >= y}),
    "not" to wrap(bbbFn, { x:Boolean -> !x})
))
