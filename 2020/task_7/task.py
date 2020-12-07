def solve_first(input: list[str]):
    gold_bag = get_gold_bag(input)

    bags_to_check = [gold_bag]
    checked_bags = set()
    while len(bags_to_check) > 0:
        bag = bags_to_check.pop()
        checked_bags.add(bag.name)

        bags_to_check += bag.contained

    return len(checked_bags) - 1


def solve_second(input):
    gold_bag = get_gold_bag(input)

    return get_contains_count(gold_bag) - 1


class Bag:
    def __init__(self, name: str) -> None:
        super().__init__()
        self.name = name
        self.contains = set()
        self.contained = set()

    def add_contains(self, bag, count: int):
        self.contains.add(ContainedBag(bag, count))

    def add_contained(self, bag):
        self.contained.add(bag)

    def __str__(self) -> str:
        return self.name


class ContainedBag:
    def __init__(self, bag, count: int) -> None:
        self.bag = bag
        self.count = count


def get_bag(bags, bag_name):
    if bag_name in bags:
        return bags[bag_name]
    else:
        bag = Bag(bag_name)
        bags[bag_name] = bag
        return bag


def clean_bag_name(bag_name):
    bag_name = bag_name.strip()
    bag_name = bag_name[:bag_name.rfind("bag")].strip()
    bag_name = bag_name.replace(".", "")

    return bag_name


def get_gold_bag(input):
    bags = {}
    for line in input:
        rule = line.split(" contain ")
        bag = get_bag(bags, clean_bag_name(rule[0]))
        contains = rule[1].strip()

        if contains == "no other bags.":
            pass
        else:
            contained = contains.split(",")
            for contained_bag_name in contained:
                contained_bag_name = clean_bag_name(contained_bag_name)
                number_delimiter_index = contained_bag_name.index(" ")

                count = int(contained_bag_name[:number_delimiter_index])
                contained_bag_name = contained_bag_name[number_delimiter_index + 1:]

                contained_bag = get_bag(bags, contained_bag_name)

                bag.add_contains(contained_bag, count)
                contained_bag.add_contained(bag)
    gold_bag = bags["shiny gold"]
    return gold_bag


def get_contains_count(bag: Bag) -> int:
    count = 1
    for contained in bag.contains:
        count += contained.count * get_contains_count(contained.bag)

    return count


if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(solve_first(input))
    print(solve_second(input))
