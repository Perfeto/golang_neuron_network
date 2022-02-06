package main

import "neurons/internal/figures"

type Job struct {
	X float32
	Y float32
}

type Result struct {
	X       float32
	Y       float32
	Predict figures.GroupID
}

func DoWork(jobChan <-chan Job, resultChan chan<- Result, logic func(Job) Result) {
	for job := range jobChan {
		resultChan <- logic(job)
	}
}
