# TODO huge amount of similar code to be removed
class Gate:
    circuit = None
    output_wire = None

    def __init__(self, circuit):
        self.circuit = circuit

    def parse_line(self, line):
        pass

    def get_wire_by_name(self, name):
        wire = self.circuit.get_wire_by_name(name)
        wire.add_gate(self)

        return wire

    def calculate_value(self):
        raise NotImplemented("Basic Gate class method 'calculate_value' should never be called")

    @classmethod
    def create_gate(cls, circuit, line):
        gate = cls._get_gate(cls, circuit)
        gate.parse_line(line)
        return gate

    @classmethod
    def _get_gate(cls, circuit):
        raise NotImplemented("Basic Gate class method '_get_gate' should never be called")

        # def __str__(self, *args, **kwargs):
        #     return "%s(output_wire:%s)" % (self.__class__.__name__, self.output_wire)
        #
        # def __repr__(self, *args, **kwargs):
        #     return self.__str__()


class RawValueGate(Gate):
    value = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        self.value = input_output_raw[0]
        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return RawValueGate(circuit)

    def calculate_value(self):
        if self.value.isdigit():
            return int(self.value)
        else:
            return self.circuit.get_wire_by_name(self.value).get_value()


class AndGate(Gate):
    first_input = None
    second_input = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" AND ")

        self.first_input = input_wires_str[0]  # self.get_wire_by_name(input_wires_str[0])
        self.second_input = input_wires_str[1]  # self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return AndGate(circuit)

    def calculate_value(self):
        second_wire = self.circuit.get_wire_by_name(self.second_input)
        # TODO supports 1 special case - 'Digit AND Wire' - to be updated to support all variants, for all wires types
        if self.first_input.isdigit():
            return int(self.first_input) & second_wire.get_value()
        else:
            first_wire = self.circuit.get_wire_by_name(self.first_input)
            return first_wire.get_value() & second_wire.get_value()


class OrGate(Gate):
    first_input = None
    second_input = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" OR ")
        self.first_input = input_wires_str[0]  # self.get_wire_by_name(input_wires_str[0])
        self.second_input = input_wires_str[1]  # self.get_wire_by_name(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return OrGate(circuit)

    def calculate_value(self):
        first_wire = self.circuit.get_wire_by_name(self.first_input)
        second_wire = self.circuit.get_wire_by_name(self.second_input)

        return first_wire.get_value() | second_wire.get_value()


class LshiftGate(Gate):
    input_wire = None
    shift_value = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" LSHIFT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[0])
        self.shift_value = int(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return LshiftGate(circuit)

    def calculate_value(self):
        input_wire_obj = self.circuit.get_wire_by_name(self.input_wire.name)

        return input_wire_obj.get_value() << self.shift_value


class RshiftGate(Gate):
    input_wire = None
    shift_value = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split(" RSHIFT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[0])
        self.shift_value = int(input_wires_str[1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return RshiftGate(circuit)

    def calculate_value(self):
        input_wire_obj = self.circuit.get_wire_by_name(self.input_wire.name)

        return input_wire_obj.get_value() >> self.shift_value


class NotGate(Gate):
    input_wire = None

    def parse_line(self, line):
        input_output_raw = line.split(" -> ")

        input_wires_str = input_output_raw[0].split("NOT ")
        self.input_wire = self.get_wire_by_name(input_wires_str[-1])

        self.output_wire = self.get_wire_by_name(input_output_raw[-1])
        self.output_wire.value_gate = self

        return self

    def _get_gate(self, circuit):
        return NotGate(circuit)

    def calculate_value(self):
        input_wire_obj = self.circuit.get_wire_by_name(self.input_wire.name)

        return ~ input_wire_obj.get_value()
