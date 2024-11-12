package functions

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Server is live!")
	}
}
