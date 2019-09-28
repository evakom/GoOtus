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
	"time"
)

const (
	JOBTIMEOUT = 8  // in seconds
	JOBSNUM    = 10 // number of all jobs
	MAXJOBS    = 5  // max concurrency jobs
	MAXERRORS  = 3  // max errors from all jobs
)

type Job func() error

// WorkerPool is the main worker pool manager.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) error {
	var eg errgroup.Group
	jobsChan := make(chan Job, maxJobs)
	errChan := make(chan error)

	// check errors from jobs
	go func() {
		countErr := 0
		for range errChan {
			countErr++
			if countErr >= MAXERRORS {
				fmt.Printf("\tTotal number of errors - %d, MAX errors: %d, aborting all jobs ...\n", countErr, maxErrors)

			}
		}
	}()

	// start workers
	for i := 0; i < maxJobs; i++ {
		i := i
		eg.Go(func() error {
			jobWorker(i, jobsChan, errChan)
			// what scenario for error?
			return nil
		})
	}

	// send jobs to workers
	for _, j := range jobs {
		jobsChan <- j
	}
	close(jobsChan) // need wait errors chan

	err := eg.Wait()
	time.Sleep(time.Millisecond)
	close(errChan)

	return err
}

func jobWorker(workerNum int, inJob <-chan Job, outErr chan<- error) {
	var err error
	for j := range inJob {
		fmt.Printf("\tWorker: %d started\n", workerNum)
		if err = j(); err != nil {
			fmt.Println(err)
			outErr <- err
		}
		fmt.Printf("\tWorker: %d finished\n", workerNum)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	// jobs slice
	for i := 0; i < JOBSNUM; i++ {
		job := func() error {
			d := rand.Intn(JOBSNUM) + 1                // random time for every job
			time.Sleep(time.Duration(d) * time.Second) // any work here
			if rand.Intn(2) == 0 {                     // error gen randomly
				return fmt.Errorf("job ended with error")
			}
			fmt.Printf("job ended successfully, duration: %d seconds\n", d)
			return nil
		}
		jobs = append(jobs, job)
	}

	// start
	if err := WorkerPool(jobs, MAXJOBS, MAXERRORS); err != nil {
		log.Fatalln("One or more jobs returned with errors!")
	}
	fmt.Println("All jobs returned successfully!")
}
