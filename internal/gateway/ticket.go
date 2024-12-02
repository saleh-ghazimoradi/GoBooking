package gateway

import (
	"github.com/saleh-ghazimoradi/GoBooking/internal/service"
	"net/http"
)

type ticket struct {
	ticketService service.Ticket
}

func getOneTicket(w http.ResponseWriter, r *http.Request) {}

func getManyTickets(w http.ResponseWriter, r *http.Request) {}

func createOneTicket(w http.ResponseWriter, r *http.Request) {}

func updateOneTicket(w http.ResponseWriter, r *http.Request) {}

func deleteOneTicket(w http.ResponseWriter, r *http.Request) {}

func NewTicketHandler(ticketService service.Ticket) *ticket {
	return &ticket{
		ticketService: ticketService,
	}
}
