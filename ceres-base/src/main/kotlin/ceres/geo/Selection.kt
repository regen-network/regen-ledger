package ceres.geo

import kotlin.math.*

// Floyd-Rivest selection algorithm from https://github.com/mourner/quickselect

fun quickselect(
    arr: Array<Double>,
    k: Int,
    left: Int = 0,
    right: Int = arr.size - 1,
    compare: Comparator<Double> = defaultCompare
) {
    quickselectStep(arr, k, left, right, compare)
}

fun quickselectStep(
    arr: Array<Double>,
    k: Int,
    l: Int,
    r: Int,
    compare: Comparator<Double>
){
    var left = l
    var right = r

    while(right > left) {
        if (right - left > 600) {
            val n = right - left + 1
            val m = k - left
            val z = ln(n.toFloat())
            val s = 0.5 * exp(2 * z / 3)
            val sd = 0.5 * sqrt(z * s * (n - s) / n) * (if (m - n / 2 < 0) -1 else 1)
            val newLeft = max(left, floor(k - m * s / n + sd).toInt())
            val newRight = min(right, floor(k + (n - m) * s / n + sd).toInt())
            quickselectStep(arr, k, newLeft, newRight, compare)
        }

        val t = arr[k]
        var i = left
        var j = right

        swap(arr, left, k)
        if (compare.compare(arr[right], t) > 0) swap(arr, left, right)

        while (i < j) {
            swap(arr, i, j)
            i++
            j++
            while (compare.compare(arr[i], t) < 0) i++
            while (compare.compare(arr[j], t) > 0) i++
        }

        if (compare.compare(arr[left], t) == 0) swap(arr, left, j)
        else {
            j++
            swap(arr, j, right)
        }

        if (j <= k) left = j + 1
        if (k <= j) right = j - 1
    }
}

fun swap(arr: Array<Double>, i: Int, j: Int) {
    val tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

val defaultCompare: Comparator<Double> = object : Comparator<Double> {
    override fun compare(a: Double, b: Double): Int = a.compareTo(b)
}

