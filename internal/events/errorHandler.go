package events

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	ticketservice "tickets/internal/service/ticket-service"
)

type ErrorHandler interface {
	Handle(msg *nats.Msg)
}

type errorHandler struct {
	storage ticketservice.TicketStorage
}

type errorReq struct {
	TicketId string `json:"ticket_id"`
	Type     string `json:"type"`
	Reason   string `json:"reason"`
}

func NewErrorHandler(storage ticketservice.TicketStorage) ErrorHandler {
	return &errorHandler{storage: storage}
}

func (h *errorHandler) Handle(msg *nats.Msg) {
	var data errorReq
	log.Println(`received new error msg: `, string(msg.Data))
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		log.Printf("error while unmarshalling error data %v\n", err)
		return
	}
	// TODO: Доделать ошибки
}
