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

func StringSum(input string) (output string, err error) {

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	if countOperands(input) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	numbers := make([]string,0)

	if strings.Count(input,"+") == 0 {
		numbers = strings.Split(input, "+")
	} else {
		numbers = strings.Split(input, "-")
	}

	sum := 0
	for _, number := range numbers {
		item, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil{
			return "", fmt.Errorf("syntax err: %w", err)
		}
		sum += item
	}

	return "", nil
}

func countOperands(str string) int {
	var count = 0
	count += strings.Count(str, "+")
	count += strings.Count(str, "-")
	return count
}
