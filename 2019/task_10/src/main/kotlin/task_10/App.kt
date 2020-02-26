package task_10

import java.io.File
import java.lang.Integer.max
import java.lang.Integer.min

class App {
    fun solveFirst(input: List<String>): Int {
        val map = Map(input)
        var (_, asteroids) = map.getBestLocation()
        return asteroids.size
    }

    fun solveSecond(input: List<String>): Int {
        return -1
    }
}

class Map(input: List<String>) {
    var map = mutableListOf<List<Int>>()
    var asteroids = mutableListOf<Point>()

    init {
        for ((y, line) in input.withIndex()) {
            val row = mutableListOf<Int>()

            for ((x, value) in line.withIndex()) {
                val fieldValue = value.getFieldValue()
                if (fieldValue.isAsteroid()) {
                    asteroids.add(Point(x, y))
                }

                row.add(fieldValue)
            }

            map.add(row)
        }
    }

    fun getBestLocation(): Pair<Point, List<Point>> {
        var bestPosition = Point(0, 0)
        var bestVisibleAsteroids = listOf<Point>()

        //TODO use channels
        for (asteroid in asteroids) {
            val visibleAsteroids = getVisibleAsteroids(asteroid, asteroids)

            if (visibleAsteroids.size > bestVisibleAsteroids.size) {
                bestPosition = asteroid
                bestVisibleAsteroids = visibleAsteroids
            }
        }

        return Pair(bestPosition, bestVisibleAsteroids)
    }

    private fun getVisibleAsteroids(asteroid: Point, asteroids: MutableList<Point>): List<Point> {
        val visibleAsteroids = mutableListOf<Point>()

        for (asteroidToCheck in asteroids) {
            if (asteroid == asteroidToCheck) {
                continue
            }

            if (hiddenByAsteroids(asteroids, asteroid, asteroidToCheck)) {
                continue
            }

            visibleAsteroids.add(asteroidToCheck)
        }

        return visibleAsteroids
    }

    private fun hiddenByAsteroids(asteroids: MutableList<Point>, fromPoint: Point, toPoint: Point): Boolean {
        for (asteroid in asteroids) {
            if (fromPoint != asteroid && toPoint != asteroid
                    && isOnSegment(asteroid, fromPoint, toPoint)) {
                return true
            }
        }

        return false
    }

    fun isOnSegment(point: Point, lineStart: Point, lineEnd: Point): Boolean {
        val isOnLine = ((lineStart.y - lineEnd.y) * point.x
                + (lineEnd.x - lineStart.x) * point.y
                + (lineStart.x * lineEnd.y - lineEnd.x * lineStart.y) == 0)

        if (!isOnLine) {
            return false
        }

        return point.x > min(lineStart.x, lineEnd.x) && point.x < max(lineStart.x, lineEnd.x)
                && point.y > min(lineStart.y, lineEnd.y) && point.y < max(lineStart.y, lineEnd.y)
    }

    fun Char.getFieldValue(): Int {
        return when (this) {
            '.' -> 0
            '#' -> 1
            else -> throw RuntimeException("unknown value of '$this'")
        }
    }

    fun Int.isAsteroid(): Boolean {
        return this == 1
    }

    data class Point(val x: Int, val y: Int)

}

fun main() {
    val app = App()
    val input = File("input.txt").readLines()

    println(app.solveFirst(input))
    println(app.solveSecond(input))
}