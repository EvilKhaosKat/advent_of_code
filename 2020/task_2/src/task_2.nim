import strutils

proc solveFirst*(input: seq[string]): int =
    var result = 0

    #1-3 a: abcde
    for line in input:
        let parts = line.split(" ")
        let countFormat = parts[0].split("-")

        let minCount = countFormat[0].parseInt
        let maxCount = countFormat[1].parseInt

        let letter = parts[1][0]

        let password = parts[2]

        var count = 0
        for c in password:
            if c == letter:
                count += 1

        if count >= minCount and count <= maxCount:
            result += 1

    return result

proc solveSecond*(input: seq[string]): int =
    var result = 0

    #1-3 a: abcde
    for line in input:
        let parts = line.split(" ")
        let countFormat = parts[0].split("-")

        let idx1 = countFormat[0].parseInt
        let idx2 = countFormat[1].parseInt

        let letter = parts[1][0]

        let password = parts[2]

        if (password[idx1-1] == letter) xor (password[idx2-1] == letter):
            result += 1

    return result

when isMainModule:
  let input = readFile("input.txt").strip.splitLines

  echo solveFirst input
  echo solveSecond input





