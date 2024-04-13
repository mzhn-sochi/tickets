package ticketservice

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"tickets/internal/broker"
	"tickets/internal/config"
	"tickets/internal/entity"
	"tickets/internal/server"
)

var _ server.TicketService = (*TicketService)(nil)

type TicketStorage interface {
	Create(ticket *entity.Ticket) error
	Find(id string) (*entity.Ticket, error)
	List(filter *entity.Filter) ([]*entity.Ticket, int64, error)
	PatchStatus(id string, status string) error
	AppendError(ticketId string, reason string) error
}

type TicketService struct {
	TicketStorage

	cfg *config.Config
	mb  broker.MessageBroker
}

func New(storage TicketStorage, mb broker.MessageBroker, cfg *config.Config) *TicketService {
	return &TicketService{TicketStorage: storage, mb: mb, cfg: cfg}
}

func (t *TicketService) Create(ticket *entity.Ticket) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	req := struct {
		ImageUrl string `json:"image_url"`
	}{ImageUrl: ticket.ImageUrl}

	var data []byte
	if data, err = json.Marshal(req); err != nil {
		return err
	}

	ticket.Id = id.String()
	if err := t.mb.Publish(
		&nats.Msg{
			Subject: t.cfg.Nats.Queues.OCR,
			Header:  map[string][]string{"ticket_id": {id.String()}},
			Data:    data,
		}); err != nil {
		return err
	}

	return t.TicketStorage.Create(ticket)
}

func (t *TicketService) PatchStatus(id string, status entity.Status) error {

	s, err := status.String()
	if err != nil {
		return err
	}

	return t.TicketStorage.PatchStatus(id, s)
}
