package ceres.lang

import ceres.data.avlMapOf
import ceres.lang.ast.EvalEnv

val BaseEvalEnv = EvalEnv(avlMapOf(
    "+." to wrap({x:Double, y:Double -> x + y}),
    "-." to wrap({x:Double, y:Double -> x - y}),
    "*." to wrap({x:Double, y:Double -> x * y}),
    "/." to wrap({x:Double, y:Double -> x / y}),
    "==" to wrap({x:Any?, y:Any? -> x == y}),
    "!=" to wrap({x:Any?, y:Any? -> x != y}),
    "<." to wrap({x:Double, y:Double -> x < y}),
    ">." to wrap({x:Double, y:Double -> x > y}),
    "<=." to wrap({x:Double, y:Double -> x <= y}),
    ">=." to wrap({x:Double, y:Double -> x >= y}),
    "not" to wrap({x:Boolean -> !x})
))
