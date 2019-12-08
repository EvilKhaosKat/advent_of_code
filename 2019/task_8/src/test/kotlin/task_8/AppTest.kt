package task_8

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull

class AppTest {
    @Test
    fun testPart1() {
        val app = App(3, 2)
        assertEquals(1, app.solveFirst(listOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2)))
    }

    @Test
    fun testPart2() {
        val app = App(2, 2)
        assertEquals(listOf(0, 1, 1, 0), app.solveSecond(listOf(0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0)))
    }

}
