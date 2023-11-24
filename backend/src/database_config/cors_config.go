package corsconfig

import "net/http"

/*
	Middleware for handle CORS requests
*/
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow all methods: POST, GET , PUT , DELETE
		w.Header().Set("Access-Control-Allow-Methods", "*")

		// Set all the headers allowed
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Handler OPTIONS requests
		if r.Method == http.MethodOptions {

			// Respond with a 200 OK status for Options Requests
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
