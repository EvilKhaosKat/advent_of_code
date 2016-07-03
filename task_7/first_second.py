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
        if not isinstance(name, str):
            raise RuntimeError("Attempt to get wire not by string name:" + str(name))

        if name.isdigit():
            raise RuntimeError("Attempt to get wire by number name: " + str(name))

        wire = self.wires.get(name)

        if not wire:
            wire = Wire(name)
            self.wires[name] = wire

        return wire

    def add_wire(self, wire):
        self.wires[wire.name] = wire

    def reset_wires_values(self):
        for wire in self.wires.values():
            wire.value = None

    def __str__(self, *args, **kwargs):
        return "Circuit(wires:" + str(self.wires) + ")"


circuit = Circuit()
circuit.read_circuit()
print("circuit:" + str(circuit))

wire_lw = circuit.wires["lw"]
print("wire_lw:" + str(wire_lw))


wire_a = circuit.wires["a"]

a_value = wire_a.get_value()
print("first a value:" + str(a_value))

circuit.reset_wires_values()
circuit.wires["b"].value = a_value

a_value = wire_a.get_value()
print("second a value:" + str(a_value))
