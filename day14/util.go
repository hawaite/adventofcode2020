package day14

import (
	"strings"
)

func int64To36BitString(value int64) string {
	bit_string := ""

	for i := 0; i < 36; i++ {
		bit_value := value & (1 << i)
		if bit_value != 0 {
			bit_string = "1" + bit_string
		} else {
			bit_string = "0" + bit_string
		}
	}
	return bit_string
}

func populateXsInBitStringWithReplacementValues(bit_string string, replacement_values string) string {
	if strings.Count(bit_string, "X") > len(replacement_values) {
		panic("Not enough replacement values to populate string")
	}

	// fmt.Println("Number of Xs:", strings.Count(bit_string, "X"))
	out_string := ""
	remaining_replacement_values := replacement_values

	for i := len(bit_string) - 1; i >= 0; i-- {
		if bit_string[i] != 'X' {
			out_string = string(bit_string[i]) + out_string
		} else {
			out_string = string(remaining_replacement_values[len(remaining_replacement_values)-1]) + out_string
			remaining_replacement_values = remaining_replacement_values[:len(remaining_replacement_values)-1]
		}
	}

	return out_string
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
