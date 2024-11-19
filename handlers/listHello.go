package handlers

import (
	"cafeaulait-server/configs"
	"fmt"
	"net/http"
)

func ListHello(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(&w)

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "cafeaulait server is live!")
	}
}
