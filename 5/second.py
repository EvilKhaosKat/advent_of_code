def is_at_least_three_vowels(string):
    vowels_count = 0

    for char in string:
        if char in "aeiou":
            vowels_count += 1

    return vowels_count >= 3


def is_contains_twice_in_row(string):
    for position in range(len(string)-1):
        if string[position] == string[position+1]:
            return True

    return False


def is_no_prohibited_strings(string):
    if "ab" not in string and "cd" not in string and "pq" not in string and "xy" not in string:
        return True
    else:
        return False


def is_string_nice(string):
    if is_at_least_three_vowels(string) and is_contains_twice_in_row(string) and is_no_prohibited_strings(string):
        return True
    else:
        return False


with open("input.txt") as f:
    nice_strings = sum(1 if is_string_nice(string) else 0 for string in f.readlines())

print("print_strings:{0}".format(nice_strings))
