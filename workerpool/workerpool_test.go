/*
 * HomeWork-5: Worker Pool
 * Created on 29.09.19 11:22
 * Copyright (c) 2019 - Eugene Klimov
 */

package workerpool

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Constants.
// const (
// 	JOBSNUM   = 10 // number of all jobs
// 	MAXJOBS   = 3  // max concurrency jobs/workers
// 	MAXERRORS = 4  // max errors from all jobs
// )

var testCases = []struct {
	jobs        []Job
	jobSum      int
	maxJobs     int
	maxErrors   int
	errExpected error
	description string
}{
	{
		jobSum:      10,
		maxJobs:     3,
		maxErrors:   1,
		errExpected: ErrWorkerAborted,
		description: "10 jobs, 3 workers, max 4 errors, no errors",
	},
}

func TestWorkerPool(t *testing.T) {
	genJobs()
	// println(len(testCases[0].jobs))
	// println(len(test.jobs))
	for _, test := range testCases {
		err := WorkerPool(test.jobs, test.maxJobs, test.maxErrors)
		if err != test.errExpected {
			if err != nil {
				t.Errorf("FAIL '%s':\n\t WorkerPool returned error '%s', expected nil error.", test.description, err)
			} else {
				t.Errorf("FAIL '%s':\n\t WorkerPool returned nil error, expected error '%s'.", test.description, ErrWorkerAborted)
			}
			continue
		}
		t.Logf("PASS WorkerPool - %s", test.description)
	}

	// start
	// if err := WorkerPool(jobs, MAXJOBS, MAXERRORS); err != nil {
	// 	t.Logf("error %v", err)
	// }
	// fmt.Println("All jobs returned successfully!")
}

func genJobs() {
	rand.Seed(time.Now().UnixNano())

	for i, test := range testCases {
		test.jobs = make([]Job, 0)
		for i := 0; i < test.jobSum; i++ {
			i := i
			job := func() error {
				d := rand.Intn(5) + 1                      // random time for every job
				n := strconv.Itoa(i)                       // job id
				time.Sleep(time.Duration(d) * time.Second) // any work here
				if rand.Intn(2) == 0 {                     // error gen randomly
					return fmt.Errorf("job '%s' returned error", n)
				}
				fmt.Printf("job '%s' ended successfully, duration: %d seconds\n", n, d)
				return nil
			}
			test.jobs = append(test.jobs, job)
			// println(len(test.jobs))
		}
		testCases[i].jobs = test.jobs
	}
	// println(len(testCases[0].jobs))
}

// var decodeTests = []struct {
// 	input       string
// 	expected    string
// 	description string
// }{
// 	{"", "", "empty string"},
// 	{"a4bc2d5e", "aaaabccddddde", "simple coded"},
// 	{"abcd", "abcd", "single characters lower"},
// 	{"45", "", "fail string"},
// 	{"XYZ", "XYZ", "single characters upper"},
// 	{"A2B3C4", "AABBBCCCC", "no single characters upper"},
// 	{"W12BW12B3W24B", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "many characters with repeat"},
// 	{" 2hsq2 qw2 2", "  hsqq qww  ", "whitespace mixed in string"},
// 	{"a2b3c4", "aabbbcccc", "no single characters lower"},
// 	{"a0b2", "bb", "with zero count"},
// 	{"z1y1x1", "zyx", "only one count per char"},
// 	{`\,1\$2\.3\*4`, ",$$...****", "esc punctuation chars"},
// 	{`qwe\4\5`, `qwe45`, "string with 2 esc numbers"},
// 	{`qwe\45`, `qwe44444`, "string with 1 esc char"},
// 	{`qwe\\5`, `qwe\\\\\`, "string with same esc character"},
// 	{`\`, "", "fail esc string"},
// 	{"А1Б2Ц3Я0", "АББЦЦЦ", "cyrillic string"},
// 	{`a4bc2d5eabcdXYZA2B3C4W12BW12B3W24B 2hsq2 qw2 2a2b3c4a0b2z1y1x1\,1\$2\.3\*4qwe\4\5qwe\45qwe\\5А1Б2Ц3`,
// 		`aaaabccdddddeabcdXYZAABBBCCCCWWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB  hsqq qww  aabbbccccbbzyx,$$...****qwe45qwe44444qwe\\\\\АББЦЦЦ`,
// 		"mixed all test strings"},
// }

// func TestUnpackString(t *testing.T) {
// 	for _, test := range decodeTests {
// 		if actual := UnpackString(test.input); actual != test.expected {
// 			t.Errorf("FAIL %s - UnpackString(%s) = '%s', expected '%s'.",
// 				test.description, test.input, actual, test.expected)
// 			continue
// 		}
// 		t.Logf("PASS UnpackString - %s", test.description)
// 	}
// }

// func BenchmarkUnpackString(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		UnpackString(`a4bc2d5eabcdXYZA2B3C4W12BW12B3W24B 2hsq2 qw2 2a2b3c4a0b2a0000b2z1y1x1\,1\$2\.3\*4qwe\4\5qwe\45qwe\\5`)
// 	}
// }
