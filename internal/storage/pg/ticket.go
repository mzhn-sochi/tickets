package pg

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
	"tickets/internal/entity"
)

type TicketStorage struct {
	db *sqlx.DB
}

func NewTicketStorage(db *sqlx.DB) *TicketStorage {
	return &TicketStorage{db: db}
}

func (t *TicketStorage) Create(ticket *entity.Ticket) error {
	return t.db.Get(
		ticket,
		fmt.Sprintf("INSERT INTO %s(id, user_id, shop_address, image_url) VALUES ($1, $2, $3, $4) RETURNING *", TICKET_TABLE),
		ticket.Id,
		ticket.UserId,
		ticket.ShopAddress,
		ticket.ImageUrl,
	)
}

func (t *TicketStorage) Find(id string) (*entity.Ticket, error) {

	ticket := new(entity.Ticket)

	if err := t.db.Get(
		ticket,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1 LIMIT 1", TICKET_TABLE),
		id,
	); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *TicketStorage) List(filter *entity.Filter) ([]*entity.Ticket, int64, error) {
	tt := make([]*entity.Ticket, 0)

	query := squirrel.Select("*").
		From(TICKET_TABLE).
		Limit(filter.Limit).
		Offset(filter.Offset).
		OrderBy("created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	if filter != nil {
		if filter.Status != nil {
			st, err := filter.Query.Status.String()
			if err != nil {
				return nil, 0, fmt.Errorf("PostgreTicketStorage.List: %w", err)
			}
			query = query.Where(squirrel.Eq{"status": st})
		}

		if filter.Query.UserId != nil {
			query = query.Where(squirrel.Eq{"user_id": filter.UserId})
		}

		if filter.TimeRange.From != nil && filter.TimeRange.To != nil {
			query = query.Where(squirrel.GtOrEq{"created_at": filter.TimeRange.From}).
				Where(squirrel.LtOrEq{"created_at": filter.TimeRange.To})
		}

		if filter.TimeRange.From != nil {
			query = query.Where(squirrel.GtOrEq{"created_at": filter.TimeRange.From})
		}

		if filter.TimeRange.To != nil {
			query = query.Where(squirrel.LtOrEq{"created_at": filter.TimeRange.To})
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create sql query: %w", err)
	}

	log.Printf("executing sql: %s, args: %v\n", sql, args)

	if err := t.db.Select(&tt, sql, args...); err != nil {
		return nil, 0, err
	}

	var count int64
	if err := t.db.Get(&count, fmt.Sprintf("SELECT COUNT(id) FROM %s", TICKET_TABLE)); err != nil {
		return nil, 0, err
	}

	return tt, count, nil
}

func (t *TicketStorage) PatchStatus(id string, status string) error {
	_, err := t.db.Exec(
		fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", TICKET_TABLE),
		status,
		id,
	)

	return err
}

func (t *TicketStorage) AppendError(ticketId string, reason string) error {
	_, err := t.db.Exec(
		fmt.Sprintf("insert into %s (ticket_id, reason) values ($1, $2)", REJECTION_REASONS_TABLE),
		ticketId,
		reason,
	)

	return err
}
