import math, strutils, strformat, sequtils

type 
  Moon* = ref object
    x*, y*, z*: int
    vX*, vY*, vZ*: int

proc solveFirst*(input: seq[string], steps: int): int
proc solveSecond*(input: seq[string]): int


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input, 1000)
  echo solveSecond(input)

proc getCoor(line: string): int =
  return line.split('=')[1].parseInt

proc getMoon(line: string): Moon =
  let parts = line[1..^2].split(',')

  return Moon(x: getCoor(parts[0]), y: getCoor(parts[1]), z: getCoor(parts[2]), vX: 0, vY: 0, vZ: 0)

proc getMoons(input: seq[string]): seq[Moon] =
  return input.mapIt(getMoon(it))

proc applyGravity(first, second: var Moon) =
  if first.x > second.x:
    first.vX -= 1
    second.vX += 1
  elif first.x < second.x:
    first.vX += 1
    second.vX -= 1

  if first.y > second.y:
    first.vY -= 1
    second.vY += 1
  elif first.y < second.y:
    first.vY += 1
    second.vY -= 1

  if first.z > second.z:
    first.vZ -= 1
    second.vZ += 1
  elif first.z < second.z:
    first.vZ += 1
    second.vZ -= 1
    
proc applyVelocity(moon: var Moon) =
  moon.x += moon.vX
  moon.y += moon.vY
  moon.z += moon.vZ

proc calculateEnergy*(moon: Moon): int =
  return (moon.x.abs + moon.y.abs + moon.z.abs)*(moon.vX.abs + moon.vY.abs + moon.vZ.abs)
 
proc calculateEnergy*(moons: seq[Moon]): int =
  return moons
          .mapIt(it.calculateEnergy)
          .foldl(a + b)

proc solveFirst*(input: seq[string], steps: int): int =
  let moons = getMoons(input)

  for step in 1..steps:
    for i in 0..<moons.len:
      for j in i+1..<moons.len:
        var first = moons[i]
        var second = moons[j]
        applyGravity(first, second)
    
    for moon in moons:
      var refMoon = moon
      refMoon.applyVelocity()

  return moons.calculateEnergy()
  

proc solveSecond*(input: seq[string]): int =
  let moons = getMoons(input)
  return -1