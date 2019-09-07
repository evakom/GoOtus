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

	for _, r := range input {
		if unicode.IsDigit(r) {
			digits += string(r)
			continue
		}

		digit, err := strconv.Atoi(digits)
		if err != nil {
			result.WriteRune(r)
			continue
		}

		result.WriteString(strings.Repeat(string(r), digit))

		digits = ""
	}

	return result.String()
}
