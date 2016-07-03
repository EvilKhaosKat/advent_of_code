from task_7.Wires import Wire
from task_7.Gates import *

from AdventOfCodeHelper import get_input_lines


class Circuit:
    wires = {}

    def read_circuit(self):
        for line in get_input_lines():
            self._parse_line(line)

    def _parse_line(self, line):
        if "AND" in line:
            self._add_and_gate(line)
        elif "OR" in line:
            pass
        elif "NOT" in line:
            pass
        elif "LSHIFT" in line:
            pass
        elif "RSHIFT" in line:
            pass
        else:  # raw value to the gate situation
            pass

    def get_wire_by_name(self, name):
        wire = self.wires.get(name)

        if not wire:
            wire = Wire(name)
            self.wires[name] = wire

        return wire

    def add_wire(self, wire):
        self.wires[wire.name] = wire

    def _add_and_gate(self, line):
        """
        '1 AND ht -> hu'

        :return: None
        """

        and_gate = AndGate.create_gate(self, line)


circuit = Circuit()
circuit.read_circuit()
