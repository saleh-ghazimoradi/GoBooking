package gateway

import (
	"github.com/saleh-ghazimoradi/GoBooking/internal/service"
	"net/http"
)

type ticket struct {
	ticketService service.Ticket
}

func (t *ticket) getOneTicket(w http.ResponseWriter, r *http.Request) {}

func (t *ticket) getManyTickets(w http.ResponseWriter, r *http.Request) {}

func (t *ticket) createOneTicket(w http.ResponseWriter, r *http.Request) {}

func (t *ticket) updateOneTicket(w http.ResponseWriter, r *http.Request) {}

func (t *ticket) deleteOneTicket(w http.ResponseWriter, r *http.Request) {}

func (t *ticket) validateTicket(w http.ResponseWriter, r *http.Request) {}

func NewTicketHandler(ticketService service.Ticket) *ticket {
	return &ticket{
		ticketService: ticketService,
	}
}
