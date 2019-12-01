import strutils, sequtils


proc solveFirst*(input: seq[string]): int 
proc solveSecond*(input: seq[string]): int 


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)
  echo solveSecond(input)




proc solveFirst*(input: seq[string]): int =
  return input.mapIt(it.parseInt)
  .mapIt(it div 3 - 2)
  .foldl(a + b)


proc solveSecond*(input: seq[string]): int =
  return -1
