class Santa:
    x = 0
    y = 0

    houses = None

    def __init__(self, houses):
        self.houses = houses

    def _make_house_coordinates(self):
        return self.x, self.y

    def move(self, direction):
        if direction == "^":
            self.y += 1
        elif direction == "v":
            self.y -= 1
        elif direction == "<":
            self.x -= 1
        elif direction == ">":
            self.x += 1

        house_coordinates = self._make_house_coordinates()
        if house_coordinates not in self.houses:
            self.houses.add(house_coordinates)

    def houses_at_least_one_present(self):
        return len(self.houses)


houses = {(0, 0)}  # common set between two santas

santa = Santa(houses)
roboSanta = Santa(houses)

with open("input.txt") as f:
    isSanta = True

    for direction in f.read():
        if isSanta:
            santa.move(direction)
            isSanta = False
        else:
            roboSanta.move(direction)
            isSanta = True

print("houses receive at least one present:" + str(santa.houses_at_least_one_present()))
