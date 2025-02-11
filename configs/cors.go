package configs

import "net/http"

func EnableCors(w *http.ResponseWriter) {
	// for local development -> http://localhost:5173
	// for prod -> https://cafeaulait.ch
	(*w).Header().Set("Access-Control-Allow-Origin", "https://cafeaulait.ch")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
