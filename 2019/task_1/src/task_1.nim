import strutils, sequtils


proc solveFirst*(input: seq[string]): int
proc solveSecond*(input: seq[string]): int


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)
  echo solveSecond(input)


proc calcFuel(mass: int): int =
  return mass div 3 - 2

proc calcFullFuel(mass: int, fullMass = 0): int =
  let fuelMass = calcFuel mass

  if fuelMass < 0:
    return fullMass

  return calcFullFuel(fuelMass, fullMass + fuelMass)

proc solveFirst*(input: seq[string]): int =
  return input.mapIt(it.parseInt)
    .mapIt(calcFuel it)
    .foldl(a + b)


proc solveSecond*(input: seq[string]): int =
  return input.mapIt(it.parseInt)
    .mapIt(calcFullFuel it)
    .foldl(a + b)
