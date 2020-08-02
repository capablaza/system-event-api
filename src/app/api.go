package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func (a *App) findEventByOperation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation := vars["operation"]
	log.Println("operation: " + operation)
	event := Event{Operation: operation}

	if err := event.getEventByOperation(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			responseError(w, "Event not found", http.StatusNotFound)
		default:
			responseError(w, "Service not Available", http.StatusServiceUnavailable)
		}
		return
	}

	response(w, event, http.StatusOK)

}

func (a *App) getEvents(w http.ResponseWriter, r *http.Request) {
	events, err := getAllEvents(a.DB)
	if checkErr(err, w, InternalError, http.StatusServiceUnavailable) {
		return
	}
	response(w, events, http.StatusOK)
}

func (a *App) addEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		responseError(w, BadRequestError, http.StatusBadRequest)
		return
	}

	if err := validateEventInputData(event); err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	event.CreatedAt = currentTime

	defer r.Body.Close()

	if err := event.addEvent(a.DB); err != nil {
		responseError(w, InternalError, http.StatusInternalServerError)
		return
	}

	response(w, event, http.StatusOK)
}

func validateEventInputData(newEvent Event) error {
	if newEvent.Operation == "" || newEvent.Description == "" {
		return errors.New("can't add empty events")
	}
	return nil
}

func main() {

	a := &App{}
	appPort := ":5000"
	a.Initialize(DbUser, DbPassword, DbName)
	a.Run(appPort)
}
