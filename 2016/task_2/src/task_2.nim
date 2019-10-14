import strutils


proc solveFirst*(input: seq[string]): string
proc solveSecond*(input: seq[string]): string


when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)
  echo solveSecond(input)


type
  Keypad[W, H: static[int]] = array[0..W-1, array[0..H-1, char]]
  Coordinate = object
    x, y: int

proc incLimit(number, max: int): int =
  if number == max: return max
  return number + 1

proc decLimit(number: int, min = 0): int =
  if number == min: return min
  return number - 1

proc applyMovement(coordinate: Coordinate, movement: char, maxX, maxY: int): Coordinate =
  result.x = coordinate.x
  result.y = coordinate.y

  case movement:
    of 'L':
      result.x = result.x.decLimit
    of 'U':
      result.y = result.y.decLimit
    of 'R':
      result.x = result.x.incLimit maxX
    of 'D':
      result.y = result.y.incLimit maxY
    else:
      echo "unknown movement " & movement

proc solve(input: seq[string], keypad: Keypad, initialCoordinate: Coordinate): string =
  var coordinate = initialCoordinate

  var code: string

  for line in input:
    for movement in line:
      let originalCoordinate = coordinate

      coordinate = applyMovement(coordinate, movement, maxX = keypad.high, maxY = keypad.high)

      if keypad[coordinate.y][coordinate.x] == ' ':
        coordinate = originalCoordinate

    code = code & keypad[coordinate.y][coordinate.x]

  return code

proc solveFirst*(input: seq[string]): string =
  let keypad: Keypad[3, 3] = [
    ['1', '2', '3'],
    ['4', '5', '6'],
    ['7', '8', '9']
  ]

  return solve(input, keypad, Coordinate(x: 1, y: 1))

proc solveSecond*(input: seq[string]): string =
  let keypad: Keypad[5, 5] = [
    [' ', ' ', '1', ' ', ' '],
    [' ', '2', '3', '4', ' '],
    ['5', '6', '7', '8', '9'],
    [' ', 'A', 'B', 'C', ' '],
    [' ', ' ', 'D', ' ', ' ']
  ]

  return solve(input, keypad, Coordinate(x: 0, y: 2))
