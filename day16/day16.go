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

func testIfRuleValidForPositionOnAllTickets(tickets []ticket, rule validationRule, ticketField int) bool {
	for _, t := range tickets {
		if !valueMeetsValidation(t[ticketField], rule) {
			return false
		}
	}
	return true
}

func updateFieldPossibleValuesWithoutGivenRule(field_possible_values map[int][]string, field_name string) map[int][]string {
	updated_field_possible_values := map[int][]string{}

	for field_ix, v := range field_possible_values {
		new_field_list := []string{}

		for _, field := range v {
			if field != field_name {
				new_field_list = append(new_field_list, field)
			}
		}

		updated_field_possible_values[field_ix] = new_field_list
	}

	return updated_field_possible_values
}

func Run(lines []string) (part1_res string, part2_res string) {

	validation_rules, my_ticket, other_tickets := parseAllInput(lines)

	valid_tickets := []ticket{}

	total := 0
	for _, ticket := range other_tickets {
		invalid_values := partOneTicketInvalidValues(ticket, validation_rules)
		if len(invalid_values) == 0 {
			valid_tickets = append(valid_tickets, ticket)
		} else {
			for _, val := range invalid_values {
				total += val
			}
		}
	}

	part1_res = fmt.Sprintf("%d", total)

	// holds all the possible rules for the numbered field
	field_possible_values := map[int][]string{}

	// holds the rules we have confirmed for each field
	field_confirmed_values := map[int]string{}

	// for each rule, work out if it is valid for every value in a given field
	// if it is, add it to the list and we'll work out if its the correct rule for that field later
	for _, rule := range validation_rules {
		for i := 0; i < len(valid_tickets[0]); i++ {
			if testIfRuleValidForPositionOnAllTickets(append(valid_tickets, my_ticket), rule, i) {
				field_possible_values[i] = append(field_possible_values[i], rule.ruleName)
			}
		}
	}

	for len(field_confirmed_values) != len(validation_rules) {
		// keep going while we dont have confirmed all fields
		// number of fields is equal to number of rules

		for field_index, possible_field_list := range field_possible_values {
			// if there is only one valid rule for a field, that must be the rule for that field
			if len(possible_field_list) == 1 {
				field_confirmed_values[field_index] = possible_field_list[0]
				field_possible_values = updateFieldPossibleValuesWithoutGivenRule(field_possible_values, possible_field_list[0])
				break // immediately go round again
			}
		}
	}

	// multiplication, so total base value is 1 instead of 0
	total = 1

	for field_ix, field := range field_confirmed_values {
		if len(field) >= 9 && field[0:9] == "departure" {
			total *= my_ticket[field_ix]
		}
	}
	part2_res = fmt.Sprintf("%d", total)

	return part1_res, part2_res
}
