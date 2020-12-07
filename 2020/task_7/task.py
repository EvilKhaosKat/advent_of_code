class Bag:

    def __init__(self, name: str) -> None:
        super().__init__()
        self.name = name
        self.contains = set()
        self.contained = set()

    def add_contains(self, bag):
        self.contains.add(bag)

    def add_contained(self, bag):
        self.contained.add(bag)

    def __str__(self) -> str:
        return self.name


def get_bag(bags, bag_name):
    if bag_name in bags:
        return bags[bag_name]
    else:
        bag = Bag(bag_name)
        bags[bag_name] = bag
        return bag


def clean_bag_name(bag_name, clean_number=True):
    bag_name = bag_name.strip()
    if clean_number:
        bag_name = bag_name[bag_name.index(" ") + 1:]
    bag_name = bag_name[:bag_name.rfind("bag")].strip()
    bag_name = bag_name.replace(".", "")

    return bag_name


def solve_first(input: list[str]):
    bags = {}
    for line in input:
        rule = line.split(" contain ")
        bag = get_bag(bags, clean_bag_name(rule[0], False))
        contains = rule[1]

        if contains == "no other bags":
            pass
        else:
            contained = contains.split(",")
            for contained_bag_name in contained:
                contained_bag_name = clean_bag_name(contained_bag_name)
                contained_bag = get_bag(bags, contained_bag_name)

                bag.add_contains(contained_bag)
                contained_bag.add_contained(bag)

    gold_bag = bags["shiny gold"]

    bags_to_check = [gold_bag]
    checked_bags = set()
    while len(bags_to_check) > 0:
        bag = bags_to_check.pop()
        checked_bags.add(bag.name)

        for contained_bag in bag.contained:
            bags_to_check.append(contained_bag)

    return len(checked_bags) - 1



if __name__ == '__main__':
    input = []
    with open("input.txt") as file:
        for line in file:
            input.append(line)

    print(solve_first(input))
