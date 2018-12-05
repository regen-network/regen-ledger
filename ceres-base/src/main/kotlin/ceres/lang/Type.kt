package ceres.lang

import ceres.data.Either
import ceres.data.Success
import ceres.geo.Point
import ceres.geo.Polygon
import ceres.lang.Expr
import ceres.lang.smtlib.RExpr
import kotlin.reflect.KClass

sealed class Type {
    abstract fun __checkSubType(subType: Type): String?
    abstract fun __union(type: Type): Type
    // TODO abstract fun intersection(type: Type): Either<String, Type>
}

fun Type.checkSubType(subType: Type) =
    this.__checkSubType(subType)

fun Type.union(type: Type) =
    when(type) {
        EmptyType -> this
        else -> this.__union(type)
    }

private fun Type.subTypeError(subType: Type) = "${subType} can't be a sub-type of ${this}"

sealed class PropertyType<T : Any> : Type() {
    abstract val kClass: KClass<T>
    abstract fun validate(value: Any): ValidationErrors
    //abstract fun fromJson(value: JSONValue): T?
    //abstract fun toJSON(value: T): JSONValue
}

data class EntityType(
    val properties: Map<String, Property<Any?>>,
    val constraints: List<Refinement>
) : PropertyType<Entity>() {

    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Entity>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        if (value !is Entity)
            return error("type.Entity")

        return properties.values.fold(noErrors) { errs, prop ->
            errs.plus(prop.validate(value.get(prop.name)))
        }
    }
}

data class DisjointEntityUnion(val entityTypes: List<EntityType>) : PropertyType<Entity>() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Entity>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

sealed class Property<T> {
    abstract fun validate(value: Any?): ValidationErrors
    abstract val name: String
    abstract val iri: String?

}

data class OneProperty<T : Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null
) : Property<T>() {
    override fun validate(value: Any?): ValidationErrors {
        if (value == null)
            TODO()
        return type.validate(value)
    }
}

data class ZeroOrOneProperty<T : Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null
) : Property<T?>() {
    override fun validate(value: Any?): ValidationErrors {
        if (value == null)
            return noErrors
        return type.validate(value)
    }
}

data class SetProperty<T : Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null,
    val minCount: Int? = null,
    val maxCount: Int? = null
) : Property<Set<T>>() {
    override fun validate(value: Any?): ValidationErrors {
        if (value !is Set<*>)
            TODO()
        // TODO check counts
        return value.fold(noErrors, { errs, x ->
            if (x == null) errs + error("nullElement")
            else errs.plus(type.validate(x))
        })
    }
}

data class ListProperty<T : Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null,
    val minCount: Int? = null,
    val maxCount: Int? = null
) : Property<List<T>>() {
    override fun validate(value: Any?): ValidationErrors {
        if (value !is List<*>)
            TODO()
        TODO()
    }
}

sealed class DataType<T : Any>() : PropertyType<T>() {
    abstract val iri: String
}

sealed class NumericType<T : Any> : DataType<T>() {
    abstract val minValue: T?
    abstract val maxValue: T?
    abstract val multipleOf: T?
}

