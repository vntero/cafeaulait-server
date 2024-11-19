package handlers

import (
	"cafeaulait-server/configs"
	"cafeaulait-server/models"

	"encoding/json"
	"log"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(&w)

	if r.Method == http.MethodPost {

		// Decode the incoming JSON payload into the struct directly
		var bookData models.BookInput
		err := json.NewDecoder(r.Body).Decode(&bookData)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Log the received data for debugging
		log.Println("Received Book form submission:", bookData)

		// Respond with the received data as confirmation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookData)

		// Send an email with the received data
		SendBookEmail(bookData)
	}
}
