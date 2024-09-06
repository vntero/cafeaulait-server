package functions

import (
	"cafeaulait-server/models"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method == http.MethodPost {

		// decode incoming JSON
		var registerData models.RegisterInput
		err := json.NewDecoder(r.Body).Decode(&registerData)

		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// log decoded data for debugging
		log.Println("Received Register form submission:", registerData)

		// confirm to client that data has been received
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(registerData)

		SendRegisterEmail(registerData)
	}
}
