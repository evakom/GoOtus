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
	"math/rand"
	"time"
)

const jobTimeOut = 5 // in seconds

type Job func() error

// WorkerPool is the job worker.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) {
	for _, job := range jobs {
		fmt.Println(job())
	}
	//jobControl := make(chan struct{})
	//for i, job := range jobs {
	//	go startJob(j, jobControl)
	//}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	for i := 0; i < 11; i++ {
		i := i // for correct parameter in func
		job := func() error {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			if rand.Intn(2) == 0 { // error gen
				return fmt.Errorf("job '%d' ended with error", i)
			}
			fmt.Printf("job '%d' ended successfully", i)
			return nil
		}
		jobs = append(jobs, job)
	}

	WorkerPool(jobs, 3, 2)
}
