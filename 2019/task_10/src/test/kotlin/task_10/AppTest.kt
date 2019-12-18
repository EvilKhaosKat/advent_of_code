package task_10

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull

class AppTest {
    @Test
    fun testExamples() {
        val app = App()

        assertEquals(8, app.solveFirst(listOf(
                ".#..#",
                ".....",
                "#####",
                "....#",
                "...##")))

        assertEquals(33, app.solveFirst(listOf(
                "......#.#.",
                "#..#.#....",
                "..#######.",
                ".#.#.###..",
                ".#..#.....",
                "..#....#.#",
                "#..#....#.",
                ".##.#..###",
                "##...#..#.",
                ".#....####")))
        /*
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####


         */
    }
}
