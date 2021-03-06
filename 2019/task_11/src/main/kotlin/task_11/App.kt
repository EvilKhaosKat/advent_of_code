package task_11

import java.io.File
import java.lang.RuntimeException

class App {
    fun solveFirst(input: String): Int {
        return getDrawnPanels(input, Color.BLACK).size
    }

    fun solveSecond(input: String): Int {
        val panels = getDrawnPanels(input, Color.WHITE)

        val minX = panels.keys.minBy { it.x }!!.x
        val maxX = panels.keys.maxBy { it.x }!!.x
        val minY = panels.keys.minBy { it.y }!!.y
        val maxY = panels.keys.maxBy { it.y }!!.y

        for (y in minY..maxY) {
            for (x in minX..maxX) {
                val colorToPaint = panels.getOrDefault(Point(x, y), Color.BLACK)
                when (colorToPaint) {
                    Color.BLACK -> print(" ")
                    Color.WHITE -> print("█")
                }
            }

            println()
        }
        return -1
    }

    private fun getDrawnPanels(input: String, startPanelColor: Color): Map<Point, Color> {
        val panels = HashMap<Point, Color>()
        panels[Point(0, 0)] = startPanelColor

        var direction = Direction.UP
        var x = 0
        var y = 0

        val computer = IntcodeComputer(input)

        val defaultColor = Color.BLACK

        while (!computer.isHalt) {
            val currentPanelColor = panels.getOrDefault(Point(x, y), defaultColor)
            computer.addInput(currentPanelColor.num.toLong())

            val paintToColor = computer.solve(true)
            if (computer.isHalt) {
                break
            }

            panels[Point(x, y)] = Color.of(paintToColor.toInt())

            direction = when (computer.solve(true)) {
                0L -> direction.left()
                1L -> direction.right()
                else -> throw RuntimeException("Unknown direction to turn")
            }

            x += direction.x
            y += direction.y
        }

        return panels
    }
}

data class Point(val x: Int, val y: Int)

class IntcodeComputer(input: String) {
    var isHalt = false

    private var state = getStateByInput(input)
    private var ptr = 0

    private var relativeBase = 0

    private val extendedMemory = HashMap<Int, Long>()

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
            var thirdOperandMode = Mode.POSITION

            if (num.specifiesParamMode()) {
                val parameterModes = num.toString()
                opcode = parameterModes[parameterModes.length - 1].toString().toLong()
                firstOperandMode = getOperandMode(parameterModes, parameterModes.length - 3)
                secondOperandMode = getOperandMode(parameterModes, parameterModes.length - 4)
                thirdOperandMode = getOperandMode(parameterModes, parameterModes.length - 5)
            }

            when (opcode) {
                1L -> {
                    ptrInc = opcodeAdd(firstOperandMode, secondOperandMode, thirdOperandMode)
                }
                2L -> {
                    ptrInc = opcodeMult(firstOperandMode, secondOperandMode, thirdOperandMode)
                }
                3L -> {
                    ptrInc = opcodeSaveTo(firstOperandMode, getInput(inputValues))
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
                    ptrInc = opcodeLessThan(firstOperandMode, secondOperandMode, thirdOperandMode)
                }
                8L -> {
                    ptrInc = opcodeEquals(firstOperandMode, secondOperandMode, thirdOperandMode)
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

    fun opcodeAdd(firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)
        val resultPtr = getIndex(thirdOperandMode, ptr + 3)

        setByIndex(resultPtr, first + second)

        return 4
    }

    fun opcodeMult(firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)
        val resultPtr = getIndex(thirdOperandMode, ptr + 3)

        setByIndex(resultPtr, first * second)

        return 4
    }

    fun opcodeSaveTo(firstOperand: Mode, input: Long): Int {
        val resultPtr = getIndex(firstOperand, ptr + 1)

        setByIndex(resultPtr, input)

        return 2
    }

    fun opcodeGetFrom(firstOperandMode: Mode): Pair<Long, Int> {
        val result = get(firstOperandMode, ptr + 1)
        //getByIndex(relativeBase + getByIndex(index).toInt())

        return Pair(result, 2)
    }

    fun opcodeJumpIfTrue(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)

        return if (first != 0L) {
            second.toInt()
        } else {
            ptr + 3
        }
    }

    fun opcodeJumpIfFalse(firstOperand: Mode, secondOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)

        return if (first == 0L) {
            second.toInt()
        } else {
            ptr + 3
        }
    }

    fun opcodeLessThan(firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)
        val resultPtr = getIndex(thirdOperandMode, ptr + 3)

        setByIndex(resultPtr, if (first < second) 1 else 0)

        return 4
    }

    fun opcodeEquals(firstOperand: Mode, secondOperand: Mode, thirdOperandMode: Mode): Int {
        val first = get(firstOperand, ptr + 1)
        val second = get(secondOperand, ptr + 2)
        val resultPtr = getIndex(thirdOperandMode, ptr + 3)

        setByIndex(resultPtr, if (first == second) 1 else 0)

        return 4
    }

    fun opcodeAdjustRelativeBase(firstOperand: Mode): Int {
        val first = get(firstOperand, ptr + 1)

        relativeBase += first.toInt()

        return 2
    }

    private fun getStateByInput(input: String) = input.split(',').map { it.toLong() }.toMutableList()

    fun get(operand: Mode, ptr: Int): Long {
        return when (operand) {
            Mode.POSITION -> getPositionMode(ptr)
            Mode.IMMEDIATE -> getImmediateMode(ptr)
            Mode.RELATIVE -> getRelativeMode(ptr, relativeBase)
        }
    }

    fun getIndex(operand: Mode, ptr: Int): Int {
        return when (operand) {
            Mode.POSITION -> getByIndex(ptr).toInt()
            Mode.RELATIVE -> relativeBase + getByIndex(ptr).toInt()
            else -> throw RuntimeException("Can't use $operand to get address to write")
        }
    }

    fun getPositionMode(index: Int): Long = getByIndex(getByIndex(index).toInt())
    fun getImmediateMode(index: Int): Long = getByIndex(index)
    fun getRelativeMode(index: Int, relativeBase: Int): Long = getByIndex(relativeBase + getByIndex(index).toInt())

    fun getByIndex(index: Int): Long {
        if (index >= state.size) {
            return extendedMemory.getOrDefault(index, 0)
        }

        return state[index]
    }

    fun setByIndex(index: Int, value: Long) {
        if (index >= state.size) {
            extendedMemory[index] = value
        } else {
            state[index] = value
        }
    }
}

enum class Direction(val x: Int, val y: Int) {
    UP(0, -1), RIGHT(1, 0), DOWN(0, 1), LEFT(-1, 0);

    fun left(): Direction {
        val values = values()
        var nextOrdinal = (ordinal - 1) % values.size

        if (nextOrdinal < 0) {
            nextOrdinal = values.size - 1
        }

        return values[nextOrdinal]
    }

    fun right(): Direction {
        val values = values()
        val nextOrdinal = (ordinal + 1) % values.size
        return values[nextOrdinal]
    }
}

enum class Color(val num: Int) {
    BLACK(0), WHITE(1);

    companion object {
        fun of(value: Int): Color {
            return when (value) {
                BLACK.num -> BLACK
                WHITE.num -> WHITE
                else -> throw RuntimeException("Unknown color value of $value")
            }
        }
    }
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