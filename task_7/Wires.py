class Wire:
    name = ""
    gates = []

    def __init__(self, name):
        self.name = name

    def add_gate(self, gate):
        self.gates.append(gate)
