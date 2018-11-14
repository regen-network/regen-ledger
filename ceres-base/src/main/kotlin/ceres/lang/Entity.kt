package ceres.lang

interface Entity {
    fun get(prop: String): Any?
}

interface PersistentEntity: Entity {
    fun set(prop: String, value: Any?): PersistentEntity
}
