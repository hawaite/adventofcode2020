package day12

func directionLetterToDirection(direction_letter string) direction {
	switch direction_letter {
	case "N":
		return north
	case "S":
		return south
	case "E":
		return east
	case "W":
		return west
	default:
		panic("Unexpected direction: " + direction_letter)
	}
}

func getNextLeftDirection(current_direction direction) direction {
	switch current_direction {
	case north:
		return west
	case west:
		return south
	case south:
		return east
	case east:
		return north
	default:
		panic("Unexpected direction: " + string(current_direction))
	}
}

func getNextRightDirection(current_direction direction) direction {
	switch current_direction {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	default:
		panic("Unexpected direction: " + string(current_direction))
	}
}

func intAbs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}
