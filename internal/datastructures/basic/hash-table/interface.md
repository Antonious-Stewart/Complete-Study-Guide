interface HashTable<K, V> {
set(key: K, value: V): void;    // Insert or update a key-value pair
get(key: K): V | undefined;     // Retrieve the value associated with the key
delete(key: K): void;           // Remove the key-value pair from the hash table
has(key: K): boolean;           // Check if a key exists in the hash table
size(): number;                 // Return the number of key-value pairs
}