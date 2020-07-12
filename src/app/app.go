package main

import (
	base "./base"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Event = base.Event
type Message = base.Message

const (
	DB_USER     = "logmaster"
	DB_PASSWORD = "9!h$%Ple1"
	DB_NAME     = "events"
)

func response(w http.ResponseWriter, r *http.Request, data []Event, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func responseError(w http.ResponseWriter, r *http.Request, message Message, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT operation, description, created_at FROM event")
	checkErr(err)

	var eventsCreated []Event

	for rows.Next() {
		var op string
		var desc string
		var created time.Time
		err = rows.Scan(&op, &desc, &created)
		checkErr(err)

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

	checkErr(err)

	if validateEventInputData(w, r, newEvent) {
		return
	}

	currentTime := time.Now()
	newEvent.CreatedAt = currentTime

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var lastInsertId int
	err = db.QueryRow("INSERT INTO event(operation, description, created_at) VALUES($1,$2,$3)  returning id;", newEvent.Operation, newEvent.Description, newEvent.CreatedAt).Scan(&lastInsertId)
	checkErr(err)

	log.Println(newEvent)

	var events []Event
	events = append(events, newEvent)
	response(w, r, events, http.StatusOK)
}

func validateEventInputData(w http.ResponseWriter, r *http.Request, newEvent Event) bool {
	if newEvent.Operation == "" || newEvent.Description == "" {
		message := Message{
			Description: "Can't add empty events",
		}
		responseError(w, r, message, http.StatusBadRequest)
		return true
	}
	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/events", getEvents).Methods("GET")
	router.HandleFunc("/events", addEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}
