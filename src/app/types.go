package main

import (
	"time"
)

type Event struct {
	Operation     string    `json:"operation"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	OperationDate string    `json:"creation_date"`
}

type Message struct {
	Description string `json:"description"`
}
