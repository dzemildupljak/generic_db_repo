package config

import "net/http"

func CorsCofing(handler http.Handler) http.Handler {
	// allowedorigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		// allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// allow specific HTTP headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// handle preflight requests
		if r.Method == "OPTIONS" {
			return
		}
	})
}
