package day11

import (
	"github.com/hawaite/aoc2020/util"
)

func generate_cell_value(row int, col int, current_state state, partTwo bool) uint8 {
	// check all 8 immediately adjacent cells to a cell, accounting for being on the edges
	// rule 1: if the cell is an unoccupied seat, and there are zero occupied seats around it, set seat to occupied
	// rule 2: if a cell is an occupied seat, and there are four or more occupied seats adjacent, set seat to unoccupied
	// rule 3: floor has no behaviour
	var adjacent_seat_tolerance int
	var count_immediate_neighbours_only bool
	if !partTwo {
		adjacent_seat_tolerance = 4
		count_immediate_neighbours_only = true
	} else {
		adjacent_seat_tolerance = 5
		count_immediate_neighbours_only = false
	}

	if current_state[row][col] == floor {
		return floor
	} else if current_state[row][col] == unoccupied_chair || current_state[row][col] == occupied_chair {
		adjacent_count := getAdjacentOccupiedSeatCount(row, col, current_state, !count_immediate_neighbours_only)
		if current_state[row][col] == unoccupied_chair && adjacent_count == 0 {
			return occupied_chair
		} else if current_state[row][col] == occupied_chair && adjacent_count >= adjacent_seat_tolerance {
			return unoccupied_chair
		} else {
			return current_state[row][col]
		}
	} else {
		panic("Did not expect value " + string(current_state[row][col]))
	}
}

func checkingDirectionFindsOccupiedSeat(row int, col int, direction uint8, current_state state, recurse bool) int {
	row_to_consider := row + direction_map[direction][0]
	col_to_consider := col + direction_map[direction][1]

	// test the next cell in the requested direction
	// if out of bounds, return 0
	// if occupied seat, return 1
	// if unoccupied seat, return 0
	// if floor, recursively call self with new cooridinates
	if !locationIsInBounds(row_to_consider, col_to_consider, current_state) {
		return 0
	}

	if current_state[row_to_consider][col_to_consider] == occupied_chair {
		return 1
	}

	if current_state[row_to_consider][col_to_consider] == unoccupied_chair {
		return 0
	}

	if !recurse {
		// immediately return 0 for floor
		return 0
	} else {
		// ignore floor and keep checking
		return checkingDirectionFindsOccupiedSeat(row_to_consider, col_to_consider, direction, current_state, true)
	}
}

func getAdjacentOccupiedSeatCount(row int, col int, current_state state, traverse bool) int {
	total := 0
	total += checkingDirectionFindsOccupiedSeat(row, col, top_left, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, top_mid, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, top_right, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, mid_left, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, mid_right, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, bottom_left, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, bottom_mid, current_state, traverse)
	total += checkingDirectionFindsOccupiedSeat(row, col, bottom_right, current_state, traverse)

	return total
}

func calculateNewState(current_state state, partTwo bool) state {
	width := len(current_state[0])
	height := len(current_state)
	new_state := buildBlankState(width, height)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			new_state[row][col] = generate_cell_value(row, col, current_state, partTwo)
		}
	}

	return new_state
}

func Run(lines []string) (string, string) {
	var part1_res int
	var part2_res int

	current_state := buildBoardFromInitialInput(lines)
	state_has_changed := true

	// part 1
	for state_has_changed {
		// make a blank board
		new_state := calculateNewState(current_state, false)

		// see if the new board is equal to the current state
		// if so, set state has changed to false
		state_has_changed = !testStatesAreEqual(new_state, current_state)

		current_state = new_state
	}

	part1_res = countOccupiedSeats(current_state)

	// Part 2
	current_state = buildBoardFromInitialInput(lines)
	state_has_changed = true

	for state_has_changed {
		// make a blank board
		new_state := calculateNewState(current_state, true)
		// see if the new board is equal to the current state
		// if so, set state has changed to false
		state_has_changed = !testStatesAreEqual(new_state, current_state)

		current_state = new_state
	}

	part2_res = countOccupiedSeats(current_state)

	return util.IntPairToStringPair(part1_res, part2_res)
}
