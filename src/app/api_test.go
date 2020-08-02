package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const tableCreationQuery = `CREATE TABLE event(
	id serial PRIMARY KEY,
	operation VARCHAR (100) NOT NULL,
	description VARCHAR (300) NOT NULL,
	created_at TIMESTAMP NOT NULL
)`

var a App

func droptable() {
	a.DB.Exec("DROP TABLE event")
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM event")
	a.DB.Exec("ALTER SEQUENCE event_id_seq RESTART WITH 1")
}

func loadEvents() {
	a.DB.Exec("INSERT INTO event(operation, description, created_at) values ('operation1', 'operation for testing', now())")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	a.Initialize(os.Getenv(DbHost), os.Getenv(DbPort), os.Getenv(DbUser), os.Getenv(DbPassword), os.Getenv(DbName))
	droptable()
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "null" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestCreateEvent(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"operation": "saveuser", "description" : "add new user on the system"}`)
	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["operation"] != "saveuser" {
		t.Errorf("Expected event name to be 'saveuser'. Got '%v'", m["operation"])
	}

	if m["description"] != "add new user on the system" {
		t.Errorf("Expected event description to be 'add new user on the system'. Got '%v'", m["description"])
	}

}

func TestFindEventByOperation(t *testing.T) {
	clearTable()
	loadEvents()
	req, _ := http.NewRequest("GET", "/events/operation1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
