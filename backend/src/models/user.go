package models

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	Password_hash string `json:"password"`
}

type UserLogin struct {
	Email         string `json:"email"`
	Password_hash string `json:"password"`
}
