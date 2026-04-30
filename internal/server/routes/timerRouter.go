package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"timer-api/pkg/domain/response"
	"timer-api/pkg/domain/timer"
	"timer-api/pkg/useCases/Handlers/timerHandler"
	"timer-api/pkg/useCases/Helpers/responseHelper"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type TimerRouter struct {
}

func (tr *TimerRouter) CreateTimer(w http.ResponseWriter, r *http.Request) {
	var timer timer.Timer
	err := json.NewDecoder(r.Body).Decode(&timer)
	if err != nil {
		responseHelper.WriteResponse(w, response.StatusBadRequest, nil)
		return
	}

	newTimer, status := timerHandler.CreateTimer(timer)
	responseHelper.WriteResponse(w, status, newTimer)
}

func (tr *TimerRouter) GetLastTimer(w http.ResponseWriter, r *http.Request) {

	timer, status := timerHandler.GetLastTimer()
	responseHelper.WriteResponse(w, status, timer)
}

func (tr *TimerRouter) DeleteTimer(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseHelper.WriteResponse(w, response.StatusBadRequest, nil)
		return
	}

	status := timerHandler.DeleteTimer(id)
	responseHelper.WriteResponse(w, status, nil)
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

	// DELETE /{id} - hard delete, allowed only from /start view (checked via header)
	r.Delete("/{id}", tr.DeleteTimer)

	return r
}
