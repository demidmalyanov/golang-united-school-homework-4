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

func StringSum(input string) (output string, err error) {

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	if countOperands(input) > 2 || countOperands(input) == 0 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	input = stringToFormat(input)

	numbers := make([]string, 0)
	q := 1

	if strings.Count(input, "+") >= 1 {
		numbers = strings.Split(input, "+")
	}
	if strings.Count(input, "-") > 1 {
		numbers = strings.Split(input, "-")

		q = -1
	}

	if strings.Count(input, "-") == 1 {
		symbol := "-"
		numbers = strings.Split(input, "-")
		numbers[1] = symbol + numbers[1]
	}

	sum := 0
	for _, number := range numbers {
		num := strings.ReplaceAll(number, " ", "")

		item, err := strconv.Atoi(num)
		if err != nil {
			return "", fmt.Errorf("Invalid format or less arguments: %w", err)
		}
		sum += item
	}

	return strconv.Itoa(sum * q), nil

}

func stringToFormat(str string) string {
	strWithoutSpaces := strings.ReplaceAll(str, " ", "")
	if strings.HasPrefix(strWithoutSpaces, "+") {
		strWithoutSpaces = strWithoutSpaces[1:]
	}
	return strWithoutSpaces
}

func countOperands(str string) int {
	var count = 0
	count += strings.Count(str, "+")
	count += strings.Count(str, "-")
	return count
}
