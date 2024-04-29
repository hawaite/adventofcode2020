package day12

import (
	"fmt"
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

type direction uint8

const (
	north direction = iota
	south direction = iota
	east  direction = iota
	west  direction = iota
)

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

// map of direction to an x/y transform used to navigate that direction
var direction_map = map[direction][]int{
	north: {0, 1},
	south: {0, -1},
	east:  {1, 0},
	west:  {-1, 0},
}

type ship struct {
	x   int
	y   int
	dir direction
}

func Run(lines []string) (string, string) {
	var part1_res string
	var part2_res string
	var ship ship
	ship.x = 0
	ship.y = 0
	ship.dir = east

	for _, line := range lines {
		dir_part := string(line[0])
		magnitude, err := strconv.Atoi(line[1:])
		util.ErrCheck(err)

		switch dir_part {
		case "F":
			direction_modifier := direction_map[ship.dir]
			ship.x += (direction_modifier[0] * magnitude)
			ship.y += (direction_modifier[1] * magnitude)
		case "L":
			degrees_without_full_rotations := magnitude % 360
			left_rotations := degrees_without_full_rotations / 90
			for i := 0; i < left_rotations; i++ {
				ship.dir = getNextLeftDirection(ship.dir)
			}
		case "R":
			degrees_without_full_rotations := magnitude % 360
			right_rotations := degrees_without_full_rotations / 90
			for i := 0; i < right_rotations; i++ {
				ship.dir = getNextRightDirection(ship.dir)
			}
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			direction_modifier := direction_map[directionLetterToDirection(dir_part)]
			ship.x += (direction_modifier[0] * magnitude)
			ship.y += (direction_modifier[1] * magnitude)
		default:
			panic("Unexpected direction: " + dir_part)
		}
	}

	manhattan_distance := intAbs(ship.x) + intAbs(ship.y)
	part1_res = fmt.Sprintf("%d", manhattan_distance)

	return part1_res, part2_res
}
