import math, strutils, strformat, sequtils

type 
  Moon* = ref tuple 
    x, y, z: int
    vX, vY, vZ: int

proc solveFirst*(input: seq[string], steps: int): int
proc solveSecond*(input: seq[string]): uint64


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input, 1000)
  echo solveSecond(input)

proc getCoor(line: string): int =
  return line.split('=')[1].parseInt

proc newMoon*(x, y, z, vX, vY, vZ: int): Moon =
  var moon: Moon
  new(moon)
  moon.x = x
  moon.y = y
  moon.z = z
  moon.vX = vX
  moon.vY = vY
  moon.vZ = vZ
  return moon 

proc getMoon(line: string): Moon =
  let parts = line[1..^2].split(',')

  return newMoon(getCoor(parts[0]), getCoor(parts[1]), getCoor(parts[2]), 0, 0, 0)

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

proc performStep(moons: seq[Moon]) =
  for i in 0..<moons.len:
    for j in i+1..<moons.len:
      var first = moons[i]
      var second = moons[j]
      applyGravity(first, second)
  
  for moon in moons:
    var refMoon = moon
    refMoon.applyVelocity()

proc same(originalMoons, moons: seq[Moon], checkX = false, checkY = false, checkZ = false): bool =
  let pairs = zip(originalMoons, moons)
  for pair in pairs:
    let first = pair[0]
    let second = pair[1]

    if checkX: 
      if first.x != second.x or first.vX != second.vX: return false

    if checkY: 
      if first.y != second.y or first.vY != second.vY: return false
    
    if checkZ: 
      if first.z != second.z or first.vZ != second.vZ: return false

  return true 

proc getLcm(x, y, z: uint64): uint64 =
  var mult: uint64 = 1

  while true:
    mult += 1

    let lcm = x * mult

    if (lcm mod y) == 0 and (lcm div y) >= 1.uint64 and (lcm mod z) == 0 and (lcm div z) >= 1.uint64:
      return lcm
    
    if mult mod 100000000 == 0:
      echo fmt"mult:{mult}"

proc solveFirst*(input: seq[string], steps: int): int =
  let moons = getMoons(input)

  for step in 1..steps:
    moons.performStep()

  return moons.calculateEnergy()
  
proc solveSecond*(input: seq[string]): uint64 =
  let originalMoons = getMoons(input)
  let moons = getMoons(input)

  var sameX = false
  var sameY = false
  var sameZ = false 
  var cycleX: uint64 = 0
  var cycleY: uint64 = 0
  var cycleZ: uint64 = 0

  var steps: uint64 = 0
  while not (sameX and sameY and sameZ):
    steps += 1
    #TODO perform calculation to only required axis
    moons.performStep()

    if not sameX:
      sameX = same(originalMoons, moons, checkX=true)
      if sameX: cycleX = steps
    if not sameY:
      sameY = same(originalMoons, moons, checkY=true)
      if sameY: cycleY = steps
    if not sameZ:
      sameZ = same(originalMoons, moons, checkZ=true)
      if sameZ: cycleZ = steps

  return getLcm(cycleX, cycleY, cycleZ) 
