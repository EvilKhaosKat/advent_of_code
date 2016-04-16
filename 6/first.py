import unittest


# commands
class Command:
    TURN_OFF = "turn off"
    TURN_ON = "turn on"
    TOGGLE = "toggle"

    type = None
    from_coord = ()
    to_coord = ()

    def __init__(self, from_coord, to_coord):
        super().__init__()
        self.from_coord = from_coord
        self.to_coord = to_coord

    def execute(self, lights_map):
        pass


class TurnOffCommand(Command):
    pass


class TurnOnCommand(Command):
    pass


class ToggleCommand(Command):
    pass
# commands ended


class CommandFactory:
    # turn on 489,959 through 759,964
    @classmethod
    def get_command(cls, command_string):
        from_coord, to_coord = cls.get_coords(command_string)

        return cls.create_command(command_string, from_coord, to_coord)

    @staticmethod
    def create_command(command_string, from_coord, to_coord):
        if command_string.startswith(Command.TURN_OFF):
            return TurnOffCommand(from_coord, to_coord)
        elif command_string.startswith(Command.TURN_ON):
            return TurnOnCommand(from_coord, to_coord)
        elif command_string.startswith(Command.TOGGLE):
            return ToggleCommand(from_coord, to_coord)
        else:
            raise ValueError("Command '{0}' - command type cannot be recognised".format(command_string))

    @classmethod
    def get_coords(cls, command_string):
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

################################################################################
# main logic
# array = [[0] * 1000] * 1000


###############################################################################
# Tests
class TestCommand(unittest.TestCase):
    def test_command_type(self):
        command = CommandFactory.get_command("turn on 489,959 through 759,964")

        self.assertTrue(isinstance(command, TurnOnCommand))
        self.assertEqual(command.from_coord, (489, 959))
        self.assertEqual(command.to_coord, (759, 964))


if __name__ == '__main__':
    unittest.main()
