package day18

import "fmt"

type token_type_name string

const (
	number            token_type_name = "NUMBER"
	addition_op       token_type_name = "ADDITION"
	multiplication_op token_type_name = "MULTIPLICATION"
	opening_paren     token_type_name = "OPENING_PAREN"
	closing_paren     token_type_name = "CLOSING_PAREN"
)

type token struct {
	token_type token_type_name
	val        int
}

func do_simple_math(token_stream []token) int {
	total := token_stream[0].val
	current_op := token_stream[1].token_type

	for ix, current_token := range token_stream[2:] {
		if ix%2 == 0 {
			// digit
			if current_op == addition_op {
				total += current_token.val
			} else if current_op == multiplication_op {
				total *= current_token.val
			} else {
				panic(fmt.Sprintf("Expected a number but got %v", current_token))
			}
		} else {
			// oper
			current_op = current_token.token_type
		}
	}

	return total
}

// performs one parentheses simplification and returns new token stream
func perform_single_simplification(token_stream []token) []token {
	opening_paren_ix := 0
	closing_paren_ix := 0
	found_simple_expression := false
	for i := 0; i < len(token_stream); i++ {
		if token_stream[i].token_type == opening_paren {
			for j := i + 1; j < len(token_stream); j++ {
				if token_stream[j].token_type == opening_paren {
					break // wasnt a simple parentheses expression
				} else if token_stream[j].token_type == closing_paren {
					// was a simple expression.
					opening_paren_ix = i
					closing_paren_ix = j
					found_simple_expression = true
					break
				}
			}
			if found_simple_expression {
				break // break from outer loop if we already found a simple expression
			}
		}
	}

	// since we have now identified a simple expression with no nested parentheses
	// evaluate the contents of the expression, then replace the expression with a number token
	math_result := do_simple_math(token_stream[opening_paren_ix+1 : closing_paren_ix])
	new_token_stream := []token{}
	new_token_stream = append(new_token_stream, token_stream[0:opening_paren_ix]...)
	new_token_stream = append(new_token_stream, token{token_type: number, val: math_result})
	new_token_stream = append(new_token_stream, token_stream[closing_paren_ix+1:]...)
	return new_token_stream
}

func token_stream_contains_parentheses(token_stream []token) bool {
	for _, t := range token_stream {
		if t.token_type == opening_paren {
			return true
		}
	}
	return false
}

func Run(lines []string) (part1_res string, part2_res string) {
	// write a maths expression parser that constantly expands paranthesis until there is no more
	// lets build a token stream

	part1_total := 0
	for _, line := range lines {
		token_stream := []token{}
		for _, char := range line {
			switch char {
			case '(':
				token_stream = append(token_stream, token{token_type: opening_paren})
			case ')':
				token_stream = append(token_stream, token{token_type: closing_paren})
			case '+':
				token_stream = append(token_stream, token{token_type: addition_op})
			case '*':
				token_stream = append(token_stream, token{token_type: multiplication_op})
			case ' ':
				continue
			default:
				token_stream = append(token_stream, token{token_type: number, val: int(char - '0')})
			}
		}

		for token_stream_contains_parentheses(token_stream) {
			token_stream = perform_single_simplification(token_stream)
		}

		// should be left with a simple expression
		expression_result := do_simple_math(token_stream)
		part1_total += expression_result
	}

	part1_res = fmt.Sprintf("%d", part1_total)
	return part1_res, part2_res
}
