package config

import (
	"encoding/json"
	"io"
	"log"
	"myprj/MyFirstApp/models"
	"os"
)

var dataFile = "data.json"

// Get data from json file
func GetUsersFromFile() ([]models.User, error) {
	var users []models.User
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return users, nil
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// override data to json
func SaveUsersToFile(users []models.User) error {
	bytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func SaveUser(user models.User) error {
	// Read data from file
	users, err := GetUsersFromFile()
	if err != nil {
		return err
	}
	// add new user and parse data
	users = append(users, user)
	return SaveUsersToFile(users)
}

// get data by username
func GetUserByUsername(username string) (models.User, error) {
	users, err := GetUsersFromFile()
	if err != nil {
		return models.User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, os.ErrNotExist
}

// update data
func UpdateUser(updatedUser models.User) error {
	users, err := GetUsersFromFile()
	if err != nil {
		return err
	}

	for i, user := range users {
		if user.Username == updatedUser.Username {
			users[i] = updatedUser
			return SaveUsersToFile(users)
		}
	}
	return os.ErrNotExist
}

// check user exist
func CheckUserExists(username string) bool {
	users, err := GetUsersFromFile()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
