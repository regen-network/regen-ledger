package ceres.data

sealed class Either<E, R>
data class Success<E, R>(val result: R): Either<E, R>()
data class Failure<E, R>(val error: E): Either<E, R>()
