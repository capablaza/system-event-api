package main

import (
	"time"
)

type Event struct {
	Operation   string    `json:"operation"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Message struct {
	Description string `json:"description"`
}
