package day3

import (
	"github.com/hawaite/aoc2020/util"
)

func traverse_grid(grid []string, right_moves int, down_moves int) int {
	hit_trees := 0
	curr_row := 0
	curr_col := 0
	for curr_row < len(grid) {
		if grid[curr_row][curr_col] == '#' {
			hit_trees++
		}
		curr_col = (curr_col + right_moves) % len(grid[curr_row])
		curr_row = curr_row + down_moves
	}

	return hit_trees
}

func Run(lines []string) (string, string) {
	return util.IntPairToStringPair(
		traverse_grid(lines, 3, 1),
		traverse_grid(lines, 1, 1)*traverse_grid(lines, 3, 1)*traverse_grid(lines, 5, 1)*traverse_grid(lines, 7, 1)*traverse_grid(lines, 1, 2))
}
