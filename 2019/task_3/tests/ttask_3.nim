import ../src/task_3
import strutils, strformat

when isMainModule:
  doAssert solveFirst(@["R8,U5,L5,D3", "U7,R6,D4,L4"]) == 6
  doAssert solveFirst(@["R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"]) == 159 
  doAssert solveFirst(@["R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"]) == 135 

