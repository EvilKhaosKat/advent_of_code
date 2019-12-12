import ../src/task_12
import strutils, strformat

when isMainModule:
  doAssert Moon(x: 8, y: 12, z: -9, vX: -7, vY: 3, vZ: 0).calculateEnergy() == 290

  let input = @["<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"]
  doAssert solveFirst(input, 10) == 179
  