import strutils, sequtils


proc solveFirst*(input: seq[string]): int
proc solveSecond*(input: seq[string]): int

proc restoreGravity(input: var seq[int]): seq[int]
proc getValue(state: seq[int], index: int): int

when isMainModule:
  let input = readFile("input.txt").strip.split(',')

  echo solveFirst(input)
  echo solveSecond(input)


proc solveFirst*(input: seq[string]): int =
  var state = input.mapIt(it.parseInt)
  #state = restoreGravity(state)

  var currOpcodePtr = 0

  while true:
    let currOpcode = state[currOpcodePtr] 

    case currOpcode:
      of 1:
        state[state[currOpcodePtr + 3]] = getValue(state, currOpcodePtr+1) + getValue(state, currOpcodePtr+2)
      of 2:
        state[state[currOpcodePtr + 3]] = getValue(state, currOpcodePtr+1) * getValue(state, currOpcodePtr+2)
      of 99:
        break
      else:
        echo "unknown opcode " & currOpcode.intToStr
        

    currOpcodePtr += 4
  
  return state[0] 

proc restoreGravity(input: var seq[int]): seq[int] =
  input[1] = 12
  input[2] = 2
  return input

proc getValue(state: seq[int], index: int): int =
  let indexToUse = state[index]
  return state[indexToUse]

proc solveSecond*(input: seq[string]): int =
  return -1
