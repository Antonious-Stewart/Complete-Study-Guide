package exercises

//1. Sort an Array in Ascending Order
//
//Problem: Given an unsorted array of integers, sort it in ascending order using Bubble Sort.
//Edge Cases:
//Array is already sorted (e.g., [1, 2, 3, 4]).
//Array is in reverse order (e.g., [4, 3, 2, 1]).
//Array contains duplicate elements (e.g., [4, 2, 2, 1]).
//Array has only one element or is empty (e.g., [1] or []).
//
//2. Sort Strings by Length
//
//Problem: Given an array of strings, sort the strings by their length in ascending order using Bubble Sort.
//Edge Cases:
//All strings have the same length (e.g., ["cat", "dog", "bat"]).
//Strings are already sorted by length (e.g., ["a", "bb", "ccc"]).
//The array contains an empty string (e.g., ["", "hello", "world"]).
//Array is empty.
//
//3. Sort Array of Objects by a Key
//
//Problem: Given an array of objects, sort the objects based on a specific numeric key using Bubble Sort. For example:
//
//const data = [{id: 3}, {id: 1}, {id: 2}];
//
//Sort by id in ascending order.
//Edge Cases:
//Objects have the same key value (e.g., [{id: 1}, {id: 1}]).
//Objects are already sorted.
//Array contains only one object or is empty.
//The key is missing in some objects.
//
//4. Find the Kth Largest Element
//
//Problem: Given an unsorted array of integers, find the Kth largest element using Bubble Sort.
//Sort the array in descending order using Bubble Sort, then pick the element at index K-1.
//Edge Cases:
//K is larger than the size of the array.
//Array contains duplicate elements (e.g., [3, 3, 3, 4]).
//K is 1 (return the largest element).
//K is equal to the size of the array (return the smallest element).
