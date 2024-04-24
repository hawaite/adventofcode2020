package day1

import (
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

func Run(lines []string) (string, string) {
	var part1_res int
	var part2_res int

	// Part 1
	for i := 0; i < len(lines); i++ {
		for j := (i + 1); j < len(lines); j++ {
			num_i, err := strconv.Atoi(lines[i])
			util.ErrCheck(err)
			num_j, err := strconv.Atoi(lines[j])
			util.ErrCheck(err)

			if (num_i + num_j) == 2020 {
				part1_res = num_i * num_j
			}
		}
	}

	// Part 2
	for i := 0; i < len(lines); i++ {
		for j := (i + 1); j < len(lines); j++ {
			for k := (j + 1); k < len(lines); k++ {
				num_i, err := strconv.Atoi(lines[i])
				util.ErrCheck(err)
				num_j, err := strconv.Atoi(lines[j])
				util.ErrCheck(err)
				num_k, err := strconv.Atoi(lines[k])
				util.ErrCheck(err)

				if (num_i + num_j + num_k) == 2020 {
					part2_res = num_i * num_j * num_k
				}
			}
		}
	}

	return util.IntPairToStringPair(part1_res, part2_res)
}
