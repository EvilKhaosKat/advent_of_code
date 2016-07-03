class Santa:
    x = 0
    y = 0

    houses = {(0, 0)}

    def __init__(self):
        pass

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


santa = Santa()

with open("input.txt") as f:
    for direction in f.read():
        santa.move(direction)

print("houses receive at least one present:" + str(santa.houses_at_least_one_present()))
