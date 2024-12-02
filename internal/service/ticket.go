package service

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/GoBooking/internal/repository"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service/service_models"
)

type Ticket interface {
	GetMany(ctx context.Context) ([]*service_models.Ticket, error)
	GetOne(ctx context.Context, id int64) (*service_models.Ticket, error)
	CreateOne(ctx context.Context, ticket *service_models.Ticket) error
	UpdateOne(ctx context.Context, ticket *service_models.Ticket) error
	DeleteOne(ctx context.Context, id int64) error
	GetWithTXT(tx *sql.Tx) Ticket
}

type ticketService struct {
	ticketRepo repository.Ticket
}

func (s *ticketService) GetMany(ctx context.Context) ([]*service_models.Ticket, error) {
	return s.ticketRepo.GetMany(ctx)
}

func (s *ticketService) GetOne(ctx context.Context, id int64) (*service_models.Ticket, error) {
	return s.ticketRepo.GetOne(ctx, id)
}

func (s *ticketService) CreateOne(ctx context.Context, ticket *service_models.Ticket) error {
	return s.ticketRepo.CreateOne(ctx, ticket)
}

func (s *ticketService) UpdateOne(ctx context.Context, ticket *service_models.Ticket) error {
	return s.ticketRepo.UpdateOne(ctx, ticket)
}

func (s *ticketService) DeleteOne(ctx context.Context, id int64) error {
	return s.ticketRepo.DeleteOne(ctx, id)
}

func (s *ticketService) GetWithTXT(tx *sql.Tx) Ticket {
	return &ticketService{
		ticketRepo: s.ticketRepo.GetWithTXT(tx),
	}
}

func NewTicketService(ticketRepo repository.Ticket) Ticket {
	return &ticketService{
		ticketRepo: ticketRepo,
	}
}
