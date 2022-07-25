package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
import (
	"fmt"
	"sync"
	"time"
)

func process(val int, wg *sync.WaitGroup) {

	fmt.Println("started processing", val)

	time.Sleep(2 * time.Second)
	fmt.Println("ended processing", val)

	wg.Done()
}

func main() {
	total := 3

	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines are finished")
}*/

/*func producer(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("wrote ", i, " into the channel")
	}
	close(ch)
}

func main() {

	ch := make(chan int, 10)

	go producer(ch)

	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read val", v, "from the channel")
		time.Sleep(2 * time.Second)
	}
}*/

type Job struct {
	id     int
	number int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {

	var sum int
	for number != 0 {
		digit := number % 10
		sum = sum + digit
		number /= 10
	}

	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.number)}

		results <- output
	}
	wg.Done()
}

func createWorkerPool(workersCount int) {
	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		number := rand.Intn(999)

		job := Job{i, number}

		jobs <- job
	}

	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input %d, sum of digits %d\n", result.job.id, result.job.number, result.sumOfDigits)
	}
	done <- true
}

func main_1() {
	start := time.Now()
	jobsCount := 100

	go allocate(jobsCount)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 50
	createWorkerPool(noOfWorkers)
	<-done

	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("Total time taken", diff)
}
