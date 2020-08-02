package main

import (
	"database/sql"
	"time"
)

func (e *Event) getEventByOperation(db *sql.DB) error {
	queryStr := "SELECT operation, description, created_at FROM event WHERE operation = $1"
	return db.QueryRow(queryStr, e.Operation).Scan(&e.Operation, &e.Description, &e.CreatedAt)
}

func getAllEvents(db *sql.DB) ([]Event, error) {
	queryStr := "SELECT operation, description, created_at FROM event"
	rows, err := db.Query(queryStr)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var eventsCreated []Event

	for rows.Next() {
		var op string
		var desc string
		var created time.Time

		if err = rows.Scan(&op, &desc, &created); err != nil {
			return nil, err
		}

		event := Event{
			Operation:   op,
			Description: desc,
			CreatedAt:   created,
		}

		eventsCreated = append(eventsCreated, event)
	}

	return eventsCreated, nil

}

func (e *Event) addEvent(db *sql.DB) error {

	var lastInsertId int
	err := db.QueryRow("INSERT INTO event(operation, description, created_at) VALUES($1,$2,$3)  returning id;", e.Operation, e.Description, e.CreatedAt).Scan(&lastInsertId)

	if err != nil {
		return err
	}

	return nil

}
