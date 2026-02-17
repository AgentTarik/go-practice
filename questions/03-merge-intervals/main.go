/*
Merge Intervals

Given an array of intervals [start, end], merge all overlapping intervals and
return an array of the non-overlapping intervals that cover all intervals in the input.
Intervals that share an endpoint are considered overlapping.

Example 1:
  Input:  [[1,3],[2,6],[8,10],[15,18]]
  Output: [[1,6],[8,10],[15,18]]

Example 2:
  Input:  [[1,4],[4,5]]
  Output: [[1,5]]  (touching endpoints merge)

Example 3:
  Input:  [[1,4],[2,3]]
  Output: [[1,4]]  (one interval fully contained in another)

Constraints:
  - 1 <= intervals.length <= 10^4
  - 0 <= start <= end <= 10^4

Asked by: Google, Amazon, Microsoft, Bloomberg
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func merge(intervals [][2]int) [][2]int {
	// TODO: implement
	return nil
}

func main() {
	type testCase struct {
		intervals [][2]int
		expected  [][2]int
	}

	cases := []testCase{
		{
			[][2]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][2]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][2]int{{1, 4}, {4, 5}},
			[][2]int{{1, 5}}, // touching endpoints merge
		},
		{
			[][2]int{{1, 4}, {2, 3}},
			[][2]int{{1, 4}}, // fully contained inside another
		},
		{
			[][2]int{{1, 2}, {3, 4}, {5, 6}},
			[][2]int{{1, 2}, {3, 4}, {5, 6}}, // no overlaps — output equals input
		},
		{
			[][2]int{{1, 4}},
			[][2]int{{1, 4}}, // single interval
		},
		{
			[][2]int{{0, 4}, {1, 3}, {2, 5}},
			[][2]int{{0, 5}}, // chain merge — all collapse into one
		},
	}

	for _, tc := range cases {
		testutil.RunDeep(
			fmt.Sprintf("merge(%v)", tc.intervals),
			tc.expected,
			merge(tc.intervals),
		)
	}
}
