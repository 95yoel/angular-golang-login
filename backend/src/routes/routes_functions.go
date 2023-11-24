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
	// Decode the post objet
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Call the function to encrypt the password
	hashed_password := auth.HashPassword(newUser.Password_hash)

	// Insert new user using the data from the post object
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

/*
LOGIN USERS FROM THE FRONTEND
*/
func Login(w http.ResponseWriter, r *http.Request) {

	var loginUser models.UserLogin
	var userRegistered models.User

	// Decode the post object
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Search the user in the database
	row := dbs.DB.QueryRow(`SELECT * FROM users WHERE email = $1`, loginUser.Email)
	/*
		Scan converts columns read from the database into Go types
	*/
	err = row.Scan(&userRegistered.Id, &userRegistered.Name, &userRegistered.Surname, &userRegistered.Email, &userRegistered.Password_hash)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Cast the stored password into bytes
	storedPasswordBytes := []byte(userRegistered.Password_hash)

	/*
		Check if password is valid, obtain the JWT token and add the claims Name , Surname and email
	*/
	token, err := auth.CheckPassword(storedPasswordBytes, loginUser.Password_hash, userRegistered.Name, userRegistered.Surname, userRegistered.Email)
	if err != nil {
		http.Error(w, "Wrong Password", http.StatusUnauthorized)
		return
	}

	// Write the token in the response
	w.Write([]byte(token))
}

/*
DELETE ALL FIELDS ON THE DATABASE
*/

func DeleteALL(w http.ResponseWriter, r *http.Request) {
	rows, err := dbs.DB.Query("DELETE FROM users")
	if err != nil {
		log.Printf("Error executing the query: %v", err)
		http.Error(w, "Error executing the query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Println("All users deleted")
}
