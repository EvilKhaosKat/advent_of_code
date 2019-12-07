package task_7

import kotlin.test.Test
import kotlin.test.assertEquals

class AppTest {
    @Test fun testAppHasAGreeting() {
        val classUnderTest = App()
    }

    @Test
    fun testIntcodeComputer() {
        val input = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31," +
                "1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104," +
                "999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
        val intcodeComputer = IntcodeComputer()

        assertEquals(999, intcodeComputer.solve(input, mutableListOf(1)))
        assertEquals(1000, intcodeComputer.solve(input, mutableListOf(8)))
        assertEquals(1001, intcodeComputer.solve(input, mutableListOf(42)))
    }
}
