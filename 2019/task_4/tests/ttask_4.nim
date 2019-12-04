import ../src/task_4
import strutils, strformat

when isMainModule:
  doAssert solveFirst(@["R8,U5,L5,D3", "U7,R6,D4,L4"]) == 6

