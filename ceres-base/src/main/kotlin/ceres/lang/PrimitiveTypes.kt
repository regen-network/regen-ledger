package ceres.lang

object Nil

expect class Integer {
    fun add(y: Integer): Integer
    fun subtract(y: Integer): Integer
    fun multiply(y: Integer): Integer
    fun divide(y: Integer): Integer
    fun remainder(y: Integer): Integer
    operator fun compareTo(other: Integer): Int
//    fun divMod(y: Integer): Pair<Integer, Integer>
}

// TODO expect or implement:
expect class Decimal constructor(unscaledVal: Integer, scale: Int) {
    fun add(y: Decimal): Decimal
    fun subtract(y: Decimal): Decimal
    fun multiply(y: Decimal): Decimal
    fun divide(y: Decimal): Decimal
    fun remainder(y: Decimal): Decimal
    operator fun compareTo(other: Decimal): Int
//    fun divMod(y: Decimal): Pair<Decimal, Decimal>
}

//expect class Date
//
//expect class Time
//
//expect class DateTime

