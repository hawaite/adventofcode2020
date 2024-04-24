package day2

import (
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

func Run(lines []string) (string, string) {
	part_one_valid_count := 0
	part_two_valid_count := 0

	for _, line := range lines {
		line_parts := strings.Split(line, " ")
		range_part := line_parts[0]
		letter_part := line_parts[1][0]
		password_part := line_parts[2]
		range_parts := strings.Split(range_part, "-")
		range_start, err := strconv.Atoi(range_parts[0])
		util.ErrCheck(err)
		range_end, err := strconv.Atoi(range_parts[1])
		util.ErrCheck(err)

		letter_count := 0
		for i := 0; i < len(password_part); i++ {
			if password_part[i] == letter_part {
				letter_count++
			}
		}

		if letter_count >= range_start && letter_count <= range_end {
			part_one_valid_count++
		}

		// logical xor
		if (password_part[range_start-1] == letter_part) != (password_part[range_end-1] == letter_part) {
			part_two_valid_count++
		}
	}

	return util.IntPairToStringPair(part_one_valid_count, part_two_valid_count)
}
