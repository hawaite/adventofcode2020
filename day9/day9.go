package day9

import (
	"fmt"
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

func validate(num_buffer []int64, num int64) bool {
	// validate if there are two numbers in the num_buffer slice which sum to "num"
	for i := 0; i < len(num_buffer); i++ {
		for j := i + 1; j < len(num_buffer); j++ {
			if num_buffer[i]+num_buffer[j] == num {
				return true
			}
		}
	}
	return false
}

func validate_consecutive(num_buffer []int64, found_weakness int64) ([]int64, bool) {
	// need to test all consecutive runs in the buffer

	// for all possible window sizes over the buffer
	for window_size := 2; window_size < len(num_buffer)+1; window_size++ {
		// for all possible locations of that window
		for window_start := 0; window_start < len(num_buffer)-window_size+1; window_start++ {
			window_slice := num_buffer[window_start : window_start+window_size]
			if sum_slice(window_slice) == found_weakness {
				return window_slice, true
			}
		}
	}
	return nil, false
}

func sum_slice(num_buffer []int64) int64 {
	total := int64(0)
	for i := 0; i < len(num_buffer); i++ {
		total += num_buffer[i]
	}
	return total
}

func slice_min_max(num_buffer []int64) (min int64, max int64) {
	min = num_buffer[0]
	max = num_buffer[0]
	for i := 0; i < len(num_buffer); i++ {
		if num_buffer[i] < min {
			min = num_buffer[i]
		}
		if num_buffer[i] > max {
			max = num_buffer[i]
		}
	}

	return min, max
}

func Run(lines []string) (string, string) {
	var part1_res int64
	var part2_res int64

	buffer_size := 25 // set buffer size to 5 when running example file
	num_buffer := []int64{}
	found_weakness := int64(0)

	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		util.ErrCheck(err)
		if len(num_buffer) < buffer_size {
			// still in preamble
			num_buffer = append(num_buffer, num)
		} else {
			// no longer in pre-amble
			// check if we find a sum in the num_buffer.
			result := validate(num_buffer, num)
			if !result {
				part1_res = num
				fmt.Println("Part 1: Nothing in", num_buffer, "add up to", num)
				found_weakness = num
				break
			}

			// strip the first number off and add this one on to the end
			num_buffer = append(num_buffer[1:], num)
		}
	}

	num_buffer = []int64{} // reset the num buffer
	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		util.ErrCheck(err)
		if len(num_buffer) < buffer_size {
			// still in preamble
			num_buffer = append(num_buffer, num)
		} else {
			// no longer in pre-amble
			// check if we find a sum in the num_buffer.
			valid_run, valid := validate_consecutive(num_buffer, found_weakness)
			if valid {
				fmt.Println("Part 2: Found consecutive run", valid_run, "which adds to the weakness of", num)
				min, max := slice_min_max(valid_run)
				part2_res = min + max
				fmt.Println("Min:", min, "Max:", max, "Result:", part2_res)
				break
			}

			// strip the first number off and add this one on to the end
			num_buffer = append(num_buffer[1:], num)
		}
	}

	return fmt.Sprintf("%d", part1_res), fmt.Sprintf("%d", part2_res)
}
