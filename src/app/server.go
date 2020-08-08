package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Printf(connectionString)
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
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))

	a.Router.HandleFunc(serviceName, a.getEvents).Methods(http.MethodGet)
	a.Router.HandleFunc(serviceName+"/{operation}", a.findEventByOperation).Methods(http.MethodGet)
	a.Router.HandleFunc(serviceName, a.addEvent).Methods(http.MethodPost)

}

func (a *App) Run(address string) {
	log.Printf("Running service in port : %s", address)
	log.Fatal(http.ListenAndServe(address, a.Router))
}
