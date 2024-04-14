package server

import (
	"errors"
	"tickets/api/ts"
	"tickets/internal/service/analytics"
)

func (s *Server) GetUserSummary(_ *ts.Empty, server ts.TicketService_GetUserSummaryServer) error {

	rows, err := s.summary.UserSummary()
	if err != nil {
		if errors.Is(err, analytics.ErrNotFound) {
			return err
		}
		return err
	}

	for userId, count := range rows {
		if err := server.Send(&ts.UserSummary{UserId: userId, Count: count}); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetStatusSummary(_ *ts.Empty, server ts.TicketService_GetStatusSummaryServer) error {

	rows, err := s.summary.StatusSummary()
	if err != nil {
		if errors.Is(err, analytics.ErrNotFound) {
			return err
		}
		return err
	}

	for status, count := range rows {
		if err := server.Send(&ts.StatusSummary{StatusId: status, Count: count}); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetShopSummary(_ *ts.Empty, server ts.TicketService_GetShopSummaryServer) error {
	rows, err := s.summary.ShopSummary()
	if err != nil {
		if errors.Is(err, analytics.ErrNotFound) {
			return err
		}
		return err
	}

	for shop, count := range rows {
		if err := server.Send(&ts.ShopSummary{ShopId: shop, Count: count}); err != nil {
			return err
		}
	}

	return nil
}
