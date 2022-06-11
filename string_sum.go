package string_sum

import (
	"errors"
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
	var (
		firstOperand  string = ""
		secondOperand string = ""
		sign          string = ""
		operation     string = ""
	)

	if input == "" {
		return "", errorEmptyInput
	}

	if strings.TrimSpace(input) == "" {
		return "", errorEmptyInput
	}

	arrSigns := []byte(input)

	for i := 0; i < len(arrSigns); i++ {
		sign = string(arrSigns[i])
		if sign == " " {
			continue
		} else {
			if sign == "+" || sign == "-" {
				if len(firstOperand) == 0 {
					firstOperand = sign
				} else {
					if len(secondOperand) == 0 {
						operation = sign
					} else {
						return "", errorNotTwoOperands
					}
				}
			} else {
				if len(operation) == 0 {
					firstOperand += sign
				} else {
					secondOperand += sign
				}
			}
		}
	}

	if len(secondOperand) == 0 {
		return "", errorNotTwoOperands
	}

	fsNum, err := strconv.Atoi(firstOperand)

	if err != nil {
		return "", err
	}

	scNum, err := strconv.Atoi(secondOperand)

	if err != nil {
		return "", err
	}

	var result int = 0

	if operation == "+" {
		result = fsNum + scNum
	} else {
		result = fsNum - scNum
	}

	output = strconv.Itoa(result)

	return output, nil
}
