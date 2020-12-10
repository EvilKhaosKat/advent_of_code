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


def get_contiguous_set(numbers, invalid_number):
    for index_1, number_1 in enumerate(numbers):
        bad_set = [number_1]
        current_sum = number_1
        for index_2 in range(index_1 + 1, len(numbers)-1):
            number_2 = numbers[index_2]

            bad_set.append(number_2)
            current_sum += number_2
            if current_sum > invalid_number:
                bad_set = []
                break
            elif current_sum == invalid_number and len(bad_set) > 1:
                return bad_set

    return []


def solve_second(input, invalid_number):
    numbers = []
    for line in input:
        numbers.append(int(line))

    bad_set = get_contiguous_set(numbers, invalid_number)
    bad_set = sorted(bad_set)

    if len(bad_set) > 2:
        return bad_set[0] + bad_set[-1]

    print("contiguous set not found")
    return -1


if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(f"first: {solve_first(input, 25)}")
    print(f"second: {solve_second(input, solve_first(input, 25))}")
