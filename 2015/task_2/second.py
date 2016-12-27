def get_smallest_perimeter(l, w, h):
    sizes = [l, w, h]

    max_size = max(sizes)
    sizes.remove(max_size)
    return sizes[0] * 2 + sizes[1] * 2


def calculate_ribbon(gift):
    l, w, h = map(int, gift.split("x"))

    return get_volume(h, l, w) + get_smallest_perimeter(l, w, h)


def get_volume(l, w, h):
    return l * w * h


total = 0
with open("input.txt") as f:
    for gift in f.readlines():
        total += calculate_ribbon(gift)

print(total)
