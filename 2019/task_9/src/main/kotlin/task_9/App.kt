/*
 * This Kotlin source file was generated by the Gradle 'init' task.
 */
package task_9

import java.io.File
import java.lang.RuntimeException

//Lets try simple hardcoded solution
class App {
    fun solveFirst(input: String): Int {
        val computer = IntcodeComputer(input)

        return -1
    }

    fun solveSecond(input: String): Int {
        return -1
    }

    private fun isCorrectPhaseSettings(aPhase: Int, bPhase: Int, cPhase: Int, dPhase: Int, ePhase: Int) =
            setOf(aPhase, bPhase, cPhase, dPhase, ePhase).size < 5
}

class IntcodeComputer(input: String) {
    var isHalt = false

    private var state = getStateByInput(input)
    private var ptr = 0

    private var relativeBase = 0

    private var inputValues = mutableListOf<Long>()
    private var outputValue = 0L

    fun addInput(vararg input: Long): IntcodeComputer {
        input.forEach { inputValues.add(it) }
        return this
    }

    fun solve(stopAtOutput: Boolean = false): Long {
        var ptrInc = 0

        while (true) {
            var finished = false
            isHalt = false
            val num = state[ptr]

            var opcode = num
            var firstOperandMode = Mode.POSITION
            var secondOperandMode = Mode.POSITION

            if (num.specifiesParamMode()) {
                val parameterModes = num.toString()
                opcode = parameterModes[parameterModes.length - 1].toString().toLong()
                firstOperandMode = getOperandMode(parameterModes, parameterModes.length - 3)
                secondOperandMode = getOperandMode(parameterModes, parameterModes.length - 4)
            }

            when (opcode) {
                1L -> {
                    ptrInc = opcodeAdd(firstOperandMode, secondOperandMode)
                }
                2L -> {
                    ptrInc = opcodeMult(firstOperandMode, secondOperandMode)
                }
                3L -> {
                    ptrInc = opcodeSaveTo(getInput(inputValues))
                }
                4L -> {
                    val result = opcodeGetFrom(firstOperandMode)
                    outputValue = result.first
                    ptrInc = result.second

                    if (stopAtOutput) finished = true
                }
                5L -> {
                    ptr = opcodeJumpIfTrue(firstOperandMode, secondOperandMode)
                    ptrInc = 0
                }
                6L -> {
                    ptr = opcodeJumpIfFalse(firstOperandMode, secondOperandMode)
                    ptrInc = 0
                }
                7L -> {
                    ptrInc = opcodeLessThan(firstOperandMode, secondOperandMode)
                }
                8L -> {
                    ptrInc = opcodeEquals(firstOperandMode, secondOperandMode)
                }
                9L -> {
                    ptrInc = opcodeAdjustRelativeBase(firstOperandMode)
                }
                99L -> {
                    isHalt = true
                    ptr = 0
                }
                else -> {
                    println("unknown value of $num")
                }
            }

            ptr += ptrInc
            if (finished || isHalt) break
        }

        return outputValue
    }

    private fun getInput(inputValues: MutableList<Long>): Long {
        val result = inputValues[0]

        inputValues.removeAt(0)

        return result
    }

    private fun getOperandMode(parameterModes: String, index: Int): Mode {
        return if (index < 0) {
            Mode.POSITION
        } else {
            Mode.of(parameterModes[index])
        }
    }

    fun Long.specifiesParamMode(): Boolean {
        return this > 99
    }

    fun opcodeAdd(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase + 1)
        val second = get(secondOperand, ptr + 2, relativeBase + 1)
        val resultPtr = get(Mode.IMMEDIATE, ptr + 3, relativeBase + 1)

        state[resultPtr.toInt()] = first + second

        return 4
    }

    fun opcodeMult(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase + 1)
        val second = get(secondOperand, ptr + 2, relativeBase + 2)
        val resultPtr = get(Mode.IMMEDIATE, ptr + 3, relativeBase + 3)

        state[resultPtr.toInt()] = first * second

        return 4
    }

    fun opcodeSaveTo(input: Long): Int {
        val resultPtr = get(Mode.IMMEDIATE, ptr + 1, relativeBase + 1)

        state[resultPtr.toInt()] = input

        return 2
    }

    fun opcodeGetFrom(firstOperandMode: Mode): Pair<Long, Int> {
        val result = get(firstOperandMode, ptr + 1, relativeBase + 1)

        return Pair(result, 2)
    }

    fun opcodeJumpIfTrue(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase + 1)
        val second = get(secondOperand, ptr + 2, relativeBase + 2)

        return if (first != 0L) {
            second.toInt()
        } else {
            ptr + 3
        }
    }

    fun opcodeJumpIfFalse(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase + 1)
        val second = get(secondOperand, ptr + 2, relativeBase + 2)

        return if (first == 0L) {
            second.toInt()
        } else {
            ptr + 3
        }
    }

    fun opcodeLessThan(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase + 1)
        val second = get(secondOperand, ptr + 2, relativeBase + 2)
        val resultPtr = get(Mode.IMMEDIATE, ptr + 3, relativeBase + 3)

        state[resultPtr.toInt()] = if (first < second) 1 else 0

        return 4
    }

    fun opcodeEquals(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase)
        val second = get(secondOperand, ptr + 2, relativeBase)
        val resultPtr = get(Mode.IMMEDIATE, ptr + 3, relativeBase)

        state[resultPtr.toInt()] = if (first == second) 1 else 0

        return 4
    }

    fun opcodeAdjustRelativeBase(firstOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1, relativeBase)

        relativeBase += first.toInt()

        return 2
    }

    private fun getStateByInput(input: String) = input.split(',').map { it.toLong() }.toMutableList()

    fun get(operand: Mode, ptr: Int, relativeBase: Int): Long {
        return when (operand) {
            Mode.POSITION -> getPositionMode(ptr)
            Mode.IMMEDIATE -> getImmediateMode(ptr)
            Mode.RELATIVE -> getRelativeMode(ptr, relativeBase)
        }
    }

    fun getPositionMode(index: Int): Long = state[state[index].toInt()]
    fun getImmediateMode(index: Int): Long = state[index]
    fun getRelativeMode(index: Int, relativeBase: Int): Long = state[state[relativeBase + index].toInt()]
}

enum class Mode {
    POSITION, IMMEDIATE, RELATIVE;

    companion object {
        fun of(value: Char): Mode {
            return when (value) {
                '0' -> POSITION
                '1' -> IMMEDIATE
                '2' -> RELATIVE
                else -> throw RuntimeException("Unknown mode value of $value")
            }
        }
    }
}

fun main() {
    val app = App()
    val input = File("input.txt").readLines()[0]

    println(app.solveFirst(input))
    println(app.solveSecond(input))
}
