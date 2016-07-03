from AdventOfCodeHelper import get_input_lines


class Circuit:
    wires = {}

    def read_schema(self):
        for line in get_input_lines():
            self._parse_line(line)

    def _parse_line(self, line):
        if "AND" in line:
            pass
        elif "OR" in line:
            pass
        elif "NOT" in line:
            pass
        elif "LSHIFT" in line:
            pass
        elif "RSHIFT" in line:
            pass
        else:  # raw value to the gate situation


circuit = Circuit()
