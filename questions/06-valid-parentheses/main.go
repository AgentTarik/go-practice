/*
Valid Parentheses

Given a string containing only '(', ')', '{', '}', '[', and ']', determine if the
input is valid. An input is valid if:
 1. Open brackets are closed by the same type of bracket.
 2. Open brackets are closed in the correct order.
 3. Every close bracket has a corresponding open bracket.

Example 1:

	Input:  "()"
	Output: true

Example 2:

	Input:  "()[]{}"
	Output: true

Example 3:

	Input:  "(]"
	Output: false

Example 4:

	Input:  "([)]"
	Output: false

Example 5:

	Input:  "{[]}"
	Output: true

Constraints:
  - 1 <= len(s) <= 10^4
  - s consists of parentheses only: '()[]{}'
*/
package main

import (
	"fmt"
	"practice/testutil"
)

type stack[T any] []T

func (s *stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v, true
}

func (s *stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	v := (*s)[len(*s)-1]
	return v, true
}

var (
	pairs = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
)

func isValid(s string) bool {
	var myStack stack[rune]
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			myStack.Push(c)
		} else {
			if popped, ok := myStack.Pop(); ok {
				if pairs[popped] != c {
					return false
				}
			} else {
				return false
			}
		}
	}
	if _, ok := myStack.Peek(); ok {
		return false
	}
	return true
}

func main() {
	type testCase struct {
		s        string
		expected bool
	}

	cases := []testCase{
		{"()", true},     // basic
		{"()[]{}", true}, // multiple types
		{"(]", false},    // wrong closing bracket
		{"([)]", false},  // interleaved — classic trap
		{"{[]}", true},   // properly nested
		{"", true},       // empty string
		{"{", false},     // unclosed
		{"}}}", false},   // only closing brackets
		{"]", false},     // starts with closing bracket
		{"((()))", true}, // deep nesting
		{"(()", false},   // one bracket left open
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("isValid(%q)", tc.s),
			tc.expected,
			isValid(tc.s),
		)
	}
}
