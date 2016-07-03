def get_smallest_area(l, w, h):
    sizes = [l, w, h]

    sizes.remove(max(sizes))

    return sizes[0] * sizes[1]


def calculate_paper(gift):
    l, w, h = map(int, gift.split("x"))

    return get_surface_area(h, l, w) + get_smallest_area(l, w, h)


def get_surface_area(l, w, h):
    return 2 * l * w + 2 * w * h + 2 * h * l


total = 0
with open("input.txt") as f:
    total = sum(calculate_paper(gift) for gift in f.readlines())

print(total)
