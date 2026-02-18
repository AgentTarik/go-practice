/*
Longest Substring Without Repeating Characters

Given a string, return the length of the longest substring that contains no
repeating characters.

Example 1:

	Input:  "abcabcbb"
	Output: 3  ("abc")

Example 2:

	Input:  "bbbbb"
	Output: 1  ("b")

Example 3:

	Input:  "pwwkew"
	Output: 3  ("wke")

Constraints:
  - 0 <= len(s) <= 5 * 10^4
  - s consists of English letters, digits, symbols and spaces

Asked by: Google, Amazon, Microsoft, Stripe, Facebook
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[rune]int)
	longestLength := 0
	acceptableIndex := 0

	for i, c := range s {
		if restoredIndex, ok := indexMap[c]; ok {
			acceptableIndex = max(restoredIndex+1, acceptableIndex)
		}
		longestLength = max(i-acceptableIndex+1, longestLength)
		indexMap[c] = i
	}

	return longestLength
}

func main() {
	type testCase struct {
		s        string
		expected int
	}

	cases := []testCase{
		{"abcabcbb", 3}, // classic — window slides on repeated 'a', 'b', 'c'
		{"bbbbb", 1},    // all same character
		{"pwwkew", 3},   // answer is not at the start
		{"", 0},         // empty string
		{"abba", 2},     // palindrome — trap for naive approaches
		{"dvdf", 3},     // "vdf" — left pointer must jump past old 'd'
		{"tmmzuxt", 5},  // "mzuxt" — interior repeat early on
		{" ", 1},        // single space counts as a character
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("lengthOfLongestSubstring(%q)", tc.s),
			tc.expected,
			lengthOfLongestSubstring(tc.s),
		)
	}
}
