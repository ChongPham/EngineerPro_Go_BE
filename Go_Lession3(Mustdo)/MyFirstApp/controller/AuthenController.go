package controller

import (
	"encoding/json"
	"myprj/MyFirstApp/config"
	"myprj/MyFirstApp/models"
	"myprj/MyFirstApp/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// check username exist
	exists := config.CheckUserExists(user.Username)
	if exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// password encoding
	user.Password, _ = utils.HashPassword(user.Password)

	// save user to json file
	err = config.SaveUser(user)
	if err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	http.Error(w, "User registered successfully", http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials map[string]string
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := config.GetUserByUsername(credentials["username"])
	if err != nil || !utils.CheckPasswordHash(credentials["password"], user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// create and send token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	} else if r.Method == http.MethodPut {
		var updatedUser models.User
		json.NewDecoder(r.Body).Decode(&updatedUser)

		user.Profile = updatedUser.Profile
		config.UpdateUser(user)
		http.Error(w, "Profile updated successfully", http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}