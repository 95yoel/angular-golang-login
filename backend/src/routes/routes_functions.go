package router

import (
	"backend/src/auth"
	dbs "backend/src/database"
	"backend/src/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
GET ALL USERS FROM BACKEND
*/
func GetAllusers(w http.ResponseWriter, r *http.Request) {
	// Query for select users
	rows, err := dbs.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error executing the query: %v", err)
		http.Error(w, "Error executing the query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store user data
	var users []models.User

	// Iterate through the query results
	for rows.Next() {
		var user models.User
		/*
			Scan converts columns read from the database into Go types
		*/
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Password_hash)
		if err != nil {
			log.Printf("Error scaning results: %v", err)
			http.Error(w, "Error scaning results", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Convert the slice of users into JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	// Configure the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	w.Write(jsonData)
}

/*
POST USER FROM FRONTEND
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser models.User
	// Decode the post data
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// encrypt the password
	hashed_password := auth.HashPassword(newUser.Password_hash)

	// Insert new user
	_, err = dbs.DB.Exec("INSERT INTO users (name, surname, email, password_hash) VALUES ($1, $2, $3, $4)",
		newUser.Name, newUser.Surname, newUser.Email, hashed_password)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	// HTTP response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User create succesfully")
}
