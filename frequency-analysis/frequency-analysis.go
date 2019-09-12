/*
 * HomeWork-3: Frequency Analysis
 * Created on 13.09.19 19:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package frequency_analysis implements counting most popular words.
package frequency_analysis

import (
	"regexp"
	"strings"
)

// NumWords sets the number of returning words.
const NumWords = 10

// Frequency is the base type for counting words.
type Frequency map[string]int

// WordCount returns the frequencies of words in a string for most popular 'num' words.
func WordCount(s string, num int) Frequency {

	result := Frequency{}
	s = strings.ToLower(s)

	var reg = regexp.MustCompile("[a-z0-9а-яё]+('[a-z0-9а-яё])*")
	words := reg.FindAllString(s, -1)

	// count words
	for _, c := range words {
		result[c]++
	}

	// sort map

	// get 'num' most popular words

	return result
}
