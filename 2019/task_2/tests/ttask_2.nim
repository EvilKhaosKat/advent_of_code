import ../src/task_2
import strutils

when isMainModule:
  doAssert solveFirst("1,0,0,0,99".split(',')) == 2
  doAssert solveFirst("1,1,1,4,99,5,6,0,99".split(',')) == 30 

    #doAssert solveSecond(@["14"]) == 2
    #doAssert solveSecond(@["1969"]) == 966

