package ceres.serialization

import ceres.util.reduceAcc

// TODO optimize this
class ByteArrayBuilder {
    private val output: MutableList<ByteArray> = mutableListOf()

    fun write(bytes: ByteArray): ByteArrayBuilder {
        output.add(bytes)
        return this
    }

    fun finish(): ByteArray {
        val len = output.reduceAcc(0, {len, arr -> len + arr.size})
        val bytes = ByteArray(len)
        var i = 0
        for(arr in output) {
            for(x in arr) {
                bytes[i] = x
                ++i
            }
        }
        return bytes
    }
}

inline fun Byte.toBytes() = byteArrayOf(this)

fun ByteArrayBuilder.write(x: Byte): ByteArrayBuilder =
    write(x.toBytes())

fun ByteArray.unsafeWrite(x: Short, idx: Int = 0): ByteArray {
    var y = x.toInt()
    for(i in 0..1) {
        this[i + idx] = y.toByte()
        y = y.ushr(8)
    }
    return this
}

fun Short.toBytes(): ByteArray = ByteArray(2).unsafeWrite(this)

fun ByteArrayBuilder.write(x: Short): ByteArrayBuilder =
        write(x.toBytes())

fun ByteArray.unsafeWrite(x: Int, idx: Int = 0): ByteArray {
    var y = x
    for(i in 0..3) {
        this[i + idx] = y.toByte()
        y = y.ushr(8)
    }
    return this
}

fun Int.toBytes(): ByteArray = ByteArray(4).unsafeWrite(this)

fun ByteArrayBuilder.write(x: Int): ByteArrayBuilder =
        write(x.toBytes())

fun ByteArray.unsafeWrite(x: Long, idx: Int = 0): ByteArray {
    var y = x
    for(i in 0..7) {
        this[i + idx] = y.toByte()
        y = y.ushr(8)
    }
    return this
}
fun Long.toBytes(): ByteArray = ByteArray(8).unsafeWrite(this)

fun ByteArrayBuilder.write(x: Long): ByteArrayBuilder =
        write(x.toBytes())

fun ByteArrayBuilder.write(x: Double): ByteArrayBuilder =
        write(x.toRawBits())

fun ByteArrayBuilder.write(x: Boolean): ByteArrayBuilder =
        write(byteArrayOf(if(x) 1 else 0))

//fun ByteArrayBuilder.writeLengthPrefixed(x: String): ByteArrayBuilder {
//    write(x.length)
//    write(x.toByteArray())
//}

