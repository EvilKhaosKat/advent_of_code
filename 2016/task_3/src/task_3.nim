import strutils, sequtils, algorithm


proc solveFirst*(input: seq[string]): int


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)


proc solveFirst*(input: seq[string]): int =
  input
    .mapIt(it.split({' '})
      .filterIt(it != "")
      .mapIt(it.parseInt)
      .sorted)
    .mapIt(if it[0] + it[1] > it[2]: 1 else: 0)
    .foldl(a + b)
