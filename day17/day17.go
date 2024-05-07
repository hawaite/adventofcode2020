package day17

import "fmt"

type dimensionKey struct {
	x int
	y int
	z int
}

func countAdjacent(current_location dimensionKey, dimension map[dimensionKey]bool) int {
	total := 0
	// for each of the 24 possible adjacent locations, see if it exists in the given dimension
	for x := current_location.x - 1; x <= current_location.x+1; x++ {
		for y := current_location.y - 1; y <= current_location.y+1; y++ {
			for z := current_location.z - 1; z <= current_location.z+1; z++ {
				exists, is_active := dimension[dimensionKey{x, y, z}]
				if exists && is_active {
					total += 1
				}
			}
		}
	}

	// remove one if we were centered on an active cube
	if dimension[current_location] {
		total--
	}

	return total
}

func buildBlankDimension() map[dimensionKey]bool {
	dimension := map[dimensionKey]bool{}

	for x := -20; x <= 20; x++ {
		for y := -20; y <= 20; y++ {
			for z := -20; z <= 20; z++ {
				dimension[dimensionKey{x, y, z}] = false
			}
		}
	}

	return dimension
}

func Run(lines []string) (part1_res string, part2_res string) {

	dimension := buildBlankDimension()

	// read initial state
	for iy, line := range lines {
		for ix, char := range line {
			if char == '#' {
				dimension[dimensionKey{ix, iy, 0}] = true
			}
		}
	}

	dimension_iteration_count := 6

	for i := range dimension_iteration_count {
		// create new blank dimension
		next_dimension := buildBlankDimension()

		// for every cube in the dimension, apply rules
		for location, is_active := range dimension {
			adjacent_count := countAdjacent(location, dimension)
			if is_active && !(adjacent_count == 2 || adjacent_count == 3) {
				// is active, make it inactive
				next_dimension[location] = false
			} else if !is_active && adjacent_count == 3 {
				// is inactive, make it active
				next_dimension[location] = true
			} else {
				//all other options result in current location being unchanged
				next_dimension[location] = dimension[location]
			}
		}

		total_active_in_iteration := 0
		for _, is_active := range next_dimension {
			if is_active {
				total_active_in_iteration++
			}
		}

		fmt.Printf("Total active in generation %d: %d\n", i, total_active_in_iteration)
		dimension = next_dimension
	}

	total_active_in_iteration := 0
	for _, is_active := range dimension {
		if is_active {
			total_active_in_iteration++
		}
	}

	part1_res = fmt.Sprintf("%d", total_active_in_iteration)
	// fmt.Println(dimension)
	// 301 too low
	return part1_res, part2_res
}
