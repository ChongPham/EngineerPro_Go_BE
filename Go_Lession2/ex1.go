package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	BirthYear int
	Job       string
}

func (p Person) CalcuAge() int {
	return time.Now().Year() - p.BirthYear
}

func (p Person) IsJobFit() bool {
	return p.BirthYear%len(p.Name) == 0
}

func Exercise1() {
	person := Person{"Trong", 1999, "Software Enginner"}

	fmt.Println("Person's Information: ")
	fmt.Println("Name: ", person.Name)
	fmt.Println("Birth Year: ", person.BirthYear)
	fmt.Println("Job: ", person.Job)

	fmt.Println("Age: ", person.CalcuAge())
	fmt.Println("Is fit for job: ", person.IsJobFit())
}
