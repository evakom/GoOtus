/*
 * HomeWork-4: Worker Pool
 * Created on 26.09.19 22:11
 * Copyright (c) 2019 - Eugene Klimov
 */
// Package workerpool implements N-workers with stopping after X-errors.
//package workerpool
package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"runtime"
	"time"
)

const (
	JOBTIMEOUT = 5  // in seconds
	JOBSNUM    = 10 // number of all jobs
	MAXJOBS    = 3  // max concurrency jobs
	MAXERRORS  = 2  // max errors from all jobs
)

type Job func() error

// WorkerPool is the main worker pool manager.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) error {
	var eg errgroup.Group
	jobsChan := make(chan Job, MAXJOBS)

	// start workers
	for i := 0; i < MAXJOBS; i++ {
		i := i
		eg.Go(func() error {
			return jobWorker(i, jobsChan)
		})
	}

	// send jobs to workers
	for _, j := range jobs {
		jobsChan <- j
	}
	close(jobsChan)

	return eg.Wait()
}

func jobWorker(workerNum int, inJob <-chan Job) error {
	for j := range inJob {
		fmt.Printf("Worker: %d started\n", workerNum)
		fmt.Println(j())
		fmt.Printf("Worker: %d finished\n", workerNum)
		runtime.Gosched() // common go
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	// jobs slice
	for i := 0; i < JOBSNUM; i++ {
		job := func() error {
			d := rand.Intn(JOBSNUM)                    // random time for every job
			time.Sleep(time.Duration(d) * time.Second) // any work here
			if rand.Intn(2) == 0 {                     // error gen randomly
				return fmt.Errorf("job ended with error")
			}
			fmt.Printf("job ended successfully, duration: %d\n", d)
			return nil
		}
		jobs = append(jobs, job)
	}

	// start
	if err := WorkerPool(jobs, MAXJOBS, MAXERRORS); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("All jobs returned successfully!")
}
