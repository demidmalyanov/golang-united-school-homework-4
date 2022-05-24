package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorWrongFormat    = errors.New("wrong format input string.")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func extractSymbolsFromNums(input string) (symbols []rune, output string, err error) {

	for i, char := range input {

		switch char {
		case '+', '-':
			if input[i] == '+' && input[i+1] == '+' {
				return symbols, output, fmt.Errorf("%w", errorWrongFormat)
			} else if input[i] == '-' && input[i+1] == '-' {
				return symbols, output, fmt.Errorf("%w", errorWrongFormat)
			}
			symbols = append(symbols, char)
			input = strings.Replace(input, string(input[i]), "*", 1)
		default:
			continue

		}

	}

	output = input

	return symbols, output, nil
}
func StringSum(input string) (output string, err error) {

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")

	symbols, str, err := extractSymbolsFromNums(input)

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if strings.HasPrefix(str, "*") {
		str = strings.Replace(str, "*", "", 1)
	}
	numbers := strings.Split(str, "*")

	if len(numbers) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	// then we should reverse our expression ("5-60" = "-60 + 5")
	if len(symbols) < 2 {
		symbol := symbols[0]
		symbols[0] = '+'
		symbols = append(symbols, symbol)
	}

	total := 0

	for i, symbol := range symbols {
		num, err := strconv.Atoi(string(symbol) + numbers[i])
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}
		total += num

	}

	return strconv.Itoa(total), nil
}
