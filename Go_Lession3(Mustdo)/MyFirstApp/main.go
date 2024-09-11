package main

import (
	"log"
	"myprj/MyFirstApp/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/profile", controller.Profile)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
