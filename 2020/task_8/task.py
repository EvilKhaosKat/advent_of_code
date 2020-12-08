def solve_first(input):
    console = create_console(input)

    while not console.loop_detected:
        console.step()

    return console.acc


def solve_second(input):
    console = create_console(input)

    for idx in range(len(console.code)):
        instruction = console.code[idx]
        if not instruction.fix_corruption():
            continue

        while not console.loop_detected and not console.terminated:
            console.step()
        if console.terminated:
            print(f"fix detected, line {idx} replaced to {console.code[idx].instruction}")
            return console.acc

        console.reset()
        instruction.fix_corruption()

    return console.acc


def create_console(input):
    code = []
    for line in input:
        instruction = line.split(" ")
        code.append(Instruction(instruction[0], instruction[1].strip()))

    return Console(code)


class Instruction:
    def __init__(self, instruction: str, operand: int) -> None:
        self.instruction = instruction
        self.operand = operand
        self.executed = False

    def __str__(self) -> str:
        return f"{self.instruction} {self.operand} executed: {self.executed}"

    def fix_corruption(self) -> False:
        if self.instruction == "jmp":
            self.instruction = "nop"
            return True
        elif self.instruction == "nop":
            self.instruction = "jmp"
            return True
        else:
            return False


class Console:
    acc: int
    pointer: int
    loop_detected: bool
    terminated: bool

    def __init__(self, code: list[Instruction]) -> None:
        self.code = code

        self.reset()

    def reset(self):
        self.loop_detected = False
        self.terminated = False
        self.pointer = 0
        self.acc = 0

        for instruction in self.code:
            instruction.executed = False

    def step(self):
        if self.terminated:
            return

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

        if self.pointer >= len(self.code):
            print("program terminated")
            self.terminated = True

        instruction.executed = True


if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(solve_first(input))
    print(solve_second(input))
