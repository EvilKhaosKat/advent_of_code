import strutils, sequtils


proc solveFirst*(input: seq[string], noun, verb: int): int
proc solveSecond*(input: seq[string]): int

proc restoreGravity(input: var seq[int], noun, verb: int): seq[int]
proc getValue(state: seq[int], address: int): int

when isMainModule:
  let input = readFile("input.txt").strip.split(',')

  echo solveFirst(input, noun = 12, verb = 2)
  echo solveSecond(input)


proc execute(state: var seq[int], noun, verb: int): int =
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

proc solveFirst*(input: seq[string], noun, verb: int): int =
  var state = input.mapIt(it.parseInt)
  state = restoreGravity(state, noun, verb)

  return execute(state, noun, verb)

proc restoreGravity(input: var seq[int], noun, verb: int): seq[int] =
  input[1] = noun 
  input[2] = verb 
  return input

proc getValue(state: seq[int], address: int): int =
  let addressOfValue = state[address]
  return state[addressOfValue]

proc solveSecond*(input: seq[string]): int =
  return -1
