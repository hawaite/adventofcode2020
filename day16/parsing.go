package day16

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/hawaite/aoc2020/util"
)

const ruleRegexStr = `^(?P<rule_name>[a-z ]+): ` +
	`(?P<min_one>[0-9]+)-(?P<max_one>[0-9]+)` +
	` or ` +
	`(?P<min_two>[0-9]+)-(?P<max_two>[0-9]+)$`

// validation rules are always two ranges which are OR'd
type validationRule struct {
	ruleName   string
	minimumOne int
	maximumOne int
	minimumTwo int
	maximumTwo int
}

type ticket []int

func parseRule(line string) (parsedValidationRule validationRule) {
	ruleRegex, err := regexp.Compile(ruleRegexStr)
	util.ErrCheck(err)

	matches := ruleRegex.FindStringSubmatch(line)
	rule_name_str := matches[ruleRegex.SubexpIndex("rule_name")]
	parsedValidationRule.ruleName = rule_name_str

	min_one_str := matches[ruleRegex.SubexpIndex("min_one")]
	min_one, err := strconv.Atoi(min_one_str)
	util.ErrCheck(err)
	parsedValidationRule.minimumOne = min_one

	max_one_str := matches[ruleRegex.SubexpIndex("max_one")]
	max_one, err := strconv.Atoi(max_one_str)
	util.ErrCheck(err)
	parsedValidationRule.maximumOne = max_one

	min_two_str := matches[ruleRegex.SubexpIndex("min_two")]
	min_two, err := strconv.Atoi(min_two_str)
	util.ErrCheck(err)
	parsedValidationRule.minimumTwo = min_two

	max_two_str := matches[ruleRegex.SubexpIndex("max_two")]
	max_two, err := strconv.Atoi(max_two_str)
	util.ErrCheck(err)
	parsedValidationRule.maximumTwo = max_two

	return
}

func parseTicket(line string) (parsedTicket ticket) {
	parts := strings.Split(line, ",")
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		util.ErrCheck(err)
		parsedTicket = append(parsedTicket, val)
	}

	return
}

func parseAllInput(lines []string) (validationRules []validationRule, myTicket ticket, otherTickets []ticket) {
	parse_rules := true
	parse_own_ticket := false
	parse_other_tickets := false

	for _, line := range lines {
		if line == "" {
			continue
		}

		if parse_rules {
			if line == "your ticket:" {
				// move to next parsing state
				parse_rules = false
				parse_own_ticket = true
				continue
			} else {
				// parse rule
				validationRules = append(validationRules, parseRule(line))
			}
		} else if parse_own_ticket {
			if line == "nearby tickets:" {
				// move to next parsing state
				parse_own_ticket = false
				parse_other_tickets = true
				continue
			} else {
				// parse my ticket
				myTicket = parseTicket(line)
			}
		} else if parse_other_tickets {
			// parse other tickets to the end of the input
			otherTickets = append(otherTickets, parseTicket(line))
		}
	}

	return
}
