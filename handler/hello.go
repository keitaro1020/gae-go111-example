package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Param string `json:"param"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("param")

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &HelloResponse{
		Message: "Hello world",
		Param: p,
	})
}
