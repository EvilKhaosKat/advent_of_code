package task_10

import java.io.File
import java.util.*

class App {
    fun solveFirst(input: List<String>): Int {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    fun solveSecond(input: List<String>): Int {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

    val greeting: String
        get() {
            return "Hello world."
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
            val visibleAsteroids = getVisibleAsteroids(asteroid)

            if (visibleAsteroids.size > bestVisibleAsteroids.size) {
                bestPosition = asteroid
                bestVisibleAsteroids = visibleAsteroids
            }
        }

        return Pair(bestPosition, bestVisibleAsteroids)
    }

    private fun getVisibleAsteroids(asteroid: Point): List<Point> {
        var visibleAsteroids = mutableListOf<Point>()

        val pointsToCheck = ArrayDeque<Point>()
        pointsToCheck.addAll(getPoints(asteroid, 1))

        while (pointsToCheck.size > 0) {
            val point = pointsToCheck.pop()
            if (map[point.x][point.y].isAsteroid()) {
                if (!hiddenByVisibleAsteroids(visibleAsteroids, asteroid, point)) {
                    visibleAsteroids.add(point)
                }
            }
        }

        return visibleAsteroids
    }

    private fun hiddenByVisibleAsteroids(visibleAsteroids: MutableList<Point>, fromPoint: Point, targetPoint: Point): Boolean {
        for (visibleAsteroid in visibleAsteroids) {
            if (isOnLine(visibleAsteroid, fromPoint, targetPoint)) {
                return true
            }
        }

        return false
    }

    private fun getPoints(asteroid: Point, radius: Int): List<Point> {
        val points = mutableListOf<Point>()

        val sideSize = 1 + radius * 2

        //up to 12 o'clock
        var currX = asteroid.x
        var currY = asteroid.y - radius

        //right top corner
        var targetX = currX + sideSize / 2
        var targetY = currY

        for (x in currX until targetX) {
            if (correctCoordinate(x, currY)) {
                points.add(Point(x, currY))
            }
        }
        currX = targetX
        currY = targetY

        //right bottom corner
        targetX = currX
        targetY = currY + radius

        for (y in currY downTo targetY) {
            if (correctCoordinate(currX, y)) {
                points.add(Point(currX, y))
            }
        }
        currX = targetX
        currY = targetY

        //left bottom corner
        targetX = currX - radius
        targetY = currY

        for (x in currX downTo  targetX) {
            if (correctCoordinate(x, currY)) {
                points.add(Point(x, currY))
            }
        }
        currX = targetX
        currY = targetY


        //left top corner
        targetX = currX
        targetY = currY - radius

        for (y in currY downTo targetY) {
            if (correctCoordinate(currX, y)) {
                points.add(Point(currX, y))
            }
        }
        currX = targetX
        currY = targetY


        //till start point
        targetX = currX + sideSize / 2
        targetY = currY

        for (x in currX until targetX) {
            if (correctCoordinate(x, currY)) {
                points.add(Point(x, currY))
            }
        }

        //(1+r*2)

        //1 -> 3 + 3 + 1 + 1 = 8
        //...
        //. .
        //...

        //2 -> 5 + 5 + 3 + 3 = 16
        //!!!!!
        //!...!
        //!. .!
        //!...!
        //!!!!!

        return points
    }

    private fun correctCoordinate(x: Int, y: Int): Boolean {
        if (x < 0 || y < 0) {
            return false
        }

        if (x >= map.size || y >= map[0].size) {
            return false
        }

        return true
    }

    fun isOnLine(point: Point, lineStart: Point, lineEnd: Point): Boolean {
        return (lineStart.y - lineEnd.y) * point.x +
                (lineEnd.x - lineStart.x) * point.y +
                (lineStart.x * lineEnd.y - lineEnd.x * lineStart.y) == 0
    }

    fun Char.getFieldValue(): Int {
        return when (this) {
            '.' -> 0
            '#' -> 1
            else -> throw RuntimeException("unknown value of $this")
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