interface Queue<T> {
enqueue(item: T): void;      // Add an element to the back of the queue
dequeue(): T | undefined;    // Remove and return the front element of the queue
front(): T | undefined;      // Return the front element without removing it
isEmpty(): boolean;          // Check if the queue is empty
size(): number;              // Return the number of elements in the queue
}