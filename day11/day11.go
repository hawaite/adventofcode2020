package day11

import (
	"github.com/hawaite/aoc2020/util"
)

type state [][]uint8

const (
	floor            uint8 = iota
	unoccupied_chair uint8 = iota
	occupied_chair   uint8 = iota
)

func buildBlankState(width int, height int) state {
	blank_board := state{}

	for i := 0; i < height; i++ {
		blank_board = append(blank_board, make([]uint8, width))
	}

	return blank_board
}

func buildBoardFromInitialInput(initial_input []string) state {
	new_state := buildBlankState(len(initial_input[0]), len(initial_input))

	for i := 0; i < len(new_state); i++ {
		for j := 0; j < len(new_state[i]); j++ {
			value_from_cell := initial_input[i][j]
			switch value_from_cell {
			case 'L':
				new_state[i][j] = unoccupied_chair
			case '#':
				new_state[i][j] = occupied_chair
			case '.':
				new_state[i][j] = floor
			default:
				panic("Did not expect input value of:" + string(value_from_cell))
			}
		}
	}

	return new_state
}

// func print_board_state(current_state state) {
// 	for i := 0; i < len(current_state); i++ {
// 		for j := 0; j < len(current_state[i]); j++ {
// 			switch current_state[i][j] {
// 			case floor:
// 				fmt.Print(".")
// 			case unoccupied_chair:
// 				fmt.Print("L")
// 			case occupied_chair:
// 				fmt.Print("#")
// 			default:
// 				panic("Did not expect input value of:" + string(current_state[i][j]))
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }

func getAdjacentOccupiedSeatCount(row int, col int, current_state state) int {
	height := len(current_state)
	width := len(current_state[0])
	total := 0
	//TL
	if col > 0 && row > 0 {
		if current_state[row-1][col-1] == occupied_chair {
			total += 1
		}
	}
	//TM
	if row > 0 {
		if current_state[row-1][col] == occupied_chair {
			total += 1
		}
	}
	//TR
	if col < width-1 && row > 0 {
		if current_state[row-1][col+1] == occupied_chair {
			total += 1
		}
	}
	//ML
	if col > 0 {
		if current_state[row][col-1] == occupied_chair {
			total += 1
		}
	}
	//MR
	if col < width-1 {
		if current_state[row][col+1] == occupied_chair {
			total += 1
		}
	}
	//BL
	if row < height-1 && col > 0 {
		if current_state[row+1][col-1] == occupied_chair {
			total += 1
		}
	}
	//BM
	if row < height-1 {
		if current_state[row+1][col] == occupied_chair {
			total += 1
		}
	}
	//BR
	if row < height-1 && col < width-1 {
		if current_state[row+1][col+1] == occupied_chair {
			total += 1
		}
	}

	return total
}

func generate_cell_value(row int, col int, current_state state) uint8 {
	// check all 8 directions on cell, accounting for being on the edges
	// rule 1: if the cell is an unoccupied seat, and there are zero occupied seats around it, set seat to occupied
	// rule 2: if a cell is an occupied seat, and there are four or more occupied seats adjacent, set seat to unoccupied
	// rule 3: floor has no behaviour
	if current_state[row][col] == floor {
		return floor
	} else if current_state[row][col] == unoccupied_chair {
		adjacent_count := getAdjacentOccupiedSeatCount(row, col, current_state)
		if adjacent_count == 0 {
			return occupied_chair
		} else {
			return unoccupied_chair
		}
	} else if current_state[row][col] == occupied_chair {
		adjacent_count := getAdjacentOccupiedSeatCount(row, col, current_state)
		if adjacent_count >= 4 {
			return unoccupied_chair
		} else {
			return occupied_chair
		}
	} else {
		panic("Did not expect value " + string(current_state[row][col]))
	}
}

func testStatesAreEqual(stateOne state, stateTwo state) bool {
	if len(stateOne) != len(stateTwo) {
		return false
	}

	for row := 0; row < len(stateOne); row++ {
		if len(stateOne[row]) != len(stateTwo[row]) {
			return false
		}

		for col := 0; col < len(stateOne[row]); col++ {
			if stateOne[row][col] != stateTwo[row][col] {
				return false
			}
		}
	}
	return true
}

func countOccupiedSeats(current_state state) int {
	total := 0

	for row := 0; row < len(current_state); row++ {
		for col := 0; col < len(current_state[row]); col++ {
			if current_state[row][col] == occupied_chair {
				total++
			}
		}
	}

	return total
}

func Run(lines []string) (string, string) {
	var part1_res int
	var part2_res int

	current_state := buildBoardFromInitialInput(lines)
	state_has_changed := true
	height := len(current_state)
	width := len(current_state[0])

	iteration_count := 0
	for state_has_changed {
		// make a blank board
		new_state := buildBlankState(width, height)

		// populate it
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				new_state[row][col] = generate_cell_value(row, col, current_state)
			}
		}
		// see if the new board is equal to the current state
		// if so, set state has changed to false
		state_has_changed = !testStatesAreEqual(new_state, current_state)

		current_state = new_state
		iteration_count++
	}

	part1_res = countOccupiedSeats(current_state)

	return util.IntPairToStringPair(part1_res, part2_res)
}
