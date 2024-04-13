package server

import (
	"context"
	"strings"
	"tickets/api/share"
	"tickets/api/ts"
	"tickets/internal/entity"
)

var _ ts.TicketServiceServer = (*Server)(nil)

type TicketService interface {
	Create(ticket *entity.Ticket) error
	List(filter *entity.Filter) ([]*entity.Ticket, int64, error)
	Find(id string) (*entity.Ticket, error)
	PatchStatus(id string, status entity.Status) error
}

type Server struct {
	service TicketService
	ts.UnimplementedTicketServiceServer
}

func New(service TicketService) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Create(ctx context.Context, request *ts.CreateRequest) (*ts.CreateResponse, error) {
	ticket := &entity.Ticket{
		UserId:      request.UserId,
		ShopAddress: request.ShopAddr,
		ImageUrl:    request.ImageUrl,
	}

	if err := s.service.Create(ticket); err != nil {
		return nil, err
	}

	return &ts.CreateResponse{
		TicketId: ticket.Id,
	}, nil
}

func (s *Server) List(ctx context.Context, request *ts.ListRequest) (*ts.ListResponse, error) {

	filter := new(entity.Filter)

	if request.Bounds != nil {
		filter.Bounds = entity.Bounds{
			Limit:  request.Bounds.Limit,
			Offset: request.Bounds.Offset,
		}
	}

	if request.Filter != nil {
		if request.Filter.Status != nil {

			var st entity.Status
			switch *request.Filter.Status {
			case ts.Statuses_WAITING_OCR:
				st = entity.StatusWaitingOcr
			case ts.Statuses_WAITING_VALIDATION:
				st = entity.StatusWaitingValidation
			case ts.Statuses_WAITING_APPROVAL:
				st = entity.StatusWaitingApproval
			case ts.Statuses_CLOSED:
				st = entity.StatusClosed
			case ts.Statuses_REJECTED:
				st = entity.StatusRejected
			}

			filter.Status = &st
		}

		if request.Filter.TimeRange != nil {
			tr := request.Filter.TimeRange
			if tr.To != nil && tr.From != nil {
				filter.TimeRange = entity.TimeRange{
					From: tr.From,
					To:   tr.To,
				}
			}

			if tr.From != nil {
				filter.TimeRange = entity.TimeRange{
					From: tr.From,
				}
			}

			if tr.To != nil {
				filter.TimeRange = entity.TimeRange{
					To: tr.To,
				}
			}
		}

		if request.Filter.UserId != nil {
			filter.UserId = request.Filter.UserId
		}
	}

	tickets, count, err := s.service.List(filter)
	if err != nil {
		return nil, err
	}

	tt := make([]*ts.Ticket, 0, len(tickets))

	for _, t := range tickets {

		tick := &ts.Ticket{
			Id:          t.Id,
			UserId:      t.UserId,
			ImageUrl:    t.ImageUrl,
			ShopAddress: t.ShopAddress,
			CreatedAt:   t.CreatedAt.Unix(),
			Status:      ts.Statuses(ts.Statuses_value[strings.ToUpper(t.Status)]),
			Reason:      t.Reason,
		}

		if t.UpdatedAt != nil {
			tick.UpdatedAt = new(int64)
			*tick.UpdatedAt = t.UpdatedAt.Unix()
		}

		tt = append(tt, tick)
	}

	return &ts.ListResponse{
		Tickets: tt,
		Count:   count,
	}, nil
}

func (s *Server) FindById(ctx context.Context, request *ts.FindByIdRequest) (*ts.Ticket, error) {
	ticket, err := s.service.Find(request.TicketId)
	if err != nil {
		return nil, err
	}

	t := &ts.Ticket{
		Id:          ticket.Id,
		UserId:      ticket.UserId,
		ImageUrl:    ticket.ImageUrl,
		ShopAddress: ticket.ShopAddress,
		CreatedAt:   ticket.CreatedAt.Unix(),
		Reason:      ticket.Reason,
	}

	if ticket.UpdatedAt != nil {
		t.UpdatedAt = new(int64)
		*t.UpdatedAt = ticket.UpdatedAt.Unix()
	}

	return t, nil
}

func (s *Server) CloseTicket(ctx context.Context, request *ts.CloseTicketRequest) (*share.Empty, error) {
	if err := s.service.PatchStatus(request.TicketId, entity.StatusClosed); err != nil {
		return nil, err
	}

	return &share.Empty{}, nil
}
