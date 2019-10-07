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
}

func TestCopyFileSeekLimit(t *testing.T) {
	for _, test := range testCases {
		from, _ := os.Open(test.fromFile)
		stat, _ := from.Stat()
		expectedBytes := stat.Size()
		actualBytes, err := CopyFileSeekLimit(test.writer, test.toFile, test.fromFile, test.offset, test.limit)
		if err != nil && !test.expError {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns error = '%s', expected = 'no error'.",
				test.description, err)
			continue
		}
		if expectedBytes != actualBytes {
			t.Errorf("FAIL %s - CopyFileSeekLimit() returns bytes = '%d', expected bytes = '%d'.",
				test.description, actualBytes, expectedBytes)
			continue
		}
		t.Logf("PASS CopyFileSeekLimit - %s", test.description)
	}
}
