package day16

import (
	"fmt"
)

func valueMeetsValidation(val int, rule validationRule) bool {
	return (val >= rule.minimumOne && val <= rule.maximumOne) ||
		(val >= rule.minimumTwo && val <= rule.maximumTwo)
}

func partOneTicketInvalidValues(t ticket, rules []validationRule) (invalidValues []int) {
	for _, val := range t {
		valid_against_one_rule := false
		for _, rule := range rules {
			if valueMeetsValidation(val, rule) {
				valid_against_one_rule = true
				break
			}
		}
		if !valid_against_one_rule {
			invalidValues = append(invalidValues, val)
		}
	}
	return
}

func Run(lines []string) (part1_res string, part2_res string) {

	validation_rules, _, other_tickets := parseAllInput(lines)

	total := 0
	for _, ticket := range other_tickets {
		invalid_values := partOneTicketInvalidValues(ticket, validation_rules)

		for _, val := range invalid_values {
			total += val
		}
	}

	part1_res = fmt.Sprintf("%d", total)
	return part1_res, part2_res
}
