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

    return diff1 * diff3

proc calc(index: int, numbers: seq[int]): int =
    if index == numbers.len() - 1:
        return 1

    var combinations = 0

    var check_index = index + 1;
    while check_index < numbers.len() - 1:
        if numbers[check_index] - numbers[index] <= 3:
            check_index += 1
        else:
            break

    echo fmt"{numbers[index]} {numbers[check_index]} {check_index - index - 1}"
    return (check_index - index - 1) * calc(check_index, numbers)

#TODO rely on dynamic programming/matrix for calculation - paths count in *graph*
proc solveSecond(input: seq[string]): int =
    var numbers: seq[int] = input.map(parseInt)
    numbers.sort()

    echo numbers

    return calc(0, numbers)

when isMainModule:
    let input = readFile("input.txt").strip.splitLines
    echo solveFirst input
    echo solveSecond input
