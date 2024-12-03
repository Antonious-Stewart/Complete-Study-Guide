package slidingWindow

//1. Maximum Sum of Subarray of Size K
//
//Problem: Given an array of integers and an integer K, find the maximum sum of any contiguous subarray of size K.
//Edge Cases:
//The array length is less than K.
//The array contains negative numbers.
//K is 1 (single element subarray).
//The array has all elements equal (e.g., [3, 3, 3, 3]).
//
//2. Longest Substring Without Repeating Characters
//
//Problem: Given a string, find the length of the longest substring without repeating characters.
//Edge Cases:
//The string is empty.
//All characters in the string are the same (e.g., "aaaaaa").
//The string has no repeating characters (e.g., "abcdef").
//The string has repeating characters at the beginning and end (e.g., "abcabcbb").
//
//3. Smallest Subarray with a Given Sum
//
//Problem: Given an array of positive integers and a target sum S, find the length of the smallest contiguous subarray whose sum is greater than or equal to S. If no such subarray exists, return 0.
//Edge Cases:
//No subarray meets the target sum.
//The target sum S is greater than the sum of the entire array.
//The entire array is the smallest subarray (e.g., [1, 2, 3] with S = 6).
//The array contains a single element that meets or exceeds S.
//
//4. Longest Subarray with Ones after Replacing at Most K Zeros
//
//Problem: Given a binary array, find the length of the longest subarray containing 1s after replacing at most K 0s.
//Edge Cases:
//The array contains only 0s or only 1s.
//K is 0 (no replacements allowed).
//The array length is smaller than K.
//All zeros are at one end of the array (e.g., [1, 1, 1, 0, 0] with K = 2)
