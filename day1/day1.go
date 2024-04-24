package day1

import (
	"fmt"
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

func Run(lines []string) (string, string) {
	var part1_res string
	var part2_res string

	// Part 1
	for i := 0; i < len(lines); i++ {
		for j := (i + 1); j < len(lines); j++ {
			num_i, err := strconv.Atoi(lines[i])
			util.ErrCheck(err)
			num_j, err := strconv.Atoi(lines[j])
			util.ErrCheck(err)

			if (num_i + num_j) == 2020 {
				fmt.Println("Found values: ", num_i, "and", num_j)
				fmt.Println("Result: ", num_i*num_j)
				part1_res = fmt.Sprintf("%d", num_i*num_j)
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
					fmt.Println("Found values: ", num_i, "and", num_j, "and", num_k)
					fmt.Println("Result: ", num_i*num_j*num_k)
					part2_res = fmt.Sprintf("%d", num_i*num_j*num_k)
				}
			}
		}
	}

	return part1_res, part2_res
}
