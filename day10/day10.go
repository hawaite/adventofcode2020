package day10

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

func Run(lines []string) (string, string) {
	var part1_res int
	// var part2_res int

	nums := []int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		util.ErrCheck(err)
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

	part1_res = gap_count_map[1] * gap_count_map[3]
	fmt.Println("Part 1 answer:", part1_res)

	return util.IntPairToStringPair(part1_res, 0)
}
