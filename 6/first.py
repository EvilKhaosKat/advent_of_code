import unittest
import itertools


class Command:
    TURN_OFF = "turn off"
    TURN_ON = "turn on"
    TOGGLE = "toggle"

    from_coord = ()
    to_coord = ()

    def __init__(self, from_coord, to_coord):
        super().__init__()
        self.from_coord = from_coord
        self.to_coord = to_coord

    def execute(self, light_object):
        """
        command to be executed for separate light_object, one-by-one
        currently - light object is a number, 0 or 1, represents state of light

        :return new light_object value
        """
        pass

    def __repr__(self, *args, **kwargs):
        return "({0}: from_coord:{1}; to_coord:{2})".format(self.__class__.__name__, self.from_coord, self.to_coord)


class TurnOffCommand(Command):
    def execute(self, light_object):
        super().execute(light_object)

        return 0


class TurnOnCommand(Command):
    def execute(self, light_object):
        super().execute(light_object)

        return 1


class ToggleCommand(Command):
    def execute(self, light_object):
        super().execute(light_object)

        return 1 - light_object


class CommandFactory:
    # turn on 489,959 through 759,964
    @classmethod
    def get_command(cls, command_string):
        from_coord, to_coord = cls._get_coords(command_string)

        return cls._create_command(command_string, from_coord, to_coord)

    @staticmethod
    def _create_command(command_string, from_coord, to_coord):
        if command_string.startswith(Command.TURN_OFF):
            return TurnOffCommand(from_coord, to_coord)
        elif command_string.startswith(Command.TURN_ON):
            return TurnOnCommand(from_coord, to_coord)
        elif command_string.startswith(Command.TOGGLE):
            return ToggleCommand(from_coord, to_coord)
        else:
            raise ValueError("Command '{0}' - command type cannot be recognised".format(command_string))

    @classmethod
    def _get_coords(cls, command_string):
        """
        calculates coordinates from input command string, returns tuple of coordinates - (from, to)
        """
        command_parts = command_string.split(' ')

        to_coords = cls.parse_coords(command_parts[-1])
        from_coords = cls.parse_coords(command_parts[-3])

        return tuple(from_coords), tuple(to_coords)

    @classmethod
    def parse_coords(cls, coords_string):
        """
        transforms "1,2" into ('1', '2')
        """
        coords_list = coords_string.split(",")
        return cls.convert_to_int(coords_list)

    @classmethod
    def convert_to_int(cls, coords_strings_tuple):
        """
        transforms ('1', '2') into (1, 2)
        """
        return tuple(map(int, coords_strings_tuple))


def execute_command(command, lights_map):
    from_coord = command.from_coord
    to_coord = command.to_coord

    for x in range(from_coord[0], to_coord[0] + 1):
        for y in range(from_coord[1], to_coord[1] + 1):
            lights_map[x][y] = command.execute(lights_map[x][y])


def get_commands():
    commands = []

    with open("input.txt") as f:
        for command_string in f.readlines():
            commands.append(CommandFactory.get_command(command_string))

    return commands


def execute_commands():
    for command in commands:
        execute_command(command, lights_map)


def create_lights_map(map_size):
    return [[0] * map_size for i in range(map_size)]


map_size = 1000
lights_map = create_lights_map(map_size)

commands = get_commands()
execute_commands()

all_lights = list(itertools.chain.from_iterable(lights_map))
lit_lights_count = sum(all_lights)
print("lit_lights_count:{0}".format(lit_lights_count))


class TestCommand(unittest.TestCase):
    def test_command_type(self):
        command = CommandFactory.get_command("turn on 489,959 through 759,964")

        self.assertTrue(isinstance(command, TurnOnCommand))
        self.assertEqual(command.from_coord, (489, 959))
        self.assertEqual(command.to_coord, (759, 964))
