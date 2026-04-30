package routes

import (
	"encoding/json"
	"net/http"

	"timer-api/pkg/domain/response"
	"timer-api/pkg/domain/timer"
	timerHandler "timer-api/pkg/useCases/Handlers/timerHandler"
	"timer-api/pkg/useCases/Helpers/responseHelper"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type TimerRouter struct {
	Handler timerHandler.Handler
}

func (tr *TimerRouter) CreateTimer(w http.ResponseWriter, r *http.Request) {
	var timer timer.Timer
	err := json.NewDecoder(r.Body).Decode(&timer)
	if err != nil {
		responseHelper.WriteResponse(w, response.StatusBadRequest, nil)
		return
	}

	status := tr.Handler.CreateTimer(timer)
	responseHelper.WriteResponse(w, status, timer)
}

func (tr *TimerRouter) GetLastTimer(w http.ResponseWriter, r *http.Request) {

	timer, status := tr.Handler.GetLastTimer()
	responseHelper.WriteResponse(w, status, timer)
}

func (tr *TimerRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
		MaxAge:             300,
	}))

	r.Post("/", tr.CreateTimer)

	r.Get("/", tr.GetLastTimer)

	return r
}
