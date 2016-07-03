class Gate:
    circuit = None

    def __init__(self, circuit):
        self.circuit = circuit

    def parse_line(self, line):
        pass

    def get_wire_by_name(self, name):
        wire = self.circuit.get_wire_by_name(name)
        wire.add_gate(self)

        return wire

    @classmethod
    def create_gate(cls, circuit, line):
        gate = cls._get_gate(cls, circuit)
        gate.parse_line(line)
        return gate

    @classmethod
    def _get_gate(cls, circuit):
        raise NotImplemented("Basic Gate class method '_get_gate' should never be called")


class RawValueGate(Gate):
    value = None
    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        self.value = input_output_raw[0]
        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return RawValueGate(circuit)


class AndGate(Gate):
    first_input = None
    second_input = None

    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" AND ")
        self.first_input = self.get_wire_by_name(input_wires_str[0])
        self.second_input = self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return AndGate(circuit)


class OrGate(Gate):
    first_input = None
    second_input = None

    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" OR ")
        self.first_input = self.get_wire_by_name(input_wires_str[0])
        self.second_input = self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return OrGate(circuit)


class LshiftGate(Gate):
    input_wire = None
    shift_value = None

    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" LSHIFT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[0])
        self.shift_value = self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return LshiftGate(circuit)


class RshiftGate(Gate):
    input_wire = None
    shift_value = None

    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" RSHIFT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[0])
        self.shift_value = self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return RshiftGate(circuit)


class NotGate(Gate):
    input_wire = None
    output_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split("NOT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[-1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])

        return self

    def _get_gate(self, circuit):
        return NotGate(circuit)
