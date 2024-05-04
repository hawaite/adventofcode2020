package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

const mask_rx = "^mask = ([01X]*)$"
const mem_rx = `^mem\[([0-9]*)\] = ([0-9]*)$`

func applyMask(value string, mask string, part2 bool) string {
	value_as_int, err := strconv.ParseInt(value, 10, 64)
	util.ErrCheck(err)

	value_bit_string := int64To36BitString(value_as_int)
	output_bit_string := ""

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case 'X':
			if !part2 {
				output_bit_string = output_bit_string + string(value_bit_string[i])
			} else {
				output_bit_string = output_bit_string + "X"
			}
		case '1':
			output_bit_string = output_bit_string + "1"
		case '0':
			if !part2 {
				output_bit_string = output_bit_string + "0"
			} else {
				output_bit_string = output_bit_string + string(value_bit_string[i])
			}
		}
	}

	return output_bit_string
}

func getExpandedLocationBitStreamList(memory_location_bit_string string) []string {
	x_count := strings.Count(memory_location_bit_string, "X")
	generated_addresses := []string{}

	for i := 0; i < IntPow(2, x_count); i++ {
		// get a bit-string representing the current "i"
		// then we will use those values as replacement values to replace
		// all the Xs in the location bit-string
		replacement_values := int64To36BitString(int64(i))

		populated := populateXsInBitStringWithReplacementValues(memory_location_bit_string, replacement_values)
		generated_addresses = append(generated_addresses, populated)
	}

	return generated_addresses
}

func Run(lines []string) (string, string) {
	var part1_res string
	var part2_res string

	mask_regex, err := regexp.Compile(mask_rx)
	util.ErrCheck(err)
	mem_regex, err := regexp.Compile(mem_rx)
	util.ErrCheck(err)

	memV1 := map[string]int64{}
	memV2 := map[string]int64{}

	current_mask := ""
	for _, line := range lines {
		is_mask := mask_regex.Match([]byte(line))
		is_mem_set := mem_regex.Match([]byte(line))

		if is_mask {
			current_mask = mask_regex.FindAllStringSubmatch(line, -1)[0][1]
		} else if is_mem_set {
			memory_location := mem_regex.FindAllStringSubmatch(line, -1)[0][1]
			memory_val := mem_regex.FindAllStringSubmatch(line, -1)[0][2]

			// start part 1
			masked_val := applyMask(memory_val, current_mask, false)
			res_value, err := strconv.ParseInt(masked_val, 2, 64)
			util.ErrCheck(err)
			memV1[memory_location] = res_value
			// end part 1

			// start part 2
			masked_location_bit_string := applyMask(memory_location, current_mask, true)
			parsed_val, err := strconv.ParseInt(memory_val, 10, 64)
			util.ErrCheck(err)

			addresses := getExpandedLocationBitStreamList(masked_location_bit_string)

			for _, address := range addresses {
				memV2[address] = parsed_val
			}
			// end part 2
		}
	}

	var total int64 = 0

	for _, val := range memV1 {
		total += val
	}
	part1_res = fmt.Sprintf("%d", total)

	total = 0
	for _, val := range memV2 {
		total += val
	}
	part2_res = fmt.Sprintf("%d", total)

	// Part 1 result: 13556564111697
	// Part 2 result: 4173715962894
	return part1_res, part2_res
}
