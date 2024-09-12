package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Profile      string `json:"profile"`
	ProfileImage string `json:"profile_image"`
}
