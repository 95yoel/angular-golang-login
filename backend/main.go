package main

import (
	databaseconfig "backend/src/database_config"
	"backend/src/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Backend Working !!!")

	//Call the HTTP router
	r := routes.Router()

	//Apply the CORS middleware to the router
	handler := databaseconfig.CorsMiddleware(r)

	// Start the server on port 8080
	http.ListenAndServe(":8080", handler)
}
