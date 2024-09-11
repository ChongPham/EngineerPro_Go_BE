package handlers

import (
	"encoding/json"
	"io"
	"myapp/database"
	"myapp/models"
	"myapp/utils"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(body, &user)

	// check if username already exits
	if database.CheckUserExists(user.Username) {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// encode password and store user data
	hashedPassword := utils.HashPassword(user.Password)
	user.password = hashedPassword
	database.SaveUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var credentials map[string]string
	json.Unmarshal(body, &credentials)

	username := credentials["username"]
	password := credentials["password"]

	// get data from database
	user, err := database.GetUserByUsername(username)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Tạo JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
