package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	board := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.TrimSuffix(line, "\n"))
	}

	fmt.Println("Part 1 tree count:", traverse_grid(board, 3, 1))
	fmt.Println(
		"Part 2 tree count:",
		traverse_grid(board, 1, 1)*traverse_grid(board, 3, 1)*traverse_grid(board, 5, 1)*traverse_grid(board, 7, 1)*traverse_grid(board, 1, 2))
}
