from strutils import parseInt, strip, splitLines
import sequtils, algorithm, strformat

proc solveFirst(input: seq[string]): int =
    var numbers: seq[int] = input.map(parseInt)
    numbers.sort()

    var diff1 = 0
    var diff2 = 0
    var diff3 = 0

    var currentJolts = 0
    for number in numbers:
        if number - 3 == currentJolts:
            diff3 += 1
            currentJolts = number
        elif number - 2 == currentJolts:
            diff2 += 1
            currentJolts = number
        elif number - 1 == currentJolts:
            diff1 += 1
            currentJolts = number
        else:
            echo fmt"impossible to choose adapter, currentJolts={currentJolts}, number={number}"

    diff3 += 1
    echo fmt"diff1={diff1}, diff2={diff2}, diff3={diff3}"

    return diff1 * diff3

when isMainModule:
    let input = readFile("input.txt").strip.splitLines
    echo solveFirst input

