package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hawaite/aoc2020/day1"
	"github.com/hawaite/aoc2020/day10"
	"github.com/hawaite/aoc2020/day11"
	"github.com/hawaite/aoc2020/day12"

	// "github.com/hawaite/aoc2020/day13"
	"github.com/hawaite/aoc2020/day14"
	// "github.com/hawaite/aoc2020/day15"
	// "github.com/hawaite/aoc2020/day16"
	// "github.com/hawaite/aoc2020/day17"
	// "github.com/hawaite/aoc2020/day18"
	// "github.com/hawaite/aoc2020/day19"
	"github.com/hawaite/aoc2020/day2"
	// "github.com/hawaite/aoc2020/day20"
	// "github.com/hawaite/aoc2020/day21"
	// "github.com/hawaite/aoc2020/day22"
	// "github.com/hawaite/aoc2020/day23"
	// "github.com/hawaite/aoc2020/day24"
	// "github.com/hawaite/aoc2020/day25"
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
		12: day12.Run,
		// 13: day13.Run,
		14: day14.Run,
		// 15: day15.Run,
		// 16: day16.Run,
		// 17: day17.Run,
		// 18: day18.Run,
		// 19: day19.Run,
		// 20: day20.Run,
		// 21: day21.Run,
		// 22: day22.Run,
		// 23: day23.Run,
		// 24: day24.Run,
		// 25: day25.Run,
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
