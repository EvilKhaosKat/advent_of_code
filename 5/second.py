# first check
def contains_repeated_pairs(string):
    """
    It contains a pair of any two letters that appears at least twice in the string without overlapping,
    like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
    """
    import re

    return bool(re.search(r"(..).*\1", string))


# second check
def contains_repeated_letter(string):
    """
    It contains at least one letter which repeats with exactly one letter between them,
    like xyx, abcdefeghi (efe), or even aaa.
    """

    for letter_number in range(0, len(string) - 2):
        first_letter = string[letter_number]
        third_letter = string[letter_number + 2]

        if first_letter == third_letter:
            return True

    return False


def is_string_nice(string):
    return contains_repeated_pairs(string) and contains_repeated_letter(string)


# main logic
with open("input.txt") as f:
    nice_strings = sum(1 if is_string_nice(string) else 0 for string in f.readlines())

print("nice_strings:{0}".format(nice_strings))