package ceres.schema

import ceres.lang.*

fun EntityType.validateEntity(entity: Entity) {
    for (prop in properties) {
        prop.validateProperty(entity.get(prop.name))
    }
}

fun Property<*>.validateProperty(value: Any?) =
    when(this) {
        is OneProperty<*> -> this.validateOneProperty(value)
        is ZeroOrOneProperty<*> -> this.validateZeroOrOneProperty(value)
        is ManyProperty<*> -> this.validateManyProperty(value)
    }

fun OneProperty<*>.validateOneProperty(value: Any?) {
    if (value == null)
        TODO()
    when(type) {
        EntityType
    }

}

fun ZeroOrOneProperty<*>.validateZeroOrOneProperty(value: Any?) {

}

fun ManyProperty<*>.validateManyProperty(value: Any?) {

}