data class IntegerType(
    override val minValue: Integer? = null, override val maxValue: Integer? = null,
    override val multipleOf: Integer? = null
) : NumericType<Integer>() {
    companion object {
        val default = IntegerType()
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#integer"

    override fun validate(value: Any): ValidationErrors {
        if (value !is Integer)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Integer>
        get() = Integer::class
}

data class DecimalType(
    val scale: Int,
    override val minValue: Decimal? = null, override val maxValue: Decimal? = null,
    override val multipleOf: Decimal? = null
) : NumericType<Decimal>() {
    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val iri: String
        get() = TODO()

    override fun validate(value: Any): ValidationErrors {
        if (value !is Decimal)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Decimal>
        get() = Decimal::class
}

data class DoubleType(
    override val minValue: Double? = null,
    override val maxValue: Double? = null,
    override val multipleOf: Double? = null,
    val exclusiveMin: Boolean = false,
    val exclusiveMax: Boolean = false
) : NumericType<Double>() {
    companion object {
        val default = DoubleType()
    }

    override fun __checkSubType(subType: Type): String? =
        when(subType) {
            is DoubleType -> null // TODO refinements
            else -> subTypeError(subType)
        }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#double"

    override fun validate(value: Any): ValidationErrors {
        if (value !is Double)
            TODO()
        if (minValue != null && value < minValue)
            TODO()
        if (maxValue != null && value > maxValue)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Double>
        get() = Double::class
}

data class StringType(val minLength: Int? = null, val maxLength: Int? = null, val regex: Regex? = null) :
    DataType<String>() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#string"

    override fun validate(value: Any): ValidationErrors {
        if (value !is String)
            TODO()
        if (regex != null && !regex.matches(value))
            TODO()
        return noErrors
    }

    override val kClass: KClass<String>
        get() = String::class
}

data class EnumValue(val value: String, val iri: String?)

data class EnumType(val values: Set<EnumValue>) : PropertyType<String>() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<String>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

// TODO
//object IRIType : PropertyType<String>() {
//    override val kClass: KClass<String>
//        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
//
//    override fun validate(value: Any): ValidationErrors {
//        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
//    }
//}

data class BoolType(
    /** Used to restrict value in refinement contexts */
    val value: Boolean? = null
) : DataType<Boolean>() {
    companion object {
        val default = BoolType()
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __checkSubType(type: Type): String? =
        when (type) {
            is BoolType -> null // TODO refinement
            else -> "Expected Bool, got ${type}"
        }

    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#boolean"

    override fun validate(v: Any): ValidationErrors {
        if (v !is Boolean)
            TODO()
        if (value != null && v != value)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Boolean>
        get() = Boolean::class
}

val boolType = BoolType()

object NilType : Type() {
    override fun __checkSubType(type: Type): String? =
            when(type) {
                NilType -> null
                else -> subTypeError(type)
            }

    override fun __union(type: Type): Type =
            when(type) {
                EmptyType -> this
                NilType -> this
                is NullableType -> type
                else -> NullableType(type)
            }
}

object EmptyType: Type() {
    override fun __checkSubType(type: Type): String? = "${this} has no subtypes"

    override fun __union(type: Type): Type = type
}

object TypeType: Type() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

}

//object DateType: DataType()
//object TimeType: DataType()
//object DateTimeType: DataType()
//object DurationType: DataType()

object PointType : DataType<Point>() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Point>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
    override val iri: String
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
}

object PolygonType : DataType<Polygon>() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Polygon>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
    override val iri: String
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
}

typealias SmtEncoder = (Array<String>) -> Either<String, String>

data class FunctionType(
    val params: List<Pair<String, Expr>>, // TODO change Type to Expr
    val ret: Expr,
    val terminationConstraint: TerminationConstraint = TerminationConstraint.Partial,
    val calledPrimitives: Set<String>? = null,
    val cost: Expr? = null,
    val smtEncoder: SmtEncoder? = null,
    val suspend: Boolean = false
// TODO refinement constraints
// TODO allowedPrimitives
// TODO contextParams (a method for introducing effects)
) : Type() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

sealed class TerminationConstraint {
    object Total : TerminationConstraint()
    //TODO data class Bounded(): TerminationConstraint()
    object Partial : TerminationConstraint()
}

// NOTE:
// SetType and ListType are not property types because we want to handle cardinality of Entity properties differently

data class NullableType(val type: Type) : Type() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class SetType(val elemType: Type) : Type() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class ListType(val elemType: Type) : Type() {
    override fun __checkSubType(type: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class RefinedType(val type: Type, val refinement: Expr, val refinementVar: String): Type() {
    override fun __checkSubType(subType: Type): String? {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override fun __union(type: Type): Type {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

}

//data class OpaquePlatformType(val id: String): Type ()

typealias ValidationErrors = List<ValidationError>

data class ValidationError(val error: String, val path: List<String> = emptyList())

fun error(error: String): ValidationErrors =
    listOf(ValidationError(error))

val noErrors: ValidationErrors = emptyList()

typealias Refinement = (varname: String) -> RExpr<Boolean>

