class Stack:
    def __init__(self):
        self.data = []

    def push(self, item):
        self.data.append(item)

    def pop(self):
        if not self.is_empty():
            return self.data.pop()
        else:
            raise Exception("Stack is empty")

    def peek(self):
        if not self.is_empty():
            return self.data[-1]
        else:
            raise Exception("Stack is empty")

    def is_empty(self):
        return len(self.data) == 0



# Contoh penggunaan
stack = Stack()
stack.push(1)
stack.push(2)
stack.push(3)

item = stack.pop()
print("Popped item:", item)

item = stack.peek()
print("Top item:", item)

print("Is stack empty?", stack.is_empty())
