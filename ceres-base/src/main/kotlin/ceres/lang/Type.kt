package ceres.lang

import ceres.lang.smtlib.RExpr
import kotlin.reflect.KClass

sealed class Type<out T>

sealed class ValidationResult<out T> {
    data class Success<T>(val value: T): ValidationResult<T>()
    data class Error<T>(val path: String, val err: String)
}

typealias ValidationErrors = List<ValidationError>

data class ValidationError(val error: String, val path: List<String> = emptyList())

fun error(error: String): ValidationErrors =
    listOf(ValidationError(error))

val noErrors: ValidationErrors = emptyList()


sealed class PropertyType<T>: Type<T>() {
    abstract fun validate(value: Any): ValidationErrors
}

sealed class DataType<T : Any>(): PropertyType<T>() {
    abstract val type: KClass<T>
}

data class EntityType(
    val properties: List<Property<Any?>>,
    val constraints: List<Refinement>): PropertyType<Any>() {
    override fun validate(value: Any): ValidationErrors {
        if(value !is Entity)
            return error("type.Entity")

        return properties.fold(noErrors) { errs, prop ->
            errs.plus(prop.validate(value.get(prop.name)))
        }
    }
}

sealed class Property<T> {
    abstract fun validate(value: Any?): ValidationErrors

    abstract val name: String
}

data class OneProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>
): Property<T>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value == null)
            TODO()
        return type.validate(value)
    }
}

data class ZeroOrOneProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>
): Property<T?>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value == null)
            return noErrors
        return type.validate(value)
    }
}

sealed class SetProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>,
    val minCount: Int? = null,
    val maxCount: Int? = null
): Property<Set<T>>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value !is Set<*>)
            TODO()
        // TODO check counts
        return value.fold(noErrors, { errs, x ->
            if(x == null) errs.plus(error("nullElement"))
            else errs.plus(type.validate(x))
        })
    }
}

sealed class ListProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>,
    val minCount: Int? = null,
    val maxCount: Int? = null
): Property<List<T>>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value !is List<*>)
            TODO()
        TODO()
    }
}

sealed class NumericType<T : Any>: DataType<T>() {
    abstract val minValue: T?
    abstract val maxValue: T?
}

data class IntegerType(override val minValue: Integer? = null, override val maxValue: Integer? = null): NumericType<Integer>() {
    override fun validate(value: Any): ValidationErrors {
        if(value !is Integer)
            TODO()
        return noErrors
    }

    override val type: KClass<Integer>
        get() = Integer::class
}

data class DoubleType(override val minValue: Double? = null, override val maxValue: Double? = null): NumericType<Double>() {
    override fun validate(value: Any): ValidationErrors {
        if(value !is Double)
            TODO()
        if(minValue != null && value < minValue)
            TODO()
        if(maxValue != null && value > maxValue)
            TODO()
        return noErrors
    }

    override val type: KClass<Double>
        get() = Double::class
}

data class StringType(val regex: Regex?): DataType<String>() {
    override fun validate(value: Any): ValidationErrors {
        if(value !is String)
            TODO()
        if(regex != null && !regex.matches(value))
            TODO()
        return noErrors
    }

    override val type: KClass<String>
        get() = String::class
}

data class BoolType(
    /** Used to restrict value in refinement contexts */ val value: Boolean?
): DataType<Boolean>() {
    override fun validate(v: Any): ValidationErrors {
        if(v !is Boolean)
            TODO()
        if(value != null && v != value)
            TODO()
        return noErrors
    }

    override val type: KClass<Boolean>
        get() = Boolean::class
}

object NilType: Type<Nil>()

//object DateType: DataType()
//
//object TimeType: DataType()
//
//object DataTimeType: DataType()

data class FunctionType(
    val params: List<Pair<String, Type<Any>>>,
    val ret: Type<Any>,
    val suspend: Boolean = false,
    val terminationConstraint: TerminationConstraint = TerminationConstraint.Partial
// TODO allowedPrimitives
// TODO contextParams
): Type<Function<Any>>()

sealed class TerminationConstraint {
    object Total: TerminationConstraint()
    //TODO data class Bounded(): TerminationConstraint()
    object Partial: TerminationConstraint()
}

data class SetType<T>(val elemType: Type<T>): Type<Set<T>>()

data class ListType<T>(val elemType: Type<T>): Type<List<T>>()

//data class OpaquePlatformType(val id: String): Type ()

typealias Refinement = (varname: String) -> RExpr<Boolean>

