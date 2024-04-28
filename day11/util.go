package day11

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

func locationIsInBounds(row int, col int, current_state state) bool {
	return (row >= 0) && (col >= 0) && (row < len(current_state)) && (col < len(current_state[0]))
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

func buildBlankState(width int, height int) state {
	blank_board := state{}

	for i := 0; i < height; i++ {
		blank_board = append(blank_board, make([]uint8, width))
	}

	return blank_board
}
