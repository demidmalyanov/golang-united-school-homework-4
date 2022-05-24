package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
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

func extractSymbolsFromNums(input string) (symbols, numbers []rune, err error) {

	for i, char := range input {

		switch char {
		case '+', '-':
			if input[i] == '+' && input[i+1] == '+' {
				return symbols, numbers, fmt.Errorf("%w", errorWrongFormat)
			} else if input[i] == '-' && input[i+1] == '-' {
				return symbols, numbers, fmt.Errorf("%w", errorWrongFormat)
			}
			symbols = append(symbols, char)
		default:
			numbers = append(numbers, char)

		}

	}

	return symbols, numbers, nil
}

func StringSum(input string) (output string, err error) {

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")

	symbols, numbers, err := extractSymbolsFromNums(input)

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if len(numbers)!= 2{
		return "",fmt.Errorf("%w",errorNotTwoOperands)
	}

	total := 0

	for i, symbol := range symbols {
		num, err := strconv.Atoi(string(symbol) + string(numbers[i]))
		if err != nil {
			return "",fmt.Errorf("%w",errorWrongFormat)
		}
		total += num

	}

	return strconv.Itoa(total), nil
}
