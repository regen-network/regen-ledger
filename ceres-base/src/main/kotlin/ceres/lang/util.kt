package ceres.lang.util

interface Source {
    val uri: String?
}

data class SourceLoc(val source: Source, val start: Int, val end: Int)

interface HasSourceLoc {
    val sourceLoc: SourceLoc?
}

