package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

type direction uint8

const (
	north direction = iota
	south direction = iota
	east  direction = iota
	west  direction = iota
)

// map of direction to an x/y transform used to navigate that direction
var direction_map = map[direction][]int{
	north: {0, 1},
	south: {0, -1},
	east:  {1, 0},
	west:  {-1, 0},
}

type point struct {
	x int
	y int
}

type pointWithDirection struct {
	coord point
	dir   direction
}

func performPartOneMove(ship *pointWithDirection, instruction string, magnitude int) {
	if !strings.Contains("FLRNSEW", instruction) {
		panic("Unexpected instruction: " + instruction)
	}

	if instruction == "F" {
		direction_modifier := direction_map[ship.dir]
		ship.coord.x += (direction_modifier[0] * magnitude)
		ship.coord.y += (direction_modifier[1] * magnitude)
	} else if instruction == "L" || instruction == "R" {
		degrees_without_full_rotations := magnitude % 360
		ninety_degree_rotations := degrees_without_full_rotations / 90
		for i := 0; i < ninety_degree_rotations; i++ {
			if instruction == "L" {
				ship.dir = getNextLeftDirection(ship.dir)
			} else {
				ship.dir = getNextRightDirection(ship.dir)
			}
		}
	} else { // N S E W
		direction_modifier := direction_map[directionLetterToDirection(instruction)]
		ship.coord.x += (direction_modifier[0] * magnitude)
		ship.coord.y += (direction_modifier[1] * magnitude)
	}
}

func performPartTwoMove(ship *pointWithDirection, waypoint *point, instruction string, magnitude int) {
	if !strings.Contains("FLRNSEW", instruction) {
		panic("Unexpected instruction: " + instruction)
	}

	if instruction == "F" {
		x_jump := magnitude * waypoint.x
		y_jump := magnitude * waypoint.y

		ship.coord.x += x_jump
		ship.coord.y += y_jump
		// no need to update waypoint as it stays the same relative location
	} else if instruction == "L" {
		degrees_without_full_rotations := magnitude % 360
		rotate_waypoint_ccw(waypoint, degrees_without_full_rotations)
	} else if instruction == "R" {
		degrees_without_full_rotations := magnitude % 360
		rotate_waypoint_cw(waypoint, degrees_without_full_rotations)
	} else { // N S E W
		direction_modifier := direction_map[directionLetterToDirection(instruction)]
		waypoint.x += (direction_modifier[0] * magnitude)
		waypoint.y += (direction_modifier[1] * magnitude)
	}
}

func Run(lines []string) (string, string) {
	var part1_res string
	var part2_res string

	var partOneShip pointWithDirection
	partOneShip.coord.x = 0
	partOneShip.coord.y = 0
	partOneShip.dir = east

	var partTwoShip pointWithDirection
	partTwoShip.coord.x = 0
	partTwoShip.coord.y = 0
	partTwoShip.dir = east

	var partTwoWaypoint point
	partTwoWaypoint.x = 10
	partTwoWaypoint.y = 1

	for _, line := range lines {
		instruction := string(line[0])
		magnitude, err := strconv.Atoi(line[1:])
		util.ErrCheck(err)

		performPartOneMove(&partOneShip, instruction, magnitude)
		performPartTwoMove(&partTwoShip, &partTwoWaypoint, instruction, magnitude)
	}

	part1_res = fmt.Sprintf("%d", intAbs(partOneShip.coord.x)+intAbs(partOneShip.coord.y))
	part2_res = fmt.Sprintf("%d", intAbs(partTwoShip.coord.x)+intAbs(partTwoShip.coord.y))

	return part1_res, part2_res
}
