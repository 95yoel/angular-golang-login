package router

import (
	"log"
	"net/http"
)

func HandleQueryErrors(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("Error executing the query: %v", err)
		// Uncomment the line below to get more info about the error
		// http.Error(w, "Error executing the query", http.StatusInternalServerError)
		return
	}
}

func HandleJSONError(err error, w http.ResponseWriter) {
	if err != nil {
		// Uncomment the line below to get more info about the error
		//http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
}
func HandleDecodingJSONerror(err error, w http.ResponseWriter) {
	if err != nil {
		// Uncomment the line below to get more info about the error
		//http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
}

func HandleErrorInsertingUser(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		// Uncomment the line below to get more info about the error
		//http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}
}

func HandleErrorScanningRow(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		// Uncomment the line below to get more info about the error
		//http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
}

func HandleWrongPasswordError(err error, w http.ResponseWriter) {
	if err != nil {
		// Uncomment the line below to get more info about the error
		//http.Error(w, "Wrong Password", http.StatusUnauthorized)
		return
	}

}
