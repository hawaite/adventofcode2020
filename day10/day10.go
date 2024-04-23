package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)

	nums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		check(err)
		nums = append(nums, num)
	}

	slices.Sort(nums)

	gap_count_map := map[int]int{}
	gap_count_map[1] = 0
	gap_count_map[2] = 0
	gap_count_map[3] = 1 // built-in adapter always a gap of 3

	for i := 0; i < len(nums); i++ {
		var num1 int = nums[i]
		var num2 int
		if i == 0 {
			num2 = 0
		} else {
			num2 = nums[i-1]
		}

		gap := num1 - num2
		_, exists := gap_count_map[gap]

		if !exists {
			panic("Had a gap other than 1, 2, or 3")
		}

		gap_count_map[gap] += 1
	}

	fmt.Println("Part 1 answer:", gap_count_map[1]*gap_count_map[3])
}
