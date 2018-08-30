package ceres.geo

interface Geometry {
    fun boundingBox(): IBoundingBox
}

data class Point(
    val lat: Double,
    val lon: Double
) : Geometry {
    override fun boundingBox(): IBoundingBox =
        BoundingBox(lon, lat, lon, lat)

}

expect class Polygon {
    val exteriorRing: List<Point>
    val interiorRings: List<List<Point>>
}

interface ToBoundingBox {
}

interface IBoundingBox {
    val minX: Double
    val minY: Double
    val maxX: Double
    val maxY: Double
}

data class BoundingBox(
    override val minX: Double,
    override val minY: Double,
    override val maxX: Double,
    override val maxY: Double
): IBoundingBox
