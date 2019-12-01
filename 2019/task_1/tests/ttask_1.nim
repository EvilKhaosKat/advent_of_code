import  ../src/task_1

when isMainModule: 
    doAssert solveFirst(@["14"]) == 2
    doAssert solveFirst(@["1969"]) == 654
    doAssert solveFirst(@["14", "1969"]) == 656

    doAssert solveSecond(@["14"]) == 2
    doAssert solveSecond(@["1969"]) == 966