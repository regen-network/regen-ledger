package ceres.schema

sealed class Property() {
    abstract val name: String

    data class DataProperty(
        override val name: String,
        val datatype: Datatype): Property()

    data class ObjectProperty(
        override val name: String,
        val cls: Class): Property()
}

data class Class(
    val name: String,
    val properties: ClassPropertyRef
)

data class ClassPropertyRef(
    val prop: Property,
    val required: Boolean
)

data class Datatype(
    val name: String
)
