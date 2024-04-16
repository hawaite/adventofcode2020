package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PassportValidations struct {
	byr_valid bool
	iyr_valid bool
	eyr_valid bool
	hgt_valid bool
	hcl_valid bool
	ecl_valid bool
	pid_valid bool
	cid_valid bool
}

func range_validation(val_to_test string, range_start int, range_end int) (bool, error) {
	number, err := strconv.Atoi(val_to_test)
	if err != nil {
		return false, err
	}

	if number >= range_start && number <= range_end {
		return true, nil
	}

	return false, nil
}

func height_validation(val_to_test string) (bool, error) {
	unit := val_to_test[len(val_to_test)-2:]
	number := val_to_test[0 : len(val_to_test)-2]

	switch unit {
	case "cm":
		is_valid, err := range_validation(number, 150, 193)
		if err != nil {
			return false, err
		}
		return is_valid, nil
	case "in":
		is_valid, err := range_validation(number, 59, 76)
		if err != nil {
			return false, err
		}
		return is_valid, nil
	default:
		return false, nil
	}
}

func regex_validation(val_to_test string, pattern string) (bool, error) {
	return regexp.Match(pattern, []byte(val_to_test))
}

func validate_passport(passport string, part1 bool) bool {
	var validations PassportValidations

	passport_props := strings.Split(passport, " ")
	for i := 0; i < len(passport_props); i++ {
		kv := strings.Split(passport_props[i], ":")
		key := kv[0]
		val := strings.TrimSuffix(kv[1], "\n")
		switch key {
		case "byr":
			if part1 {
				validations.byr_valid = true
				continue
			}
			is_valid, err := range_validation(val, 1920, 2002)
			check(err)
			validations.byr_valid = is_valid
		case "iyr":
			if part1 {
				validations.iyr_valid = true
				continue
			}
			is_valid, err := range_validation(val, 2010, 2020)
			check(err)
			validations.iyr_valid = is_valid
		case "eyr":
			if part1 {
				validations.eyr_valid = true
				continue
			}
			is_valid, err := range_validation(val, 2020, 2030)
			check(err)
			validations.eyr_valid = is_valid
		case "hgt":
			if part1 {
				validations.hgt_valid = true
				continue
			}
			is_valid, err := height_validation(val)
			check(err)
			validations.hgt_valid = is_valid
		case "hcl":
			if part1 {
				validations.hcl_valid = true
				continue
			}
			is_valid, err := regex_validation(val, "^#[0-9a-f]{6}$")
			check(err)
			validations.hcl_valid = is_valid
		case "ecl":
			if part1 {
				validations.ecl_valid = true
				continue
			}
			valid_ecl_values := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			validations.ecl_valid = slices.Contains(valid_ecl_values, val)
		case "pid":
			if part1 {
				validations.pid_valid = true
				continue
			}
			is_valid, err := regex_validation(val, "^[0-9]{9}$")
			check(err)
			validations.pid_valid = is_valid
		case "cid":
			validations.cid_valid = true
		}
	}

	if validations.byr_valid &&
		validations.iyr_valid &&
		validations.eyr_valid &&
		validations.hgt_valid &&
		validations.hcl_valid &&
		validations.ecl_valid &&
		validations.pid_valid {
		return true
	}

	return false
}

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	valid_passport_count_part1 := 0
	valid_passport_count_part2 := 0
	passport_line_buffer := []string{}

	for scanner.Scan() {
		current_line := strings.TrimSuffix(scanner.Text(), "\n")

		if current_line == "" {
			if validate_passport(strings.Join(passport_line_buffer, " "), true) {
				valid_passport_count_part1++
			}
			if validate_passport(strings.Join(passport_line_buffer, " "), false) {
				valid_passport_count_part2++
			}

			// reset the line buffer.
			passport_line_buffer = []string{}
		} else {
			passport_line_buffer = append(passport_line_buffer, current_line)
		}
	}

	// test if we finished scanning without emptying the buffer
	// happens if the file doesnt end with a blank line
	if len(passport_line_buffer) != 0 {
		if validate_passport(strings.Join(passport_line_buffer, " "), true) {
			valid_passport_count_part1++
		}
		if validate_passport(strings.Join(passport_line_buffer, " "), false) {
			valid_passport_count_part2++
		}
	}

	fmt.Println("Valid passport count for part 1:", valid_passport_count_part1)
	fmt.Println("Valid passport count for part 2:", valid_passport_count_part2)
}
