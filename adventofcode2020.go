package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hawaite/aoc2020/day1"
	"github.com/hawaite/aoc2020/util"
)

func main() {
	day_flag := flag.Int("day", 0, "the day to execute [REQUIRED]")
	input_option_flag := flag.Bool("example", false, "indicates should the example input be used or the full input")
	flag.Parse()

	if *day_flag == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Running Day", *day_flag)
	fmt.Println("=====================")
	var part1_result string
	var part2_result string
	switch *day_flag {
	case 1:
		part1_result, part2_result = day1.Run(util.GetLinesForDay(1, *input_option_flag))
	default:
		panic("Not implemented")
	}
	fmt.Println("=====================")
	fmt.Printf("Part 1 result: %s\nPart 2 result: %s\n", part1_result, part2_result)
}
