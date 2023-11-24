package main

import (
	corsconfig "backend/src/database_config"
	router "backend/src/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Backend Working !!!")

	//Call the HTTP router
	r := router.Router()

	//Apply the CORS middleware to the router
	handler := corsconfig.CorsMiddleware(r)

	// Start the server on port 8080
	http.ListenAndServe(":8080", handler)
}
