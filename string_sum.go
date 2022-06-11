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

func CreateOperands(arrSigns []byte) (firstOperand, secondOperand, operation string, err error) {
	var (
		sign string = ""
	)

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
						return "", "", "", fmt.Errorf("error! message = %q ", errorNotTwoOperands)
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

	return firstOperand, secondOperand, operation, nil
}

func ConvertIntoNum(first, second string) (fsNum, scNum int, errors error) {
	var (
		strerr *strconv.NumError
		err    error
	)

	if len(second) == 0 {
		return 0, 0, fmt.Errorf("error! message = %q ", errorNotTwoOperands)
	}

	fsNum, err = strconv.Atoi(first)

	if err != nil {
		strerr = &strconv.NumError{Func: "Atoi", Num: first, Err: strconv.ErrSyntax}
		return 0, 0, fmt.Errorf("error! message = %q", strerr)
	}

	scNum, err = strconv.Atoi(second)

	if err != nil {
		strerr = &strconv.NumError{Func: "Atoi", Num: second, Err: strconv.ErrSyntax}
		return 0, 0, fmt.Errorf("error! message = %q", strerr)
	}

	return fsNum, scNum, nil
}

func Execute(fOper, secOper int, sign string) int {
	if sign == "+" {
		return fOper + secOper
	} else {
		return fOper - secOper
	}
}

func StringSum(input string) (output string, err error) {
	var (
		fsNum         int    = 0
		scNum         int    = 0
		firstOperand  string = ""
		secondOperand string = ""
		operation     string = ""
	)

	if input == "" {
		return "", fmt.Errorf("error! message = %q ", errorEmptyInput)
	}

	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("error! message = %q ", errorEmptyInput)
	}

	arrSigns := []byte(input)
	firstOperand, secondOperand, operation, err = CreateOperands(arrSigns)

	if err != nil {
		return "", err
	}

	fsNum, scNum, err = ConvertIntoNum(firstOperand, secondOperand)

	if err != nil {
		return "", err
	}

	var result int = 0

	result = Execute(fsNum, scNum, operation)
	output = strconv.Itoa(result)

	return output, nil
}
