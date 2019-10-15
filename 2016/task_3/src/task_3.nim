import strutils, sequtils, algorithm


proc solveFirst*(input: seq[string]): int


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)


proc solveFirst*(input: seq[string]): int =
  for line in input:
    var numbers = line
      .split({' '})
      .filterIt(it != "")
      .mapIt(it.parseInt)
      .sorted

    if numbers[0] + numbers[1] > numbers[2]: result.atomicInc
