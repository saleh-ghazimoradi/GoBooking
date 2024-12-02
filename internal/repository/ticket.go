package repository

import (
	"context"
	"database/sql"
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

type ticketRepo struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (t *ticketRepo) GetMany(ctx context.Context) ([]*service_models.Ticket, error) {
	return nil, nil
}

func (t *ticketRepo) GetOne(ctx context.Context, id int64) (*service_models.Ticket, error) {
	return nil, nil
}

func (t *ticketRepo) CreateOne(ctx context.Context, ticket *service_models.Ticket) error {
	return nil
}

func (t *ticketRepo) UpdateOne(ctx context.Context, ticket *service_models.Ticket) error {
	return nil
}

func (t *ticketRepo) DeleteOne(ctx context.Context, id int64) error {
	return nil
}

func (t *ticketRepo) GetWithTXT(tx *sql.Tx) Ticket {
	return &ticketRepo{
		dbWrite: t.dbWrite,
		dbRead:  t.dbRead,
		tx:      tx,
	}
}

func NewTicketRepository(dbWrite *sql.DB, dbRead *sql.DB) Ticket {
	return &ticketRepo{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
