package router

import "net/http"

/*
Returns a router HTTP for the backend
*/
func Router() http.Handler {
	r := http.NewServeMux()
	/*
		Write below all the routes of the backend
	*/

	r.HandleFunc("/", GetAllusers)     // GET
	r.HandleFunc("/post", CreateUser)  // POST
	r.HandleFunc("/login", Login)      // POST
	r.HandleFunc("/delete", DeleteALL) //DELETE

	return r
}
