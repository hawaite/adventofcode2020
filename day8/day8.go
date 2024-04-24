package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

type op struct {
	operator string
	operand  int
}

func parse_op(op_line string) op {
	parts := strings.Split(op_line, " ")
	parsed_operand, err := strconv.Atoi(parts[1][1:])
	util.ErrCheck(err)

	if string(parts[1][0]) != "+" {
		parsed_operand *= -1
	}

	return op{operator: parts[0], operand: parsed_operand}
}

func exec_program(program []op) (halted bool, looped bool, accumulator int) {
	program_counter := 0
	executed_pc_set := map[int]bool{}
	acc := 0
	resulted_in_loop := false
	resulted_in_halt := false

	for {
		if program_counter == len(program) {
			resulted_in_halt = true
			break
		}

		current_op := program[program_counter]

		// test if we have hit this instruction before
		_, have_executed_this_pc_before := executed_pc_set[program_counter]
		if have_executed_this_pc_before {
			resulted_in_loop = true
			break
		}

		executed_pc_set[program_counter] = true
		switch current_op.operator {
		case "nop":
			program_counter++
		case "acc":
			acc += current_op.operand
			program_counter++
		case "jmp":
			program_counter += current_op.operand
		}
	}

	return resulted_in_halt, resulted_in_loop, acc
}

func Run(lines []string) (string, string) {
	var part1_res int
	var part2_res int

	program := []op{}
	for _, line := range lines {
		program = append(program, parse_op(line))
	}

	_, looped, accumulator := exec_program(program)
	if looped {
		part1_res = accumulator
		fmt.Println("Part 1 Final Accumulator:", accumulator)
	} else {
		panic("Program did not infinite loop when infinite loop was expected")
	}

	// Part 2 bruteforce
	// make copy of program
	// modify one nop or jmp
	for i := 0; i < len(program); i++ {
		// only execute modified program if it was a nop or jmp
		if program[i].operator != "acc" {
			modified_program := append([]op{}, program...)
			if modified_program[i].operator == "jmp" {
				modified_program[i].operator = "nop"
			} else {
				modified_program[i].operator = "jmp"
			}

			halted, _, accumulator := exec_program(modified_program)

			if halted {
				part2_res = accumulator
				fmt.Println("Part 2 Final Accumulator:", accumulator)
				break
			}
		}
	}

	return util.IntPairToStringPair(part1_res, part2_res)
}
