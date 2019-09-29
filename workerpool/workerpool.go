/*
 * HomeWork-5: Worker Pool
 * Created on 27.09.19 22:11
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
	MAXJOBS    = 2  // max concurrency jobs/workers
	MAXERRORS  = 2  // max errors from all jobs
)

type Job func() error

// WorkerPool is the main worker pool manager.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) error {
	var eg errgroup.Group

	jobsChan := make(chan Job, maxJobs)
	errChan := make(chan error)
	abortChan := make(chan bool)

	// check errors from jobs
	go func() {
		countErr := 0
		for range errChan {
			countErr++
			if countErr >= maxErrors {
				fmt.Printf("\tTotal number of errors - %d, MAX errors: %d, aborting all jobs ...\n", countErr, maxErrors)
				close(abortChan)
				//close(jobsChan)  // abort all workers
				return
			}
		}
	}()

	// start workers
	for i := 0; i < maxJobs; i++ {
		i := i
		eg.Go(func() error {
			//jobWorker(i, jobsChan, errChan, abortChan)
			// what scenario for error?
			for job := range jobsChan {
				select {
				case <-abortChan:
					fmt.Printf("\tWorker: %d aborted\n", i)
					//return err
					return nil
				default:
					break
				}
				fmt.Printf("\tWorker: %d started\n", i)
				if err := job(); err != nil {
					fmt.Println(err)
					errChan <- err
				}
				fmt.Printf("\tWorker: %d finished\n", i)
			}
			fmt.Printf("\tWorker: %d exited\n", i)
			//return err
			return nil
		})
	}

	// send jobs to workers
	for _, j := range jobs {
		jobsChan <- j
	}
	//if _, ok := <-jobsChan; ok {
	close(jobsChan)
	//}

	err := eg.Wait()
	time.Sleep(time.Millisecond)
	close(errChan)

	return err
}

//func jobWorker(workerNum int, inJob <-chan Job, outErr chan<- error, inAbort <-chan bool) {
//	for job := range inJob {
//		select {
//		case <-inAbort:
//			fmt.Printf("\tWorker: %d aborted\n", workerNum)
//			return
//		default:
//			break
//		}
//		fmt.Printf("\tWorker: %d started\n", workerNum)
//		if err := job(); err != nil {
//			fmt.Println(err)
//			outErr <- err
//		}
//		fmt.Printf("\tWorker: %d finished\n", workerNum)
//	}
//}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	// jobs slice
	for i := 0; i < JOBSNUM; i++ {
		job := func() error {
			d := rand.Intn(5) + 1                      // random time for every job
			time.Sleep(time.Duration(d) * time.Second) // any work here
			if rand.Intn(2) == 0 {                     // error gen randomly
				return fmt.Errorf("error from job")
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
