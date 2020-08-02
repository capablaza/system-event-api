package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.routeHandle()
}

func (a *App) routeHandle() {
	serviceName := "/events"
	a.Router.HandleFunc(serviceName, a.getEvents).Methods(http.MethodGet)
	a.Router.HandleFunc(serviceName+"/{operation}", a.findEventByOperation).Methods(http.MethodGet)
	a.Router.HandleFunc(serviceName, a.addEvent).Methods(http.MethodPost)
}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}
