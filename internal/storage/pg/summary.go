package pg

import (
	"github.com/Masterminds/squirrel"
	"log"
)

func (t *TicketStorage) StatusSummary() (map[string]int64, error) {

	r := make(map[string]int64)

	query, args, err := squirrel.Select("t.status, count(t.id)").
		From(TICKET_TABLE + " t").
		GroupBy("t.status").
		OrderBy("count(t.id)").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Printf("executing query: %s", query)

	rows, err := t.db.Query(query, args...)
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int64
		err = rows.Scan(&status, &count)
		if err != nil {
			log.Printf("error scanning row: %s", err)
			continue
		}
		r[status] = count
	}

	return r, nil
}

func (t *TicketStorage) ShopSummary() (map[string]int64, error) {

	r := make(map[string]int64)

	query, args, err := squirrel.Select("t.shop_name, count(t.id)").
		From(TICKET_TABLE + " t").
		GroupBy("t.shop_name").
		OrderBy("count(t.id)").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Printf("executing query: %s", query)

	rows, err := t.db.Query(query, args...)
	defer rows.Close()

	for rows.Next() {
		var shop string
		var count int64
		err = rows.Scan(&shop, &count)
		if err != nil {
			log.Printf("error scanning row: %s", err)
			continue
		}
		r[shop] = count
	}

	return r, nil
}

func (t *TicketStorage) UserSummary() (map[string]int64, error) {

	r := make(map[string]int64)

	query, args, err := squirrel.Select("t.user_id, count(t.id)").
		From(TICKET_TABLE + " t").
		GroupBy("t.user_id").
		OrderBy("count(t.id)").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Printf("executing query: %s", query)

	rows, err := t.db.Query(query, args...)
	defer rows.Close()

	for rows.Next() {
		var userId string
		var count int64
		err = rows.Scan(&userId, &count)
		if err != nil {
			log.Printf("error scanning row: %s", err)
			continue
		}
		r[userId] = count
	}

	return r, nil
}
