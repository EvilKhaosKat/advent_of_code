class Wire:
    name = ""
    value = None

    gates = []
    value_gate = None

    def __init__(self, name):
        self.name = name

    def add_gate(self, gate):
        self.gates.append(gate)

    def get_value(self):
        if not self.value:
            value_gate = self.value_gate
            if not value_gate:
                print("problematic wire:" + str(self))
                raise RuntimeError("value_gate is empty for wire " + str(self.name))

            self.value = value_gate.calculate_value()
            if self.value is None:
                print("self.value_gate.value:" + str(self.value_gate.value))
                raise RuntimeError("Calculated value must not be None. self:" + str(self))

        return self.value

    def __str__(self, *args, **kwargs):
        return "Wire(name:%s, value_gate:%s)" % (self.name, self.value_gate)

    def __repr__(self, *args, **kwargs):
        return self.__str__()  # TODO that is not correct version of repr
