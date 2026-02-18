/*
Course Schedule (Topological Sort)

There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that
you must take course bi before course ai. Return true if you can finish all courses,
false if there is a cycle.

Example 1:
  Input:  numCourses = 2, prerequisites = [[1,0]]
  Output: true  (take 0, then 1)

Example 2:
  Input:  numCourses = 2, prerequisites = [[1,0], [0,1]]
  Output: false  (cycle between 0 and 1)

Example 3:
  Input:  numCourses = 4, prerequisites = [[1,0], [2,1], [3,2]]
  Output: true  (take 0 -> 1 -> 2 -> 3)

Constraints:
  - 1 <= numCourses <= 2000
  - 0 <= prerequisites.length <= 5000
  - prerequisites[i].length == 2
  - 0 <= ai, bi < numCourses
  - All prerequisite pairs are unique

Asked by: Amazon, Microsoft, Facebook, Google
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// TODO: implement
	return false
}

func main() {
	type testCase struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}

	cases := []testCase{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{1, [][]int{}, true},
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, true},
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{5, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3}}, true},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("canFinish(%d, %v)", tc.numCourses, tc.prerequisites),
			tc.expected,
			canFinish(tc.numCourses, tc.prerequisites),
		)
	}
}
