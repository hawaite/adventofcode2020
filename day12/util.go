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

func rotate_waypoint_ccw(waypoint *point, degrees int) {
	original_waypoint_x := waypoint.x
	original_waypoint_y := waypoint.y
	if degrees == 90 {
		waypoint.x = original_waypoint_y * -1
		waypoint.y = original_waypoint_x
	} else if degrees == 180 {
		waypoint.x = original_waypoint_x * -1
		waypoint.y = original_waypoint_y * -1
	} else if degrees == 270 {
		waypoint.x = original_waypoint_y
		waypoint.y = original_waypoint_x * -1
	} else {
		panic("Rotation degrees not 90, 180, or 270")
	}
}

func rotate_waypoint_cw(waypoint *point, degrees int) {
	original_waypoint_x := waypoint.x
	original_waypoint_y := waypoint.y
	if degrees == 270 {
		waypoint.x = original_waypoint_y * -1
		waypoint.y = original_waypoint_x
	} else if degrees == 180 {
		waypoint.x = original_waypoint_x * -1
		waypoint.y = original_waypoint_y * -1
	} else if degrees == 90 {
		waypoint.x = original_waypoint_y
		waypoint.y = original_waypoint_x * -1
	} else {
		panic("Rotation degrees not 90, 180, or 270")
	}
}
