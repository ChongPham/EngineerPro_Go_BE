package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Employee struct {
	Id           int    `json:"id"`
	Name         string `json:"employee_name"`
	Salary       int    `json:"employee_salary"`
	Age          int    `json:"employee_age"`
	ProfileImage string `json:"profile_image"`
}

type ApiRespone struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}

func getEmployees() ([]Employee, error) {
	// Call Api
	res, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var apiResponse ApiRespone
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}

func worker(id int, jobs <-chan Employee, results chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for employee := range jobs {
		time.Sleep(1 * time.Second)
		if employee.Age == 0 {
			continue
		}
		result := float64(employee.Salary) / float64(employee.Age)
		fmt.Printf("Worker %d: Employee %s, Salary/Age = %.2f\n", id, employee.Name, result)
		results <- result
	}
}

func main() {
	// get employee infor from API
	employees, err := getEmployees()
	if err != nil {
		log.Fatalf("Error get data from API: %v", err)
	}

	numWorkers := 5
	jobs := make(chan Employee, len(employees))
	results := make(chan float64, len(employees))

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for _, employee := range employees {
		jobs <- employee
	}
	close(jobs)

	// wait workers finish
	wg.Wait()
	close(results)
}
