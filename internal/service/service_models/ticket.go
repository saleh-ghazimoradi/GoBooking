package service_models

import "time"

type Ticket struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"event_id"`
	Event     Event     `json:"event"`
	Entered   bool      `json:"entered"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ValidateTicket struct {
	TicketId int64 `json:"ticket_id"`
}
