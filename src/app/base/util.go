package base

import (
	"strconv"
	"time"
)

func loadData(events []Event) {
	for i := 1; i < 6; i++ {
		op := "operation " + string(strconv.Itoa(i))
		currentTime := time.Now()
		event := Event{Operation: op, CreatedAt: currentTime}
		events = append(events, event)
	}
}