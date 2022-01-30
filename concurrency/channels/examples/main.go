package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalculation(i Job) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete int %v\n", i, duration)
	return int(i) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	fmt.Println("The jobs", jobs)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}
	return jobs
}

func runJob(resultChan chan int, i Job) {
	resultChan <- longCalculation(i)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := makeJobs()

	resultChan := make(chan int, 10)

	for i := 0; i < len(jobs); i++ {
		go runJob(resultChan, jobs[i])
	}

	resultCount := 0
	sum := 0
	for {
		result := <-resultChan
		sum += result
		resultCount++
		if resultCount == len(jobs) {
			break
		}
	}
	fmt.Println("sum is", sum)

}
