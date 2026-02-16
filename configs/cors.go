package configs

import "net/http"

func EnableCors(w *http.ResponseWriter, r *http.Request) {
    allowedOrigins := []string{
        "https://cafeaulait.ch",
        "https://saminthekitchen.ch",
        "http://localhost:5173", // for local dev
    }
    
    origin := r.Header.Get("Origin")
    
    // Check if the origin is in our allowed list
    for _, allowedOrigin := range allowedOrigins {
        if origin == allowedOrigin {
            (*w).Header().Set("Access-Control-Allow-Origin", origin)
            (*w).Header().Set("Access-Control-Allow-Credentials", "true")
            break
        }
    }
    
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
