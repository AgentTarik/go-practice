/*
Number of Connected Components in an Undirected Graph (Graph BFS/DFS)

You have n nodes labeled from 0 to n-1 and a list of undirected edges.
Return the number of connected components in the graph.

Example 1:
  Input:  n = 5, edges = [[0,1], [1,2], [3,4]]
  Output: 2  (components: {0,1,2} and {3,4})

Example 2:
  Input:  n = 5, edges = [[0,1], [1,2], [2,3], [3,4]]
  Output: 1  (all nodes connected)

Example 3:
  Input:  n = 4, edges = []
  Output: 4  (each node is its own component)

Constraints:
  - 1 <= n <= 2000
  - 0 <= edges.length <= 5000
  - edges[i].length == 2
  - 0 <= edges[i][0], edges[i][1] < n
  - No duplicate edges

Asked by: Google, Amazon, Facebook, LinkedIn
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func countComponents(n int, edges [][]int) int {
	// TODO: implement
	return 0
}

func main() {
	type testCase struct {
		n        int
		edges    [][]int
		expected int
	}

	cases := []testCase{
		{5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
		{4, [][]int{}, 4},
		{1, [][]int{}, 1},
		{4, [][]int{{0, 1}, {2, 3}}, 2},
		{6, [][]int{{0, 1}, {2, 3}, {4, 5}, {0, 2}}, 2},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("countComponents(%d, %v)", tc.n, tc.edges),
			tc.expected,
			countComponents(tc.n, tc.edges),
		)
	}
}
