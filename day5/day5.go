package day5

import (
	"errors"
	"fmt"

	"github.com/hawaite/aoc2020/util"
)

// take a string and interprets it as binary, using whatever character is passed as representing '1'
func binary_string_to_int(row string, char_representing_one string) int {
	factor := 1
	total := 0
	str_len := len(row)
	for i := 0; i < str_len; i++ {
		letter := row[str_len-i-1 : str_len-i]
		if letter == char_representing_one {
			total += factor
		}
		factor *= 2
	}
	return total
}

func get_missing_col(found_cols []string) (string, error) {
	all_cols := []string{"LLL", "LLR", "LRL", "LRR", "RLL", "RLR", "RRL", "RRR"}

	for _, col := range all_cols {
		found := false
		for _, found_col := range found_cols {
			if found_col == col {
				found = true
			}
		}
		if !found {
			return col, nil
		}
	}

	return "", errors.New("lists were the same")
}

func ticket_to_seatid(ticket string) int {
	row_segment := ticket[0:7]
	col_segment := ticket[7:10]
	row_num := binary_string_to_int(row_segment, "B")
	col_num := binary_string_to_int(col_segment, "R")
	return (row_num * 8) + col_num
}

func Run(lines []string) (string, string) {
	var part1_res int
	var part2_res int

	max_seen_ticket := 0
	row_map := map[string][]string{}
	seatid_list := []int{}

	for _, line := range lines {
		ticket := line
		row_segment := ticket[0:7]
		col_segment := ticket[7:10]
		seat_id := ticket_to_seatid(ticket)
		seatid_list = append(seatid_list, seat_id)

		if seat_id > max_seen_ticket {
			max_seen_ticket = seat_id
		}

		_, exists := row_map[row_segment]
		if !exists {
			row_map[row_segment] = []string{col_segment}
		} else {
			row_map[row_segment] = append(row_map[row_segment], col_segment)
		}
	}

	part1_res = max_seen_ticket
	fmt.Println("(Part 1 answer) Max SID:", max_seen_ticket)

	found_empty_seats := []string{}
	for key, value := range row_map {
		// find rows containing 7 cols
		// work out what col is missing
		if len(value) == 7 {
			missing_col, err := get_missing_col(value)
			util.ErrCheck(err)
			found_empty_seats = append(found_empty_seats, key+missing_col)
		}
	}

	for _, seat := range found_empty_seats {
		// validate if the seat means the +/-1 rule
		higher_seat := ticket_to_seatid(seat) + 1
		lower_seat := ticket_to_seatid(seat) - 1
		found_higher := false
		found_lower := false

		for _, seatid := range seatid_list {
			if seatid == higher_seat {
				found_higher = true
			} else if seatid == lower_seat {
				found_lower = true
			}

			if found_lower && found_higher {
				break
			}
		}

		if found_lower && found_higher {
			part2_res = ticket_to_seatid(seat)
			fmt.Println("(part 2 answer) Found valid empty Seat:", seat, "with seatid", ticket_to_seatid(seat))
		}
	}

	return util.IntPairToStringPair(part1_res, part2_res)
}
