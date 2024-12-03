interface Stack {
push(item: T): void;      // Add an element to the top of the stack
pop(): T | undefined;     // Remove and return the top element of the stack
peek(): T | undefined;    // Return the top element without removing it
isEmpty(): boolean;       // Check if the stack is empty
size(): number;           // Return the number of elements in the stack
}