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
		description: "10 jobs, 3 workers, max 1 errors, workers return errors",
	},
	{
		jobSum:      10,
		maxJobs:     3,
		maxErrors:   5,
		errExpected: nil,
		description: "10 jobs, 3 workers, max 5 errors, no workers errors",
	},
}

func TestWorkerPool(t *testing.T) {
	genJobs()
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
		t.Logf("PASS WorkerPool - '%s'", test.description)
	}
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
		}
		testCases[i].jobs = test.jobs
	}
}

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
