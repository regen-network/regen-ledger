package ceres.lang.smtlib

import ceres.cache.Cache
import ceres.lang.Integer
import ceres.lang.Decimal

// TODO maybe deprecate Real/Decimal support?

sealed class SMTExpr {
    data class SList(val xs: List<SMTExpr>): SMTExpr()
    data class Sym(val value: String): SMTExpr()
    data class IntL(val value: String): SMTExpr()
    data class RealL(val value: String): SMTExpr()
    data class StrL(val value: String): SMTExpr()
}

fun sym(x: String) = SMTExpr.Sym(x)

fun list(vararg xs: SMTExpr) = SMTExpr.SList(xs.toList())


/**
 * A refinement typing expression
 */
sealed class RExpr<T> {
    override fun toString(): String =
        when(this) {
            is DeclareConst -> "(declare-fun $name () $sort)"
            is Assert -> "(assert $x)"
            is Var -> name
            is BoolE -> x.toString()
            is Not -> "(not $x)"
            is Implies -> "(=> $x $y)"
            is And -> "(and $x $y)"
            is Or -> "(or $x $y)"
            is Xor -> "(xor $x $y)"
            is Eq<*> -> "(= $x $y)"
            is Distinct<*> -> "(distinct $x $y)"
            is Ite -> "(ite $x $y $z)"
            is IntegerE -> x.toString()
            is DecimalE -> x.toString()
            is Neg -> "(- $x)"
            is Add -> "(+ $x $y)"
            is Subtract -> "(- $x $y)"
            is Mult -> "(* $x $y)"
            is DecimalDiv -> "(/ $x $y)"
            is IntDiv -> "(div $x $y)"
            is IntMod -> "(mod $x $y)"
            is IntAbs -> "(abs $x)"
            is LTE<*> -> "(<= $x $y)"
            is LT<*> -> "(< $x $y)"
            is GT<*> -> "(> $x $y)"
            is GTE<*> -> "(>= $x $y)"
            is ToReal -> "(to_real $x)"
            is IoInt -> "(to_int $x)"
            is IsInt -> "(is_int $x)"
        }
}

// Basics
data class DeclareConst(val name: String, val sort: Sort): RExpr<Unit>()
data class Assert(val x: RExpr<Boolean>): RExpr<Unit>()
data class Var<T>(val name: String): RExpr<T>()

// Core theory: http://smtlib.cs.uiowa.edu/theories-Core.shtml
data class BoolE(val x: Boolean): RExpr<Boolean>()
data class Not(val x: RExpr<Boolean>): RExpr<Boolean>()
data class Implies(val x: RExpr<Boolean>, val y: RExpr<Boolean>): RExpr<Boolean>()
data class And(val x: RExpr<Boolean>, val y: RExpr<Boolean>): RExpr<Boolean>()
data class Or(val x: RExpr<Boolean>, val y: RExpr<Boolean>): RExpr<Boolean>()
data class Xor(val x: RExpr<Boolean>, val y: RExpr<Boolean>): RExpr<Boolean>()
data class Eq<T>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class Distinct<T>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class Ite<T>(val x: RExpr<Boolean>, val y: RExpr<T>, val z: RExpr<T>): RExpr<T>()

// Ints_Decimals theory: http://smtlib.cs.uiowa.edu/theories-Decimals_Ints.shtml
data class IntegerE(val x: Integer): RExpr<Integer>()
data class DecimalE(val x: Decimal): RExpr<Decimal>()
data class Neg<T: Number>(val x: RExpr<T>): RExpr<T>()
data class Add<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<T>()
data class Subtract<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<T>()
data class Mult<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<T>()
data class DecimalDiv(val x: RExpr<Decimal>, val y: RExpr<Decimal>): RExpr<Decimal>()
data class IntDiv(val x: RExpr<Integer>, val y: RExpr<Integer>): RExpr<Integer>()
data class IntMod(val x: RExpr<Integer>, val y: RExpr<Integer>): RExpr<Integer>()
data class IntAbs(val x: RExpr<Integer>): RExpr<Integer>()
data class LTE<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class LT<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class GT<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class GTE<T: Number>(val x: RExpr<T>, val y: RExpr<T>): RExpr<Boolean>()
data class ToReal(val x: RExpr<Integer>): RExpr<Decimal>()
data class IoInt(val x: RExpr<Decimal>): RExpr<Integer>()
data class IsInt(val x: RExpr<Decimal>): RExpr<Boolean>()

// TODO: FloatingPoint theory: http://smtlib.cs.uiowa.edu/theories-FloatingPoint.shtml
//data class DoubleE(val x: Double): RExpr<Double>()
// TODO: FixedSizeBitVectors: http://smtlib.cs.uiowa.edu/theories-FixedSizeBitVectors.shtml
// TODO: ArraysEx: http://smtlib.cs.uiowa.edu/theories-ArraysEx.shtml

sealed class Sort {
    override fun toString(): String =
        when(this) {
            Decimal -> "Decimal"
            Int -> "Int"
            Bool -> "Bool"
        }
    object Decimal: Sort()
    object Int: Sort()
    object Bool: Sort()
}

enum class CheckSat {Sat, Unsat, Unknown }

interface SmtEngine {
    fun checkSat(ctx: SmtContext): CheckSat
}

class CachingSmtEngine(val engine: SmtEngine): SmtEngine {
    val cache: Cache<SmtContext, CheckSat> = TODO()

    // TODO just pass in list of assertions and env of vars and let this derive the context
    override fun checkSat(ctx: SmtContext): CheckSat {
        val hit = cache[ctx]
        if(hit != null)
            return hit
        val sat = engine.checkSat(ctx)
        cache.put(ctx, sat)
        return sat
    }
}

data class SmtContext(val txt: String) {
    fun add(vararg statements: RExpr<Unit>) =
            SmtContext(txt + "\n" + statements.joinToString("\n"))
}

//interface SmtEngine {
//    fun exec(str: String): String
//}
//
//class SmtContext(val engine: SmtEngine) {
//    private var curStack : MutableList<RExpr<Unit>> = mutableListOf()
//    private var statementStacks: MutableList<MutableList<RExpr<Unit>>> = mutableListOf()
//
//    init {
//        engine.exec("push")
//    }
//
//    fun add(statement: RExpr<Unit>) {
//        engine.exec(statement.toString())
//        curStack.add(statement)
//    }
//
//    fun push() {
//        engine.exec("(push)")
//        statementStacks.add(curStack)
//        curStack = mutableListOf()
//    }
//
//    fun pop() {
//        engine.exec("(pop)")
//        val n = statementStacks.size
//        if(n > 0) {
//            curStack = statementStacks[n - 1]
//            statementStacks.removeAt(n - 1)
//        } else {
//            engine.exec("(push)")
//            curStack = mutableListOf()
//        }
//    }
//
//    fun reset() {
//        engine.exec("(reset)")
//        engine.exec("(push)")
//        curStack = mutableListOf()
//        statementStacks = mutableListOf()
//    }
//
//    fun checkSat(): Boolean =
//        when(engine.exec("(check-sat)")) {
//            "sat" -> true
//            "unsat" -> false
//            else -> false
//        }
//}
