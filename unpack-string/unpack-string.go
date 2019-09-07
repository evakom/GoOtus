/*
 * HomeWork-2: Unpack String
 * Created on 07.09.19 12:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package unpackstring implements unpacking string.
package unpackstring

import (
	"strconv"
	"strings"
	"unicode"
)

// UnpackString returns string unpacked.
func UnpackString(input string) string {

	//result := strings.Builder{}
	result := ""
	digits := ""
	char := ""

	for i, r := range input {

		//if unicode.IsLetter(r) || unicode.IsSpace(r) {
		//	char = string(r)
		//	//digits = "1"
		//	//result.WriteString(char)
		//	continue
		//}

		if unicode.IsDigit(r) {
			digits += string(r)
			if i != len(input)-1 { // last char is number
				continue
			}
		}

		digit, _ := strconv.Atoi(digits) // no need to check error - IsDigit checks it
		if digit == 0 && len(digits) > 0 {
			//println(input, char, digits, digit)
			result = result[:len(result)-1]
			char = ""
			//digits = ""
			//continue
		}

		if digit > 0 {
			//println("in digit", input, char, digit, digits)
			//result.WriteString(strings.Repeat(char, digit-1)) // one char already appended
			result += strings.Repeat(char, digit-1) // one char already appended
			digits = ""
			char = ""
		}

		if unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			char = string(r)
			//digits = "1"
			//result.WriteString(char)
			result += char
			//continue
		}
	}

	//return result.String()
	return result
}
