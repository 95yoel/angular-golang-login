package main

import (
	dbs "backend/src/database"
	corsconfig "backend/src/database_config"
	router "backend/src/routes"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // import Postgres driver for Go
)

func main() {
	fmt.Println("Backend Working !!!")

	dbs.InitializeDb()

	//Call the HTTP router
	r := router.Router()

	//Apply the CORS middleware to the router
	handler := corsconfig.CorsMiddleware(r)

	// Start the server on port 8080
	http.ListenAndServe(":8080", handler)
}
