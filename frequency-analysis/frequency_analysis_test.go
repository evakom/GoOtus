/*
 * HomeWork-3: Frequency Analysis tests
 * Created on 13.09.19 19:20
 * Copyright (c) 2019 - Eugene Klimov
 */

package frequency_analysis

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	input       string
	output      Frequency
}{
	{
		"empty string",
		"",
		Frequency{},
	},
	{
		"count one word",
		"word",
		Frequency{"word": 1},
	},
	{
		"count one of each word",
		"one of each - каждого по одному",
		Frequency{"each": 1, "of": 1, "one": 1, "каждого": 1, "по": 1, "одному": 1},
	},
	{
		"multiple occurrences of a word",
		"one fish two fish red fish blue fish",
		Frequency{"blue": 1, "fish": 4, "one": 1, "red": 1, "two": 1},
	},
	{
		"handles cramped lists",
		"one,two,three",
		Frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"handles expanded lists",
		"one,\ntwo,\nthree",
		Frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"ignore punctuation",
		"car: carpet as java: javascript!!&@$%^&",
		Frequency{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		"include numbers",
		"testing, 1, 2 testing",
		Frequency{"1": 1, "2": 1, "testing": 2},
	},
	{
		"normalize case",
		"go Go GO Stop stop",
		Frequency{"go": 3, "stop": 2},
	},
	{
		"with apostrophes",
		"First: don't laugh. Then: don't cry.",
		Frequency{"cry": 1, "don't": 2, "first": 1, "laugh": 1, "then": 1},
	},
	{
		"with quotations",
		"Joe can't tell between 'large' and large.",
		Frequency{"and": 1, "between": 1, "can't": 1, "joe": 1, "large": 2, "tell": 1},
	},
	{
		"multiple spaces not detected as a word",
		" multiple   whitespaces",
		Frequency{"multiple": 1, "whitespaces": 1},
	},
	{
		"alternating word separators not detected as a word",
		",\n,one,\n ,two \n 'three'",
		Frequency{"one": 1, "three": 1, "two": 1},
	},
}

func TestWordCount(t *testing.T) {
	for _, tt := range testCases {
		expected := tt.output
		actual := WordCount(tt.input, NumWords)
		if len(actual) != NumWords {
			t.Errorf("111")
			continue
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%s\n\tExpected: %v\n\tGot: %v", tt.description, expected, actual)
			continue
		}
		t.Logf("PASS: %s", tt.description)
	}
}

func BenchmarkWordCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			WordCount(tt.input, NumWords)
		}
	}
}
