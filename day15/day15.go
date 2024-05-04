package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

func getNextNumber(num_list []int, last_turn_map map[int]int) int {
	last_num := num_list[len(num_list)-1]

	turn_last_seen, exists := last_turn_map[last_num]
	last_turn_map[last_num] = len(num_list)
	if exists {
		new_value := len(num_list) - turn_last_seen
		return new_value
	} else {
		return 0
	}
}

func Run(lines []string) (part1_res string, part2_res string) {
	num_list := []int{}
	last_turn_map := map[int]int{}
	last_number_added_to_map := -1

	// add the initial state
	for i, part := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(part)
		util.ErrCheck(err)
		num_list = append(num_list, num)

		last_turn_map[num] = i + 1
		last_number_added_to_map = num
	}

	// strip the last number from the seen map
	delete(last_turn_map, last_number_added_to_map)

	turn_to_fetch_part_one := 2020
	turn_to_fetch_part_two := 30000000

	for i := len(num_list); i < turn_to_fetch_part_two; i++ {
		num_list = append(num_list, getNextNumber(num_list, last_turn_map))
	}

	part1_res = fmt.Sprintf("%d", num_list[turn_to_fetch_part_one-1])
	part2_res = fmt.Sprintf("%d", num_list[turn_to_fetch_part_two-1])

	return part1_res, part2_res
}
