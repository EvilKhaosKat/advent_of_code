import strutils, strformat, sequtils


proc solveFirst*(input: seq[string]): int
proc solveSecond*(input: seq[string]): int

type
  Password* = seq[int]


proc inc(password: var Password): Password
proc inc(password: var Password, index: int): Password
proc isCorrectPart1*(password: Password): bool
proc isCorrectPart2*(password: Password): bool
proc isSame(password: Password, value: string): bool


# TODO skip excess subranges here instead of validation func
when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst(input)
  echo solveSecond(input)

proc getPasswordAsSeq(password: string): seq[int] =
  return password.toSeq().mapIt(parseInt($it))

proc solve(input: seq[string], isCorrect: proc (password: Password): bool): int =
  let fromPassword = input[0]
  let toPassword = input[1]

  var password = Password(fromPassword.getPasswordAsSeq)

  var passwordsCount = 0
  while true:
    password = password.inc()

    if password.isSame(toPassword):
      break

    if isCorrect(password):
      passwordsCount += 1

  return passwordsCount

proc solveFirst*(input: seq[string]): int =
  return solve(input, isCorrectPart1) 

proc solveSecond*(input: seq[string]): int =
  return solve(input, isCorrectPart2) 

proc inc(password: var Password): Password =
  return password.inc(password.len - 1)

proc inc(password: var Password, index: int): Password =
  if index < 0:
    echo fmt"Error: can't increment password any further than {password}"
    return password

  var val = password[index]
  if val == 9:
    password[index] = 0
    return password.inc(index - 1)
  else:
    password[index] = val + 1

  return password

proc isCorrectPart1*(password: Password): bool =
  var hasSameAdjacentsDigits = false

  var prevDigit = password[0]
  for index in 1..<password.len:
    var digit = password[index]
    if prevDigit == digit:
      hasSameAdjacentsDigits = true

    if digit < prevDigit:
      return false

    prevDigit = digit

  return hasSameAdjacentsDigits

proc isCorrectPart2*(password: Password): bool =
  var hasSameAdjacentsDigits = false
  var groupSize = 1

  var prevDigit = password[0]
  for index in 1..<password.len:
    var digit = password[index]

    if prevDigit == digit:
      groupSize += 1
    else:
      if groupSize == 2:
        hasSameAdjacentsDigits = true

      groupSize = 1

    if digit < prevDigit:
      return false

    prevDigit = digit

  if groupSize == 2:
    hasSameAdjacentsDigits = true

  return hasSameAdjacentsDigits

proc isSame(password: Password, value: string): bool =
  return password.join == value
