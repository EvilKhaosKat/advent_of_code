import ../src/task_12
import strutils, strformat

when isMainModule:
  doAssert newMoon(8, 12, -9, -7, 3, 0).calculateEnergy() == 290

  let input1 = @["<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"]
  doAssert solveFirst(input1, 10) == 179
  
  let input2 = @["<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"]
  let result2 = solveSecond(input2)
  echo result2
  doAssert result2 == 2772.uint64

  let longInput = @["<x=-8, y=-10, z=0>", "<x=5, y=5, z=10>", "<x=2, y=-7, z=3>", "<x=9, y=-8, z=-3>"] 
  let resultLongInput = solveSecond(longInput)
  echo resultLongInput
  doAssert resultLongInput == 4686774924.uint64