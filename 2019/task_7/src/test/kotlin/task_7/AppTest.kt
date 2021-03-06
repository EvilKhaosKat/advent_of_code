package task_7

import kotlin.test.Test
import kotlin.test.assertEquals

class AppTest {
    @Test
    fun testIntcodeComputer() {
        val input = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31," +
                "1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104," +
                "999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

        assertEquals(1, IntcodeComputer("3,9,8,9,10,9,4,9,99,-1,8").addInput(8).solve(true))
        assertEquals(1, IntcodeComputer("3,9,7,9,10,9,4,9,99,-1,8").addInput(7).solve(true))
        assertEquals(1, IntcodeComputer("3,3,1108,-1,8,3,4,3,99").addInput(8).solve(true))
        assertEquals(1, IntcodeComputer("3,3,1107,-1,8,3,4,3,99").addInput(7).solve(true))

        assertEquals(999, IntcodeComputer(input).addInput(1).solve())
        assertEquals(1000, IntcodeComputer(input).addInput(8).solve())
        assertEquals(1001, IntcodeComputer(input).addInput(42).solve())
    }

    @Test
    fun testFirstPart() {
        val app = App()

        assertEquals(43210, app.solveFirst("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"))
        assertEquals(54321, app.solveFirst("3,23,3,24,1002,24,10,24,1002,23,-1,23," +
                "101,5,23,23,1,24,23,23,4,23,99,0,0"))
    }

    @Test
    fun testSecondPart() {
        val app = App()

        assertEquals(139629729, app.solveSecond(
                "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26," +
                        "27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
        ))
        assertEquals(18216, app.solveSecond(
                "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54," +
                        "-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4," +
                        "53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
        ))
    }
}
