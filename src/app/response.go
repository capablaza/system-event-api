package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, data []Event, status int) {
	w = prepareResponse(w, status)
	json.NewEncoder(w).Encode(data)
}

func responseError(w http.ResponseWriter, message Message, status int) {
	w = prepareResponse(w, status)
	json.NewEncoder(w).Encode(message)
}

func prepareResponse(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return w
}

func checkErr(err error, w http.ResponseWriter, errorMessage string) bool {
	if err != nil {
		message := Message{
			Description: errorMessage,
		}
		log.Println("error: ", err.Error())
		responseError(w, message, http.StatusBadRequest)
		return true
	}
	return false
}
