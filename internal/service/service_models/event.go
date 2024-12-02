package service_models

import "time"

type Event struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

type EventPayload struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type UpdateEventPayload struct {
	Name     *string `json:"name" validate:"omitempty,max=35"`
	Location *string `json:"location" validate:"omitempty,max=35"`
}
