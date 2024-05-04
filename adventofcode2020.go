package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hawaite/aoc2020/day1"
	"github.com/hawaite/aoc2020/day10"
	"github.com/hawaite/aoc2020/day11"
	"github.com/hawaite/aoc2020/day12"

	// "github.com/hawaite/aoc2020/day13"
	"github.com/hawaite/aoc2020/day14"
	"github.com/hawaite/aoc2020/day15"

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

const new_file_template = `package day%d

func Run(lines []string) (part1_res string, part2_res string) {
	return part1_res, part2_res
}
`

func initNewDay(day_num int) {
	fmt.Println("Initializing day", day_num)
	fmt.Println("=====================")

	dir_path := path.Join(".", fmt.Sprintf("day%d", day_num))
	input_path := path.Join(dir_path, "input")

	fmt.Println("Building directory", input_path)
	err := os.MkdirAll(input_path, os.ModeDir|0755)
	util.ErrCheck(err) // will fail if already exists

	fmt.Println(path.Join(input_path, "input.txt"))
	err = os.WriteFile(path.Join(input_path, "input.txt"), []byte{}, 0644)
	util.ErrCheck(err)

	fmt.Println(path.Join(input_path, "example.txt"))
	err = os.WriteFile(path.Join(input_path, "example.txt"), []byte{}, 0644)
	util.ErrCheck(err)

	fmt.Println(path.Join(dir_path, fmt.Sprintf("day%d.go", day_num)))
	err = os.WriteFile(
		path.Join(dir_path, fmt.Sprintf("day%d.go", day_num)),
		[]byte(fmt.Sprintf(new_file_template, day_num)),
		0644)
	util.ErrCheck(err)
}

func execDay(day_num int, useExampleInput bool) {
	fmt.Println("Running Day", day_num)
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
		15: day15.Run,
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

	run_func, exists := day_run_func_map[day_num]
	if exists {
		part1_result, part2_result = run_func(util.GetLinesForDay(day_num, useExampleInput))
	} else {
		panic("Not implemented")
	}

	fmt.Println("=====================")
	fmt.Printf("Part 1 result: %s\nPart 2 result: %s\n", part1_result, part2_result)
}

func main() {
	day_flag := flag.Int("day", 0, "the day to execute [REQUIRED]")
	run_flag := flag.Bool("run", false, "Execute the given day. Cannot be used with -init")
	init_flag := flag.Bool("init", false, "Initialise the given day. Cannot be used with -run")
	use_example_input_flag := flag.Bool("example", false, "indicates should the example input be used or the full input")
	flag.Parse()

	// day alway required
	if *day_flag == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// cannot have both init and run flags set
	if *run_flag && *init_flag {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// cannot have neither init and run flags set
	if !*run_flag && !*init_flag {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *run_flag {
		execDay(*day_flag, *use_example_input_flag)
	} else if *init_flag {
		initNewDay(*day_flag)
	}

}
