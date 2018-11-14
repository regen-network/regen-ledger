package ceres.lang

import ceres.data.PersistentMap

interface Entity: Map<String, Any>

interface PersistentEntity: PersistentMap<String, Any>

class EntityImpl(val map: PersistentMap<String, Any>): PersistentEntity, PersistentMap<String, Any> by map
