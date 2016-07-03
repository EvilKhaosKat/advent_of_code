class Wire:
    name = ""
    gates = []

    def __init__(self, name):
        self.name = name

    def add_gate(self, gate):
        self.gates.append(gate)

    def __str__(self, *args, **kwargs):
        return "Wire(name:%s)" % self.name

    def __repr__(self, *args, **kwargs):
        return self.__str__()  # TODO that is not correct version of repr
