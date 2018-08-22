package ceres.geo

interface Geometry

data class Point(
    val lat: Double,
    val lon: Double
) : Geometry

data class Polygon(
    val exteriorRing: List<Point>,
    val interiorRings: List<List<Point>>
) : Geometry
