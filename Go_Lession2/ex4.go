package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PersonInformation struct {
	Name      string
	Job       string
	BirthYear int
}

func ReadFromFile(filename string) ([]PersonInformation, error) {
	var people []PersonInformation
	// open file
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	// Close file after process
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// line by line
		line := scanner.Text()
		// Split line by "|"
		parts := strings.Split(line, "|")

		name := strings.ToUpper(parts[0])
		job := strings.ToLower(parts[1])

		var birthYear int
		fmt.Sscanf(parts[2], "%d", &birthYear)

		//Create instance and add to slice
		personInFor := PersonInformation{
			Name:      name,
			Job:       job,
			BirthYear: birthYear,
		}

		people = append(people, personInFor)
	}
	//Check error from scanner
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

func Exercise4() {
	people, err := ReadFromFile("textFile.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		// Format specifier used to print the value of the person variable with the field name of the struct
		fmt.Printf("%+v\n", person)
	}
}
