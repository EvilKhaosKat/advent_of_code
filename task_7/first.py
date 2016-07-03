from task_7.Wires import Wire
from task_7.Gates import *

from AdventOfCodeHelper import get_input_lines


class Circuit:
    wires = {}

    def read_circuit(self):
        for line in get_input_lines():
            self._parse_line(line.strip())

    def _parse_line(self, line):
        if "AND" in line:
            AndGate.create_gate(self, line)
        elif "OR" in line:
            OrGate.create_gate(self, line)
        elif "NOT" in line:
            NotGate.create_gate(self, line)
        elif "LSHIFT" in line:
            LshiftGate.create_gate(self, line)
        elif "RSHIFT" in line:
            RshiftGate.create_gate(self, line)
        else:  # raw value to the gate situation
            RawValueGate.create_gate(self, line)

    def get_wire_by_name(self, name):
        wire = self.wires.get(name)

        if not wire:
            wire = Wire(name)
            self.wires[name] = wire

        return wire

    def add_wire(self, wire):
        self.wires[wire.name] = wire

    def __str__(self, *args, **kwargs):
        return "Circuit(wires:" + str(self.wires) + ")"


circuit = Circuit()
circuit.read_circuit()
print("circuit:" + str(circuit))