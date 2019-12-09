package task_9

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
    fun testApp() {
        val app = App()

        assertEquals(99, app.solveFirst("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"))
        assertEquals(1125899906842624L, app.solveFirst("104,1125899906842624,99"))
        assertEquals(1219070632396864L, app.solveFirst("1102,34915192,34915192,7,4,7,99,0"))
    }
}
