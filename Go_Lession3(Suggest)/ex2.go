package main

import (
	"fmt"
	"time"
)

type Employees struct {
	Id           int
	Name         string
	Salary       int
	Age          int
	ProfileImage string
}

func worker(id int, jobs <-chan Employees, results chan<- float64) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", "employee", j.Name)
		time.Sleep(time.Second)
		result := float64(j.Salary) / float64(j.Age)
		fmt.Println("worker", id, "finished job", "employee", j.Name)
		results <- result
	}

}
func Exercise2() {
	numJob := 5
	jobs := make(chan Employees)
	results := make(chan float64, numJob)

	employees := []Employees{
		{1, "Tiger Nixon", 320800, 61, ""},
		{2, "Garrett Winters", 170750, 63, ""},
		{3, "Ashton Cox", 86000, 66, ""},
		{4, "Cedric Kelly", 433060, 22, ""},
		{5, "Airi Satou", 162700, 33, ""},
	}

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for _, employee := range employees {
		jobs <- employee
	}
	close(jobs)

	for i := 1; i <= numJob; i++ {
		result := <-results
		fmt.Printf("Result: %.2f\n", result)
	}
}
