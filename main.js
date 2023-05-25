class Stack {
    constructor() {
        this.data = [];
    }

    push(item) {
        this.data.push(item);
    }

    pop() {
        if (!this.isEmpty()) {
            return this.data.pop();
        } else {
            throw new Error('Stack is empty');
        }
    }

    peek() {
        if (!this.isEmpty()) {
            return this.data[this.data.length - 1];
        } else {
            throw new Error('Stack is empty');
        }
    }

    isEmpty() {
        return this.data.length === 0;
    }
}

// Contoh penggunaan
const stack = new Stack();
stack.push(1);
stack.push(2);
stack.push(3);

const item = stack.pop();
console.log('Popped item:', item);

const topItem = stack.peek();
console.log('Top item:', topItem);

console.log('Is stack empty?', stack.isEmpty());
