def solve_first(input):
    code = []

    for line in input:
        instruction = line.split(" ")
        code.append(Instruction(instruction[0], instruction[1].strip()))

    console = Console(code)

    while not console.loop_detected:
        console.step()

    return console.acc


class Instruction:
    def __init__(self, instruction: str, operand: int) -> None:
        self.instruction = instruction
        self.operand = operand
        self.executed = False

    def __str__(self) -> str:
        return f"{self.instruction} {self.operand} executed: {self.executed}"


class Console:
    def __init__(self, code: list[Instruction]) -> None:
        self.code = code
        self.acc = 0
        self.pointer = 0

        self.loop_detected = False

    def step(self):
        instruction = self.code[self.pointer]
        if instruction.executed:
            print(f"loop detected, pointer: {self.pointer}, command: {instruction}")
            self.loop_detected = True
            return

        if instruction.instruction == "nop":
            self.pointer += 1
        elif instruction.instruction == "acc":
            self.acc += int(instruction.operand)
            self.pointer += 1
        elif instruction.instruction == "jmp":
            self.pointer += int(instruction.operand)

        instruction.executed = True


if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(solve_first(input))
