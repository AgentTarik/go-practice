/*
Island Detection

Given a 2D grid of 1s (land) and 0s (water), count the number of distinct islands.
An island is a group of 1s connected horizontally or vertically (not diagonally).

Example 1:
  Input:
    1 0 1 1
    0 0 1 0
    1 0 0 1
    0 0 0 1
  Output: 4

Example 2:
  Input:
    1 0 1
    1 0 1
    1 1 1
  Output: 1  (U-shape, all connected through the bottom)

Example 3:
  Input:
    1 0 1
    1 1 1
  Output: 1  (two tips joined by a bridge)

Constraints:
  - Grid is non-empty
  - Values are 0 or 1
  - Connectivity is 4-directional (up, down, left, right)
*/
package main

import "practice/testutil"

func islandDetector(ocean [][]int) int {
	innerLength := len(ocean[0])
	length := len(ocean)
	visited := make([][]bool, length)
	for i := range visited {
		visited[i] = make([]bool, innerLength)
	}

	total := 0
	for i := 0; i < length; i++ {
		for j := 0; j < innerLength; j++ {
			if ocean[i][j] == 1 && !visited[i][j] {
				total++
				visited[i][j] = !visited[i][j]
				spreadFunction(ocean, visited, i, j, length, innerLength)
			}
		}
	}
	return total
}

func spreadFunction(ocean [][]int, visited [][]bool, i int, j int, outerLenght int, innerLength int) {
	beforeI := i - 1
	beforeJ := j - 1
	afterI := i + 1
	afterJ := j + 1
	if beforeI >= 0 && !visited[beforeI][j] && ocean[beforeI][j] == 1 {
		visited[beforeI][j] = true
		spreadFunction(ocean, visited, beforeI, j, outerLenght, innerLength)
	}
	if afterI < outerLenght && !visited[afterI][j] && ocean[afterI][j] == 1 {
		visited[afterI][j] = true
		spreadFunction(ocean, visited, afterI, j, outerLenght, innerLength)
	}
	if beforeJ >= 0 && !visited[i][beforeJ] && ocean[i][beforeJ] == 1 {
		visited[i][beforeJ] = true
		spreadFunction(ocean, visited, i, beforeJ, outerLenght, innerLength)
	}
	if afterJ < innerLength && !visited[i][afterJ] && ocean[i][afterJ] == 1 {
		visited[i][afterJ] = true
		spreadFunction(ocean, visited, i, afterJ, outerLenght, innerLength)
	}
}

func main() {
	type testCase struct {
		name     string
		ocean    [][]int
		expected int
	}

	cases := []testCase{
		{"original (4 islands)", [][]int{
			{1, 0, 1, 1},
			{0, 0, 1, 0},
			{1, 0, 0, 1},
			{0, 0, 0, 1},
		}, 4},
		{"U-shape (1 island)", [][]int{
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, 1},
		{"bridge (1 island)", [][]int{
			{1, 0, 1},
			{1, 1, 1},
		}, 1},
		{"three separate islands (3 islands)", [][]int{
			{1, 0, 0},
			{0, 0, 1},
			{1, 0, 0},
		}, 3},
	}

	for _, tc := range cases {
		testutil.Run(tc.name, tc.expected, islandDetector(tc.ocean))
	}
}
