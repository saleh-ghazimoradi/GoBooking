package service_models

import "time"

type Event struct {
	ID        int64
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
