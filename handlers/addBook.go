package handlers

import (
	"cafeaulait-server/configs"
	"cafeaulait-server/data"

	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(&w)

	// responds to preflight OPTIONS requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// check if it's a POST method
	if r.Method == http.MethodPost {
		// Decode the incoming JSON payload into the struct directly
		var bookData data.BookData
		err := json.NewDecoder(r.Body).Decode(&bookData)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate incoming data
		validate := validator.New()
		if err := validate.Struct(bookData); err != nil {
			log.Printf("Invalid data: %v", bookData)
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}

		// Respond with the received data as confirmation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookData)

		// Send an email with the received data
		SendBookEmail(bookData)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
