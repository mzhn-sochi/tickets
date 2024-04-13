package analytics

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("no data found")

type Summary interface {
	UserSummary() (map[string]int64, error)
	ShopSummary() (map[string]int64, error)
	StatusSummary() (map[string]int64, error)
}

type Analytics struct {
	s Summary
}

func New(s Summary) *Analytics {
	return &Analytics{s: s}
}

func (a *Analytics) UserSummary() (map[string]int64, error) {
	r, err := a.s.UserSummary()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error at UserSummary: %w", err)
	}

	return r, nil
}

func (a *Analytics) ShopSummary() (map[string]int64, error) {
	r, err := a.s.ShopSummary()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error at ShopSummary: %w", err)
	}

	return r, nil
}

func (a *Analytics) StatusSummary() (map[string]int64, error) {
	r, err := a.s.StatusSummary()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error at StatusSummary: %w", err)
	}

	return r, nil
}
