package day17

import (
	"fmt"
	"sync"
)

type dimensionKey struct {
	x int
	y int
	z int
	w int
}

func countAdjacent(current_location dimensionKey, dimension map[dimensionKey]bool, checkFourthDimension bool) int {
	total := 0
	// for each of the 24 possible adjacent locations, see if it exists in the given dimension
	for x := current_location.x - 1; x <= current_location.x+1; x++ {
		for y := current_location.y - 1; y <= current_location.y+1; y++ {
			for z := current_location.z - 1; z <= current_location.z+1; z++ {
				if checkFourthDimension {
					for w := current_location.w - 1; w <= current_location.w+1; w++ {
						exists, is_active := dimension[dimensionKey{x, y, z, w}]
						if exists && is_active {
							total += 1
						}
					}
				} else {
					exists, is_active := dimension[dimensionKey{x, y, z, 0}]
					if exists && is_active {
						total += 1
					}
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

func buildBlankDimension(hasFourDimensions bool) map[dimensionKey]bool {
	dimension := map[dimensionKey]bool{}
	i := 15
	for x := -i; x <= i; x++ {
		for y := -i; y <= i; y++ {
			for z := -i; z <= i; z++ {
				if hasFourDimensions {
					for w := -i; w <= i; w++ {
						dimension[dimensionKey{x, y, z, w}] = false
					}
				} else {
					dimension[dimensionKey{x, y, z, 0}] = false
				}
			}
		}
	}

	return dimension
}

func calculateAndSetLocationOnDimension(
	location dimensionKey,
	current_dimension map[dimensionKey]bool,
	target_dimension map[dimensionKey]bool,
	checkFourthDimension bool,
	waitGroup *sync.WaitGroup,
	mutex *sync.Mutex) {

	adjacent_count := countAdjacent(location, current_dimension, checkFourthDimension)
	is_active := current_dimension[location]

	mutex.Lock()
	if is_active && !(adjacent_count == 2 || adjacent_count == 3) {
		// is active, make it inactive
		target_dimension[location] = false
	} else if !is_active && adjacent_count == 3 {
		// is inactive, make it active
		target_dimension[location] = true
	} else {
		//all other options result in current location being unchanged
		target_dimension[location] = current_dimension[location]
	}
	mutex.Unlock()

	waitGroup.Done()
}

func simulate(lines []string, useFourthDimension bool) int {
	dimension := buildBlankDimension(useFourthDimension)

	// read initial state
	for iy, line := range lines {
		for ix, char := range line {
			if char == '#' {
				dimension[dimensionKey{ix, iy, 0, 0}] = true
			}
		}
	}

	dimension_iteration_count := 6
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	for range dimension_iteration_count {
		// create new blank dimension
		next_dimension := buildBlankDimension(useFourthDimension)

		for location := range dimension {
			waitGroup.Add(1)
			go calculateAndSetLocationOnDimension(location, dimension, next_dimension, useFourthDimension, &waitGroup, &mutex)
		}

		waitGroup.Wait()

		total_active_in_iteration := 0
		for _, is_active := range next_dimension {
			if is_active {
				total_active_in_iteration++
			}
		}

		dimension = next_dimension
	}

	total_active_in_iteration := 0
	for _, is_active := range dimension {
		if is_active {
			total_active_in_iteration++
		}
	}

	return total_active_in_iteration
}

func Run(lines []string) (part1_res string, part2_res string) {
	part1_res = fmt.Sprintf("%d", simulate(lines, false))
	part2_res = fmt.Sprintf("%d", simulate(lines, true))
	return part1_res, part2_res
}
