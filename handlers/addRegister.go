package handlers

import (
	"cafeaulait-server/configs"
	"cafeaulait-server/data"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func AddRegister(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(&w)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// check if it's a POST method
	if r.Method == http.MethodPost {
		// Decode the incoming JSON payload into the struct directly
		var registerData data.ResgisterData
		err := json.NewDecoder(r.Body).Decode(&registerData)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate incoming data
		validate := validator.New()
		if err := validate.Struct(registerData); err != nil {
			log.Printf("Invalid data: %v", registerData)
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}

		// confirm to client that data has been received
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(registerData)

		// Send an email with the received data
		SendRegisterEmail(registerData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
