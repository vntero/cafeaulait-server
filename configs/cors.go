package configs

import "net/http"

func EnableCors(w *http.ResponseWriter) {
	// for local development -> http://localhost:5173
	// for prod -> https://stehliantero.com
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // Allows requests from any origin
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
} 