package day6

import (
	"fmt"

	"github.com/hawaite/aoc2020/util"
)

func build_answered_question_count_map(answered_questions []string) map[rune]int {
	answer_map := map[rune]int{}
	for _, answered_question_list := range answered_questions {
		for _, rune := range answered_question_list {
			_, exists := answer_map[rune]
			if !exists {
				answer_map[rune] = 1
			} else {
				answer_map[rune] = answer_map[rune] + 1
			}
		}
	}
	return answer_map
}

func build_answered_questions_set(answered_questions []string) string {
	answer_map := build_answered_question_count_map(answered_questions)

	output_str := ""
	for k, _ := range answer_map {
		output_str = output_str + string(k)
	}

	return output_str
}

func build_answered_questions_intersection(answered_questions []string) string {
	answer_map := build_answered_question_count_map(answered_questions)

	output_str := ""
	for k, v := range answer_map {
		if v == len(answered_questions) {
			// this rune appeared in all rows
			output_str = output_str + string(k)
		}
	}

	return output_str
}

func Run(lines []string) (string, string) {

	line_group_buffer := []string{}
	total_part1 := 0
	total_part2 := 0
	for _, line := range lines {
		if line == "" {
			// processing
			answered_question_set := build_answered_questions_set(line_group_buffer)
			answered_question_intersection := build_answered_questions_intersection(line_group_buffer)
			total_part1 += len(answered_question_set)
			total_part2 += len(answered_question_intersection)
			line_group_buffer = []string{}
		} else {
			line_group_buffer = append(line_group_buffer, line)
		}
	}

	// handle no trailing blank line causing buffer to still have data
	if len(line_group_buffer) != 0 {
		answered_question_set := build_answered_questions_set(line_group_buffer)
		answered_question_intersection := build_answered_questions_intersection(line_group_buffer)
		total_part1 += len(answered_question_set)
		total_part2 += len(answered_question_intersection)
	}

	fmt.Println("(part 1) Total :", total_part1)
	fmt.Println("(part 2) Total :", total_part2)

	return util.IntPairToStringPair(total_part1, total_part2)
}
