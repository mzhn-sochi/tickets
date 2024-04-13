package events

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
	ticketservice "tickets/internal/service/ticket-service"
)

type ErrorHandler interface {
	Handle(msg *nats.Msg)
}

type errorHandler struct {
	storage ticketservice.TicketStorage
}

type errorReq struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
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

	ticketId := strings.TrimSpace(msg.Header.Get("ticket_id"))
	if ticketId == "" {
		log.Printf("empty ticket id\n")
		return
	}

	// TODO: Переделать статус
	if err := h.storage.PatchStatus(ticketId, "rejected"); err != nil {
		log.Println("cannot change status to rejected")
		return
	}

	if err := h.storage.AppendError(ticketId, data.Reason); err != nil {
		log.Println("cannot append error to ticket")
		return
	}
}
