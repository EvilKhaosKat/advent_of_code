import ../src/task_4
import strutils, strformat


when isMainModule:
  doAssert isCorrectPart1 Password(@[1, 1, 1, 1, 1, 1])
  doAssert not isCorrectPart1 Password(@[2, 2, 3, 4, 5, 0])
  doAssert not isCorrectPart1 Password(@[1, 2, 3, 7, 8, 9])

  doAssert isCorrectPart2 Password(@[1, 1, 2, 2, 3, 3])
  doAssert not isCorrectPart2 Password(@[1, 2, 3, 4, 4, 4])
  doAssert isCorrectPart2 Password(@[1, 1, 1, 1, 2, 2])

