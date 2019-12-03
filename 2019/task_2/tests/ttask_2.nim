import ../src/task_2
import strutils

when isMainModule:
  doAssert solveFirst("1,0,0,0,99".split(','), noun = 0, verb = 0) == 2
  doAssert solveFirst("1,1,1,4,99,5,6,0,99".split(','), noun = 1, verb = 1) == 30 

