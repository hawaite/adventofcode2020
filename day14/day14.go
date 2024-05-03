package day14

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hawaite/aoc2020/util"
)

const mask_rx = "^mask = ([01X]*)$"
const mem_rx = `^mem\[([0-9]*)\] = ([0-9]*)$`

func applyMask(value string, mask string) string {
	value_as_int, err := strconv.ParseInt(value, 10, 64)
	util.ErrCheck(err)

	value_bit_string := int64To36BitString(value_as_int)
	output_bit_string := ""

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case 'X':
			output_bit_string = output_bit_string + string(value_bit_string[i])
		case '1':
			output_bit_string = output_bit_string + "1"
		case '0':
			output_bit_string = output_bit_string + "0"
		}
	}

	return output_bit_string
}

func Run(lines []string) (string, string) {
	var part1_res string
	var part2_res string

	mask_regex, err := regexp.Compile(mask_rx)
	util.ErrCheck(err)
	mem_regex, err := regexp.Compile(mem_rx)
	util.ErrCheck(err)

	mem := map[string]int64{}
	current_mask := ""
	for _, line := range lines {
		is_mask := mask_regex.Match([]byte(line))
		is_mem_set := mem_regex.Match([]byte(line))

		if is_mask {
			current_mask = mask_regex.FindAllStringSubmatch(line, -1)[0][1]
		} else if is_mem_set {
			memory_location := mem_regex.FindAllStringSubmatch(line, -1)[0][1]
			memory_val := mem_regex.FindAllStringSubmatch(line, -1)[0][2]

			masked_val := applyMask(memory_val, current_mask)
			res_value, err := strconv.ParseInt(masked_val, 2, 64)
			util.ErrCheck(err)

			mem[memory_location] = res_value
		}
	}

	var total int64 = 0

	for key, val := range mem {
		fmt.Println(key, " = ", val)
		total += val
	}
	part1_res = fmt.Sprintf("%d", total)

	return part1_res, part2_res
}
