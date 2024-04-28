package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hawaite/aoc2020/day1"
	"github.com/hawaite/aoc2020/day10"
	"github.com/hawaite/aoc2020/day11"
	"github.com/hawaite/aoc2020/day2"
	"github.com/hawaite/aoc2020/day3"
	"github.com/hawaite/aoc2020/day4"
	"github.com/hawaite/aoc2020/day5"
	"github.com/hawaite/aoc2020/day6"
	"github.com/hawaite/aoc2020/day7"
	"github.com/hawaite/aoc2020/day8"
	"github.com/hawaite/aoc2020/day9"
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

	day_run_func_map := map[int]func([]string) (string, string){
		1:  day1.Run,
		2:  day2.Run,
		3:  day3.Run,
		4:  day4.Run,
		5:  day5.Run,
		6:  day6.Run,
		7:  day7.Run,
		8:  day8.Run,
		9:  day9.Run,
		10: day10.Run,
		11: day11.Run,
	}

	var part1_result string
	var part2_result string

	run_func, exists := day_run_func_map[*day_flag]
	if exists {
		part1_result, part2_result = run_func(util.GetLinesForDay(*day_flag, *input_option_flag))
	} else {
		panic("Not implemented")
	}

	fmt.Println("=====================")
	fmt.Printf("Part 1 result: %s\nPart 2 result: %s\n", part1_result, part2_result)
}
