package handlers

import (
	"cafeaulait-server/configs"
	"fmt"
	"net/http"
)

func ListHello(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(&w, r)

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "mittwuch monolithic server is live!")
	}
}
