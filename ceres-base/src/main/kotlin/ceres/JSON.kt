package ceres

sealed class JSONValue

data class JSONObject(val map: Map<String, JSONValue>): Map<String, JSONValue> by map

data class JSONArray(val list: List<JSONValue>): List<JSONValue> by list

data class JSONNumber(val value: Number)

data class JSONString(val value: String)

data class JSONBoolean(val value: Boolean)

object JSONNull
