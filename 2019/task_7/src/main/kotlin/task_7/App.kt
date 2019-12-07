package task_7

import java.io.File
import java.lang.RuntimeException

//Lets try simple hardcode solution
class App {
    fun solveFirst(input: String): Int {
        var maxThruster = 0

        val intcodeComputer = IntcodeComputer()

        for (aPhase in 0..4) {
            val aOutput = intcodeComputer.solve(input, mutableListOf(aPhase, 0))
            for (bPhase in 0..4) {
                val bOutput = intcodeComputer.solve(input, mutableListOf(bPhase, aOutput))
                for (cPhase in 0..4) {
                    val cOutput = intcodeComputer.solve(input, mutableListOf(cPhase, bOutput))
                    for (dPhase in 0..4) {
                        val dOutput = intcodeComputer.solve(input, mutableListOf(dPhase, cOutput))
                        for (ePhase in 0..4) {
                            val eOutput = intcodeComputer.solve(input, mutableListOf(ePhase, dOutput))

                            //not optimal to do it here, but simple
                            if (setOf(aPhase, bPhase, cPhase, dPhase, ePhase).size < 5) {
                                continue
                            }

                            if (eOutput > maxThruster) {
                                maxThruster = eOutput
                            }
                        }
                    }
                }
            }
        }

        return maxThruster
    }

    fun solveSecond(input: String): Int {
        return -1
    }
}

class IntcodeComputer {
    fun solve(input: String, inputValues: MutableList<Int>): Int {
        val state = getStateByInput(input)

        var outputValue = 0

        var ptr = 0
        var ptrInc = 0

        while (true) {
            var finished = false
            val num = state[ptr]

            var opcode = num
            var firstOperandMode = Mode.POSITION
            var secondOperandMode = Mode.POSITION
            var thirdOperandMode = Mode.POSITION

            if (num.specifiesParamMode()) {
                val parameterModes = num.toString()
                opcode = parameterModes[parameterModes.length - 1].toString().toInt()
                firstOperandMode = getOperandMode(parameterModes, parameterModes.length - 3)
                secondOperandMode = getOperandMode(parameterModes, parameterModes.length - 4)
                thirdOperandMode = getOperandMode(parameterModes, parameterModes.length - 5)
            }

            when (opcode) {
                1 -> {
                    ptrInc = state.opcodeAdd(ptr, firstOperandMode, secondOperandMode)
                }
                2 -> {
                    ptrInc = state.opcodeMult(ptr, firstOperandMode, secondOperandMode)
                }
                3 -> {
                    ptrInc = state.opcodeSaveTo(ptr, inputValues[0])
                    inputValues.removeAt(0)
                }
                4 -> {
                    val result = state.opcodeGetFrom(ptr, firstOperandMode)
                    outputValue = result.first
                    ptrInc = result.second
                }
                5 -> {
                    ptr = state.opcodeJumpIfTrue(ptr, firstOperandMode, secondOperandMode)
                    ptrInc = 0
                }
                6 -> {
                    ptr = state.opcodeJumpIfFalse(ptr, firstOperandMode, secondOperandMode)
                    ptrInc = 0
                }
                7 -> {
                    ptrInc = state.opcodeLessThan(ptr, firstOperandMode, secondOperandMode, thirdOperandMode)
                }
                8 -> {
                    ptrInc = state.opcodeEquals(ptr, firstOperandMode, secondOperandMode, thirdOperandMode)
                }
                99 -> {
                    finished = true
                }
                else -> {
                    println("unknown value of $num")
                }
            }

            if (finished) break
            ptr += ptrInc
        }

        return outputValue
    }

    private fun getOperandMode(parameterModes: String, index: Int): Mode {
        return if (index < 0) {
            Mode.POSITION
        } else {
            Mode.of(parameterModes[index])
        }
    }

    fun Int.specifiesParamMode(): Boolean {
        return this > 99
    }

    fun MutableList<Int>.opcodeAdd(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 3)

        this[resultPtr] = first + second

        return 4
    }

    fun MutableList<Int>.opcodeMult(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 3)

        this[resultPtr] = first * second

        return 4
    }

    fun MutableList<Int>.opcodeSaveTo(ptr: Int, input: Int): Int {
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 1)

        this[resultPtr] = input

        return 2
    }

    fun MutableList<Int>.opcodeGetFrom(ptr: Int, firstOperandMode: Mode): Pair<Int, Int> {
        val result = firstOperandMode.get(this, ptr + 1)

        return Pair(result, 2)
    }

    fun MutableList<Int>.opcodeJumpIfTrue(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)

        return if (first != 0) {
            second
        } else {
            ptr + 3
        }
    }

    fun MutableList<Int>.opcodeJumpIfFalse(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)

        return if (first == 0) {
            second
        } else {
            ptr + 3
        }
    }

    fun MutableList<Int>.opcodeLessThan(ptr: Int, firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 3)

        this[resultPtr] = if (first < second) 1 else 0

        return 4
    }

    fun MutableList<Int>.opcodeEquals(ptr: Int, firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 3)

        this[resultPtr] = if (first == second) 1 else 0

        return 4
    }

    private fun getStateByInput(input: String) = input.split(',').map { it.toInt() }.toMutableList()
}

enum class Mode {
    POSITION {
        override fun get(state: List<Int>, ptr: Int): Int = state.getPositionMode(ptr)
    },
    IMMEDIATE {
        override fun get(state: List<Int>, ptr: Int): Int = state.getImmediateMode(ptr)
    };

    abstract fun get(state: List<Int>, ptr: Int): Int

    companion object {
        fun of(value: Char): Mode {
            return when (value) {
                '0' -> POSITION
                '1' -> IMMEDIATE
                else -> throw RuntimeException("Unknown mode value of $value")
            }
        }
    }
}

fun List<Int>.getPositionMode(index: Int): Int = this[this[index]]
fun List<Int>.getImmediateMode(index: Int): Int = this[index]

fun main() {
    val app = App()
    val input = File("input.txt").readLines()[0]

    println(app.solveFirst(input))
    println(app.solveSecond(input))
}
