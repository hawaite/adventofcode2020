package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)

	part_one_valid_count := 0
	part_two_valid_count := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_parts := strings.Split(line, " ")
		range_part := line_parts[0]
		letter_part := line_parts[1][0]
		password_part := line_parts[2]
		range_parts := strings.Split(range_part, "-")
		range_start, err := strconv.Atoi(range_parts[0])
		check(err)
		range_end, err := strconv.Atoi(range_parts[1])
		check(err)

		letter_count := 0
		for i := 0; i < len(password_part); i++ {
			if password_part[i] == letter_part {
				letter_count++
			}
		}

		if letter_count >= range_start && letter_count <= range_end {
			part_one_valid_count++
		}

		// logical xor
		if (password_part[range_start-1] == letter_part) != (password_part[range_end-1] == letter_part) {
			part_two_valid_count++
		}
	}

	fmt.Println("Total valid passwords for part one: ", part_one_valid_count)
	fmt.Println("Total valid passwords for part two: ", part_two_valid_count)
}
