package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	buffer_size := 25
	num_buffer := []int64{}
	found_weakness := int64(0)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
		// fmt.Println("Current number:", num)
		if len(num_buffer) < buffer_size {
			// still in preamble
			num_buffer = append(num_buffer, num)
		} else {
			// no longer in pre-amble
			// check if we find a sum in the num_buffer.
			result := validate(num_buffer, num)
			if result == false {
				fmt.Println("Part 1: Nothing in", num_buffer, "add up to", num)
				found_weakness = num
				break
			}

			// strip the first number off and add this one on to the end
			num_buffer = append(num_buffer[1:], num)
		}
	}

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	num_buffer = []int64{} // reset the num buffer
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
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
				fmt.Println("Min:", min, "Max:", max, "Result:", min+max)
				break
			}

			// strip the first number off and add this one on to the end
			num_buffer = append(num_buffer[1:], num)
		}
	}
}
