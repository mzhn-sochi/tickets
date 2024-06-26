package pg

import (
	"context"
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
		fmt.Sprintf("INSERT INTO %s(id, user_id, shop_name, shop_address, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING *", TICKET_TABLE),
		ticket.Id,
		ticket.UserId,
		ticket.ShopName,
		ticket.ShopAddress,
		ticket.ImageUrl,
	)
}

func (t *TicketStorage) Find(id string) (*entity.Ticket, error) {

	ticket := new(entity.Ticket)

	query, args, err := squirrel.Select(fmt.Sprintf("%s.*, %s.reason", TICKET_TABLE, REJECTION_REASONS_TABLE)).
		From(TICKET_TABLE).
		Limit(1).
		Where(squirrel.Eq{"id": id}).
		OrderBy("created_at DESC").
		LeftJoin(REJECTION_REASONS_TABLE).
		JoinClause(fmt.Sprintf("ON %s.ticket_id = %s.id", REJECTION_REASONS_TABLE, TICKET_TABLE)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err := t.db.Get(ticket, query, args...); err != nil {
		return nil, err
	}

	var item entity.Item
	query = fmt.Sprintf("select ti.product, ti.description, ti.price, ti.amount, ti.unit, ti.overprice from %s ti where ti.ticket_id = $1 limit 1;", TICKETS_ITEM_TABLE)
	err = t.db.QueryRow(query, id).Scan(&item.Product, &item.Description, &item.Price, &item.Measure.Amount, &item.Measure.Unit, &item.Overprice)
	if err == nil {
		ticket.Item = &item
	}

	return ticket, nil
}

func (t *TicketStorage) List(filter *entity.Filter) ([]*entity.Ticket, int64, error) {
	tt := make([]*entity.Ticket, 0)

	query := squirrel.Select(fmt.Sprintf("%s.*, %s.reason", TICKET_TABLE, REJECTION_REASONS_TABLE)).
		From(TICKET_TABLE).
		Limit(filter.Limit).
		Offset(filter.Offset).
		OrderBy("created_at DESC").
		LeftJoin(REJECTION_REASONS_TABLE).
		JoinClause(fmt.Sprintf("ON %s.ticket_id = %s.id", REJECTION_REASONS_TABLE, TICKET_TABLE)).
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

	query, args, err := squirrel.Update(TICKET_TABLE).
		Set("status", status).
		Set("updated_at", squirrel.Expr("now()")).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to create sql query: %w", err)
	}

	if _, err := t.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func (t *TicketStorage) AppendError(ticketId string, reason string) error {
	log.Println("reason: ", reason)
	_, err := t.db.Exec(
		fmt.Sprintf("insert into %s (ticket_id, reason) values ($1, $2)", REJECTION_REASONS_TABLE),
		ticketId,
		reason,
	)

	return err
}

func (t *TicketStorage) SetOverprice(ticketId string, overprice uint) error {
	tx, err := t.db.BeginTx(context.Background(), nil)
	if err != nil {
		fmt.Println("cannot start tx")
	}

	var count int
	if err := tx.QueryRow(`select count() from tickets_item where ticket_id=?;`, ticketId).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		query := fmt.Sprintf(`update tickets_item i set overprice=? where i.ticket_id=?;`)
		_, err := tx.Exec(query, overprice, ticketId)
		if err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf(`insert into tickets_item(ticket_id, overprice) values (?, ?);`)
		_, err := tx.Exec(query, ticketId, overprice)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (t *TicketStorage) AddItem(ticketId string, item *entity.Item) error {
	query := fmt.Sprintf(`insert into %s(ticket_id, product, description, price, amount, unit) values ($1,$2,$3,$4,$5,$6);`, TICKETS_ITEM_TABLE)
	_, err := t.db.Exec(query, ticketId, item.Product, item.Description, item.Price, item.Measure.Amount, item.Measure.Unit)

	return err
}
