package main

import (
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
	a.Initialize(DbUser, DbPassword, DbName)
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
