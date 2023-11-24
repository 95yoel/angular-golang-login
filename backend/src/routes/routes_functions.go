package router

import (
	dbs "backend/src/database"
	"backend/src/models"
	"encoding/json"
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
