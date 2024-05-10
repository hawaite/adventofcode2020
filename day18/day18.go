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

func operator_one_has_higher_or_equal_precedence_part_two(operatorOne token, operatorTwo token) bool {
	// only scenario with tokens we have
	return (operatorOne.token_type == addition_op && operatorTwo.token_type == multiplication_op) || (operatorOne.token_type == operatorTwo.token_type)
}

func operator_one_has_higher_or_equal_precedence_part_one(operatorOne token, operatorTwo token) bool {
	// return true except when the first token is an opening paren
	return operatorOne.token_type != opening_paren
}

// takes a list of tokens and organises them in to RPN
func shunting_algo(token_stream []token, precedence_func func(token, token) bool) []token {
	output_queue := []token{}
	operator_stack := []token{}

	for _, current_token := range token_stream {
		if current_token.token_type == number {
			output_queue = append(output_queue, current_token)
		} else if current_token.token_type == addition_op || current_token.token_type == multiplication_op {
			// pop operators from the stack while the one on the top has higher precedence or we hit an opening paren
			for len(operator_stack) > 0 && precedence_func(operator_stack[len(operator_stack)-1], current_token) {
				most_recent_operator := operator_stack[len(operator_stack)-1]
				operator_stack = operator_stack[:len(operator_stack)-1]
				// and push them to output queue
				output_queue = append(output_queue, most_recent_operator)
			}

			// then add current token to the operator stack
			operator_stack = append(operator_stack, current_token)

		} else if current_token.token_type == opening_paren {
			operator_stack = append(operator_stack, current_token)
		} else if current_token.token_type == closing_paren {
			// pop operators until we hit an opening paren
			// discard that paren
			for len(operator_stack) != 0 {
				most_recent_operator := operator_stack[len(operator_stack)-1]
				operator_stack = operator_stack[:len(operator_stack)-1]
				if most_recent_operator.token_type == opening_paren {
					break
				} else {
					output_queue = append(output_queue, most_recent_operator)
				}
			}
		}
	}

	// flush all remaining operators to the output
	for len(operator_stack) > 0 {
		op := operator_stack[len(operator_stack)-1]
		operator_stack = operator_stack[:len(operator_stack)-1]
		output_queue = append(output_queue, op)
	}

	return output_queue
}

func do_rpn_math(token_stream []token) int {
	the_stack := []int{}
	for _, t := range token_stream {
		if t.token_type == number {
			the_stack = append(the_stack, t.val)
		} else if t.token_type == addition_op {
			item_one := the_stack[len(the_stack)-1]
			item_two := the_stack[len(the_stack)-2]
			the_stack = the_stack[:len(the_stack)-2]
			the_stack = append(the_stack, item_one+item_two)
		} else if t.token_type == multiplication_op {
			item_one := the_stack[len(the_stack)-1]
			item_two := the_stack[len(the_stack)-2]
			the_stack = the_stack[:len(the_stack)-2]
			the_stack = append(the_stack, item_one*item_two)
		}
	}

	return the_stack[0]
}

func Run(lines []string) (part1_res string, part2_res string) {
	part1_total := 0
	part2_total := 0
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

		part1_total += do_rpn_math(shunting_algo(token_stream, operator_one_has_higher_or_equal_precedence_part_one))

		part2_total += do_rpn_math(shunting_algo(token_stream, operator_one_has_higher_or_equal_precedence_part_two))
	}

	part1_res = fmt.Sprintf("%d", part1_total)
	part2_res = fmt.Sprintf("%d", part2_total)
	return part1_res, part2_res
}
