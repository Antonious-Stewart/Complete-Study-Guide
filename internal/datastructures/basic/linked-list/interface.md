interface LinkedList<T> {
insertAtHead(item: T): void;   // Insert an element at the beginning of the list
insertAtTail(item: T): void;   // Insert an element at the end of the list
delete(item: T): void;         // Remove an element from the list
find(item: T): T | undefined;  // Find and return an element from the list
isEmpty(): boolean;            // Check if the list is empty
size(): number;                // Return the number of elements in the list
}