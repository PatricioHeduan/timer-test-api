package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
	ERR_ALREADY_COMMITTED = "already been committed"
)

func New() http.Handler {
	r := chi.NewRouter()

	tr := TimerRouter{}

	r.Mount("/timer", tr.Routes())
	return r
}
