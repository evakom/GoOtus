/*
 * HomeWork-6: FileCopy utility like dd
 * Created on 07.10.2019 17:26
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var testCases = []struct {
	writer           io.Writer
	fromFile, toFile string
	offset, limit    int64
	expError         bool
	description      string
}{
	{
		ioutil.Discard,
		"/var/log/auth.log",
		"test.log",
		0, 0,
		false,
		"simple copy all file with no limit and offset",
	},
	{
		ioutil.Discard,
		"fake.fak",
		"test.log",
		0, 0,
		true,
		"error open source file",
	},
	{
		ioutil.Discard,
		"/var/log/auth.log",
		"",
		0, 0,
		true,
		"error create destination file",
	},
}

func TestCopyFileSeekLimit(t *testing.T) {
	for _, test := range testCases {
		actualBytes, err := CopyFileSeekLimit(test.writer, test.toFile, test.fromFile, test.offset, test.limit)
		if err != nil {
			if !test.expError {
				t.Errorf("FAIL %s - CopyFileSeekLimit() returns error = '%s', expected = 'no error'.",
					test.description, err)
			} else {
				t.Logf("PASS CopyFileSeekLimit - %s", test.description)
			}
			continue
		}
		if test.expError {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns = 'no error', expected error = '%s'.",
				test.description, err)
			continue
		}
		from, err := os.Open(test.fromFile)
		if err != nil {
			t.Fatalf("can't open test file: %s", err)
		}
		stat, err := from.Stat()
		if err != nil {
			t.Fatalf("can't get file stat: %s", err)
		}
		expectedBytes := stat.Size()
		if expectedBytes != actualBytes {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns bytes = '%d', expected bytes = '%d'.",
				test.description, actualBytes, expectedBytes)
			continue
		}
		t.Logf("PASS CopyFileSeekLimit - %s", test.description)
	}
}
