package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("param")
	fmt.Fprintf(w, "Hello, World! : %s", p)
}
