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
}
