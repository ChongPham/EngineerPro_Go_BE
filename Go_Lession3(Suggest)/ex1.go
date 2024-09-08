package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func Exercise1() {
	// Call Api
	res, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		fmt.Println("Error api call", err)
		return
	}

	defer res.Body.Close()

	// Read data from the respone body
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error body read", readErr)
		return
	}

	// Parse json
	var apiRespone ApiRespone
	jsonErr := json.Unmarshal(body, &apiRespone)
	if jsonErr != nil {
		fmt.Println("Error json parse", jsonErr)
		return
	}

	// apiResponse.Data slice iteration
	for _, employee := range apiRespone.Data {
		fmt.Printf("ID: %d\n", employee.Id)
		fmt.Printf("Name: %s\n", employee.Name)
		fmt.Printf("Salary: %d\n", employee.Salary)
		fmt.Printf("Age: %d\n", employee.Age)
		fmt.Printf("Profile Image: %s\n", employee.ProfileImage)
		fmt.Println("-----------------------")
	}
}
