package task_7

import java.io.File
import java.lang.RuntimeException

//Lets try simple hardcode solution
class App {
    fun solveFirst(input: String): Int {
        var maxThruster = 0

        for (aPhase in 0..4) {
            for (bPhase in 0..4) {
                for (cPhase in 0..4) {
                    for (dPhase in 0..4) {
                        for (ePhase in 0..4) {
                            //not optimal to do it here, but simple
                            if (setOf(aPhase, bPhase, cPhase, dPhase, ePhase).size < 5) {
                                continue
                            }

                            val aOutput = IntcodeComputer(input).addInput(aPhase, 0).solve()
                            val bOutput = IntcodeComputer(input).addInput(bPhase, aOutput).solve()
                            val cOutput = IntcodeComputer(input).addInput(cPhase, bOutput).solve()
                            val dOutput = IntcodeComputer(input).addInput(dPhase, cOutput).solve()
                            val eOutput = IntcodeComputer(input).addInput(ePhase, dOutput).solve()

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
        var maxThruster = 0

        for (aPhase in 5..9) {
            for (bPhase in 5..9) {
                for (cPhase in 5..9) {
                    for (dPhase in 5..9) {
                        for (ePhase in 5..9) {
                            if (setOf(aPhase, bPhase, cPhase, dPhase, ePhase).size < 5) {
                                continue
                            }

                            val intcodeComputerA = IntcodeComputer(input)
                            val intcodeComputerB = IntcodeComputer(input)
                            val intcodeComputerC = IntcodeComputer(input)
                            val intcodeComputerD = IntcodeComputer(input)
                            val intcodeComputerE = IntcodeComputer(input)

                            var aOutput = intcodeComputerA.addInput(aPhase, 0).solve(true)
                            var bOutput = intcodeComputerB.addInput(bPhase, aOutput).solve(true)
                            var cOutput = intcodeComputerC.addInput(cPhase, bOutput).solve(true)
                            var dOutput = intcodeComputerD.addInput(dPhase, cOutput).solve(true)
                            var eOutput = intcodeComputerE.addInput(ePhase, dOutput).solve(true)

                            var cycles = 0
                            while (!intcodeComputerE.isHalt) {
                                cycles++
                                if (cycles > 9_000_000) {
                                    throw RuntimeException("sanity limit - too much cycles: $aPhase,$bPhase,$cPhase,$dPhase,$ePhase")
                                }

                                aOutput = intcodeComputerA.addInput(eOutput).solve(true)
                                bOutput = intcodeComputerB.addInput(aOutput).solve(true)
                                cOutput = intcodeComputerC.addInput(bOutput).solve(true)
                                dOutput = intcodeComputerD.addInput(cOutput).solve(true)
                                eOutput = intcodeComputerE.addInput(dOutput).solve(true)

                                if (eOutput > maxThruster) {
                                    maxThruster = eOutput
                                }
                            }
                        }
                    }
                }
            }
        }

        return maxThruster
    }
}

class IntcodeComputer(input: String) {
    var isHalt = false
    private var state = getStateByInput(input)
    private var ptr = 0
    private var inputValues = mutableListOf<Int>()
    private var outputValue = 0

    fun addInput(vararg input: Int): IntcodeComputer {
        input.forEach { inputValues.add(it) }
        return this
    }

    fun solve(stopAtOutput: Boolean = false): Int {
        var ptrInc = 0

        while (true) {
            var finished = false
            isHalt = false
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
                    ptrInc = state.opcodeSaveTo(ptr, getInput(inputValues))
                }
                4 -> {
                    val result = state.opcodeGetFrom(ptr, firstOperandMode)
                    outputValue = result.first
                    ptrInc = result.second
                    if (stopAtOutput) finished = true
//                    println("output: $outputValue")
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
                    ptrInc = state.opcodeLessThan(ptr, firstOperandMode, secondOperandMode)
                }
                8 -> {
                    ptrInc = state.opcodeEquals(ptr, firstOperandMode, secondOperandMode)
                }
                99 -> {
                    finished = true
                    isHalt = true
                    ptr = 0
//                    println("halt!")
                }
                else -> {
                    println("unknown value of $num")
                }
            }

            ptr += ptrInc
            if (finished) break
        }

        return outputValue
    }

    private fun getInput(inputValues: MutableList<Int>): Int {
        val result = inputValues[0]

//        if (inputValues.size > 1) {
            inputValues.removeAt(0)
//        }

        return result
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

    fun MutableList<Int>.opcodeLessThan(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
        val first = firstOperand.get(this, ptr + 1)
        val second = secondOperand.get(this, ptr + 2)
        val resultPtr = Mode.IMMEDIATE.get(this, ptr + 3)

        this[resultPtr] = if (first < second) 1 else 0

        return 4
    }

    fun MutableList<Int>.opcodeEquals(ptr: Int, firstOperand: Mode, secondOperand: Mode): Int {
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
