package handler

import (
	"net/http"
	"fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("param")
	fmt.Fprintf(w, "Hello, World! : %s", p)
}
