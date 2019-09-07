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

	result := strings.Builder{} // faster then string +=
	digits := ""
	char := ""

	for _, r := range input {

		//if !unicode.IsLetter(r) && digits == "" { // fail string
		//	println("---------fail-----------", input, string(r), result.String())
		//	return ""
		//}

		if unicode.IsLetter(r) && char == "" { // 1st char
			char = string(r)
			result.WriteString(char)
			//println("---------1st char-----------", input, char, result.String())
			continue
		}

		if unicode.IsLetter(r) && char != "" { // char after char
			char = string(r)
			result.WriteString(char)
			//println("---------char char-----------", input, char, result.String())
			continue
		}

		if unicode.IsDigit(r) {
			digits += string(r)
			continue
		}

		digit, err := strconv.Atoi(digits)
		if err != nil {
			//println("---------------------")
			result.WriteRune(r)
			continue
		}

		result.WriteString(strings.Repeat(string(char), digit))

		digits = ""
		char = ""
	}

	return result.String()
}
