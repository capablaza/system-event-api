package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func getEvents(w http.ResponseWriter, r *http.Request) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if checkErr(err, w, InternalError) {
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT operation, description, created_at FROM event")
	if checkErr(err, w, InternalError) {
		return
	}

	var eventsCreated []Event

	for rows.Next() {
		var op string
		var desc string
		var created time.Time
		err = rows.Scan(&op, &desc, &created)
		if checkErr(err, w, InternalError) {
			return
		}

		nevent := Event{
			Operation:   op,
			Description: desc,
			CreatedAt:   created,
		}
		eventsCreated = append(eventsCreated, nevent)
	}

	response(w, r, eventsCreated, http.StatusOK)
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEvent)

	if checkErr(err, w, InternalError) {
		return
	}

	if validateEventInputData(w, newEvent) {
		return
	}

	currentTime := time.Now()
	newEvent.CreatedAt = currentTime

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	if checkErr(err, w, InternalError) {
		return
	}
	defer db.Close()

	var lastInsertId int
	err = db.QueryRow("INSERT INTO event(operation, description, created_at) VALUES($1,$2,$3)  returning id;", newEvent.Operation, newEvent.Description, newEvent.CreatedAt).Scan(&lastInsertId)
	if checkErr(err, w, InternalError) {
		return
	}

	var events []Event
	events = append(events, newEvent)
	response(w, r, events, http.StatusOK)
}

func validateEventInputData(w http.ResponseWriter, newEvent Event) bool {
	if newEvent.Operation == "" || newEvent.Description == "" {
		message := Message{
			Description: "Can't add empty events",
		}
		responseError(w, message, http.StatusBadRequest)
		return true
	}
	return false
}

func main() {
	router := mux.NewRouter()

	serviceName := "/events"

	router.HandleFunc(serviceName, getEvents).Methods(http.MethodGet)
	router.HandleFunc(serviceName, addEvent).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":5000", router))
}
