package ceres.lang

import ceres.JSONValue
import ceres.geo.Point
import ceres.geo.Polygon
import ceres.lang.smtlib.RExpr
import kotlin.reflect.KClass

sealed class Type

sealed class PropertyType<T : Any>: Type() {
    abstract val kClass: KClass<T>
    abstract fun validate(value: Any): ValidationErrors
    //abstract fun fromJson(value: JSONValue): T?
    //abstract fun toJSON(value: T): JSONValue
}

data class EntityType(
    val properties: List<Property<Any?>>,
    val constraints: List<Refinement>): PropertyType<Entity>() {
    override val kClass: KClass<Entity>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        if(value !is Entity)
            return error("type.Entity")

        return properties.fold(noErrors) { errs, prop ->
            errs.plus(prop.validate(value.get(prop.name)))
        }
    }
}

data class DisjointEntityUnion(val entityTypes: List<EntityType>): PropertyType<Entity>() {
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

data class OneProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null
): Property<T>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value == null)
            TODO()
        return type.validate(value)
    }
}

data class ZeroOrOneProperty<T: Any>(
    override val name: String,
    val type: PropertyType<T>,
    override val iri: String? = null
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
    override val iri: String? = null,
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
    override val iri: String? = null,
    val minCount: Int? = null,
    val maxCount: Int? = null
): Property<List<T>>() {
    override fun validate(value: Any?): ValidationErrors {
        if(value !is List<*>)
            TODO()
        TODO()
    }
}

sealed class DataType<T : Any>(): PropertyType<T>() {
    abstract val iri: String
}

sealed class NumericType<T : Any>: DataType<T>() {
    abstract val minValue: T?
    abstract val maxValue: T?
    abstract val multipleOf: T?
}

data class IntegerType(override val minValue: Integer? = null, override val maxValue: Integer? = null,
                       override val multipleOf: Integer? = null
): NumericType<Integer>() {
    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#integer"

    override fun validate(value: Any): ValidationErrors {
        if(value !is Integer)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Integer>
        get() = Integer::class
}

data class DoubleType(
    override val minValue: Double? = null,
    override val maxValue: Double? = null,
    override val multipleOf: Double? = null,
    val exclusiveMin: Boolean = false,
    val exclusiveMax: Boolean = false
): NumericType<Double>() {
    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#double"

    override fun validate(value: Any): ValidationErrors {
        if(value !is Double)
            TODO()
        if(minValue != null && value < minValue)
            TODO()
        if(maxValue != null && value > maxValue)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Double>
        get() = Double::class
}

data class StringType(val minLength: Int? = null, val maxLength: Int? = null, val regex: Regex?): DataType<String>() {
    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#string"

    override fun validate(value: Any): ValidationErrors {
        if(value !is String)
            TODO()
        if(regex != null && !regex.matches(value))
            TODO()
        return noErrors
    }

    override val kClass: KClass<String>
        get() = String::class
}

data class EnumValue(val value: String, val iri: String?)

data class EnumType(val values: Set<EnumValue>): PropertyType<String>() {
    override val kClass: KClass<String>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

// TODO
object IRIType: PropertyType<String>() {
    override val kClass: KClass<String>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.

    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }
}

data class BoolType(
    /** Used to restrict value in refinement contexts */ val value: Boolean?
): DataType<Boolean>() {
    override val iri: String
        get() = "http://www.w3.org/2001/XMLSchema#boolean"

    override fun validate(v: Any): ValidationErrors {
        if(v !is Boolean)
            TODO()
        if(value != null && v != value)
            TODO()
        return noErrors
    }

    override val kClass: KClass<Boolean>
        get() = Boolean::class
}

object NilType: Type()

//object DateType: DataType()
//object TimeType: DataType()
//object DateTimeType: DataType()
//object DurationType: DataType()

object PointType: DataType<Point>() {
    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Point>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
    override val iri: String
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
}

object PolygonType: DataType<Polygon>() {
    override fun validate(value: Any): ValidationErrors {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    override val kClass: KClass<Polygon>
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
    override val iri: String
        get() = TODO("not implemented") //To change initializer of created properties use File | Settings | File Templates.
}

data class FunctionType(
    val params: List<Pair<String, Type>>,
    val ret: Type,
    val suspend: Boolean = false,
    val terminationConstraint: TerminationConstraint = TerminationConstraint.Partial
// TODO allowedPrimitives
// TODO contextParams
): Type()

sealed class TerminationConstraint {
    object Total: TerminationConstraint()
    //TODO data class Bounded(): TerminationConstraint()
    object Partial: TerminationConstraint()
}

// NOTE:
// SetType and ListType are not property types because we want to handle cardinality of Entity properties differently

data class SetType<T>(val elemType: Type): Type()

data class ListType<T>(val elemType: Type): Type()

//data class OpaquePlatformType(val id: String): Type ()

typealias ValidationErrors = List<ValidationError>

data class ValidationError(val error: String, val path: List<String> = emptyList())

fun error(error: String): ValidationErrors =
    listOf(ValidationError(error))

val noErrors: ValidationErrors = emptyList()

typealias Refinement = (varname: String) -> RExpr<Boolean>

