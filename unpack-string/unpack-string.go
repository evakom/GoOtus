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

	result := strings.Builder{}
	digits := ""
	char := ""

	for _, r := range input {

		//if !unicode.IsLetter(r) && digits == "" { // fail string
		//	println("---------fail-----------", input, string(r), result.String())
		//	return ""
		//}
		digit, _ := strconv.Atoi(digits) // no need to check error - IsDigit below check it
		//println(char, digit, digits)
		if digit > 0 { // unpack series
			//digit, _ := strconv.Atoi(digits) // no need to check error - IsDigit below check it
			//if err != nil { // write one char only if digit unknown - need tests?
			//	println("----------fail digit-----------", digit)
			//	result.WriteString(char)
			//	continue
			//}
			//println(char, digit)
			//if digit != 0 { // one char already appended
			//	digit--
			//}
			result.WriteString(strings.Repeat(char, digit-1))
			digits = ""
			char = ""
		}

		if unicode.IsLetter(r) { //} && char == "" { // 1st char
			char = string(r)
			result.WriteString(char)
			//println("---------1st char-----------", input, char, result.String())
			continue
		}

		//if unicode.IsLetter(r) && char != "" { // char after char
		//	char = string(r)
		//	result.WriteString(char)
		//	//println("---------char char-----------", input, char, result.String())
		//	continue
		//}

		if unicode.IsDigit(r) {
			digits += string(r)
			continue
		}

		//digit, err := strconv.Atoi(digits)
		//if err != nil {
		//	//println("---------------------")
		//	result.WriteRune(r)
		//	continue
		//}

		//result.WriteString(strings.Repeat(string(char), digit))

		//digits = ""
		//char = ""
	}

	return result.String()
}
