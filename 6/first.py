import unittest

class Command:
    TURN_OFF = "turn off"
    TURN_ON = "turn on"
    TOGGLE = "toggle"

    type = None
    from_coord = ()
    to_coord = ()

    # turn on 489,959 through 759,964
    def __init__(self, command_string):
        self.type = Command.get_command_type(command_string)
        self.from_coord, self.to_coord = Command.get_coords(command_string)

    @classmethod
    def get_command_type(cls, command_string):
        if command_string.startswith(cls.TURN_OFF):
            return cls.TURN_OFF
        elif command_string.startswith(cls.TURN_ON):
            return cls.TURN_ON
        elif command_string.startswith(cls.TOGGLE):
            return cls.TOGGLE
        else:
            raise ValueError("Command '{0}' - command type cannot be recognised".format(command_string))

    @classmethod
    def get_coords(cls, command_string):
        command_parts = command_string.split(' ')

        to_coords = cls.parse_coords(command_parts[-1])
        from_coords = cls.parse_coords(command_parts[-3])

        #print("to_coords:" + str(to_coords))
        #print("from_coords:" + str(from_coords))

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


#array = [[0] * 1000] * 1000


###############################################################################
# Tests
class TestCommand(unittest.TestCase):
    def test_command_type(self):
        command = Command("turn on 489,959 through 759,964")

        self.assertEqual(command.type, Command.TURN_ON)
        self.assertEqual(command.from_coord, (489, 959))
        self.assertEqual(command.to_coord, (759, 964))


if __name__ == '__main__':
    unittest.main()
