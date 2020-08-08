package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, data interface{}, status int) {
	w = prepareResponse(w, status)
	response, _ := json.Marshal(data)
	w.Write(response)
}

func responseError(w http.ResponseWriter, msgDescription string, status int) {
	message := Message{
		Description: msgDescription,
	}
	response(w, message, status)
}

func prepareResponse(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	return w
}

func checkErr(err error, w http.ResponseWriter, errorMessage string, statusCode int) bool {
	if err != nil {
		log.Println("error: ", err.Error())
		responseError(w, errorMessage, statusCode)
		return true
	}
	return false
}
