package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service/service_models"
)

type Ticket interface {
	GetMany(ctx context.Context, fq service_models.PaginationFeedQuery) ([]service_models.Ticket, error)
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

func (t *ticketRepo) GetMany(ctx context.Context, fq service_models.PaginationFeedQuery) ([]service_models.Ticket, error) {
	query := `SELECT id, event_id, entered, created_at, updated_at FROM tickets ORDER BY id ` + fq.Sort + ` LIMIT $1 OFFSET $2`

	rows, err := t.dbRead.QueryContext(ctx, query, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []service_models.Ticket
	for rows.Next() {
		var ticket service_models.Ticket
		err = rows.Scan(
			&ticket.ID,
			&ticket.EventID,
			&ticket.Entered,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (t *ticketRepo) GetOne(ctx context.Context, id int64) (*service_models.Ticket, error) {
	query := `SELECT id, event_id, entered, created_at, updated_at FROM tickets WHERE id=$1`

	var ticket service_models.Ticket
	err := t.dbRead.QueryRowContext(ctx, query, id).Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.Entered,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &ticket, nil

}

func (t *ticketRepo) CreateOne(ctx context.Context, ticket *service_models.Ticket) error {
	query := `INSERT INTO tickets (event_id, entered) VALUES ($1, $2) RETURNING id, created_at, updated_at`

	err := t.dbWrite.QueryRowContext(ctx, query, ticket.EventID, ticket.Entered).Scan(&ticket.ID, &ticket.CreatedAt, &ticket.UpdatedAt)
	if err != nil {
		return err
	}

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
