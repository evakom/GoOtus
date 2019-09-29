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
const (
	JOBSNUM   = 10 // number of all jobs
	MAXJOBS   = 3  // max concurrency jobs/workers
	MAXERRORS = 4  // max errors from all jobs
)

func TestWorkerPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	// jobs slice
	for i := 0; i < JOBSNUM; i++ {
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
		jobs = append(jobs, job)
	}

	// start
	if err := WorkerPool(jobs, MAXJOBS, MAXERRORS); err != nil {
		t.Logf("error %v", err)
	}
	fmt.Println("All jobs returned successfully!")
}
