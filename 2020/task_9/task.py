def is_valid(number, numbers):
    for index_1, number_1 in enumerate(numbers):
        for index_2 in range(index_1, len(numbers)):
            number_2 = numbers[index_2]

            if number_1 + number_2 == number:
                return True

    return False


def solve_first(input, preamble_size: int):
    numbers = []
    for line in input:
        number = int(line)
        if len(numbers) < preamble_size:
            numbers.append(number)
        else:
            if is_valid(number, numbers):
                numbers = numbers[1:]
                numbers.append(number)
            else:
                return number

    print("invalid number not found")
    return -1


if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(f"first: {solve_first(input, 25)}")
