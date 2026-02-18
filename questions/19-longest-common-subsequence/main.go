/*
Longest Common Subsequence

Given two strings text1 and text2, return the length of their longest common subsequence.
A subsequence is a sequence that can be derived from the string by deleting some
(or no) characters without changing the order of the remaining characters.
If there is no common subsequence, return 0.

Example 1:
  Input:  text1 = "abcde", text2 = "ace"
  Output: 3  ("ace")

Example 2:
  Input:  text1 = "abc", text2 = "abc"
  Output: 3

Example 3:
  Input:  text1 = "abc", text2 = "def"
  Output: 0

Constraints:
  - 1 <= text1.length, text2.length <= 1000
  - text1 and text2 consist of only lowercase English characters

Asked by: Amazon, Google, Facebook, Bloomberg
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	// TODO: implement
	return 0
}

func main() {
	type testCase struct {
		text1    string
		text2    string
		expected int
	}

	cases := []testCase{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"oxcpqrsvwf", "shmtulqrypy", 2},
		{"", "abc", 0},
		{"a", "a", 1},
		{"abcba", "abcbcba", 5},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("LCS(%q, %q)", tc.text1, tc.text2),
			tc.expected,
			longestCommonSubsequence(tc.text1, tc.text2),
		)
	}
}
