package events

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
	"tickets/internal/entity"
	ticketservice "tickets/internal/service/ticket-service"
)

const (
	OCRType        = "ocr"
	ValidationType = "validation"
)

type StatusHandler interface {
	Handle(msg *nats.Msg)
}

type statusHandler struct {
	storage ticketservice.TicketStorage
}

type statusReq struct {
	OperationType string `json:"type"`
}

func NewStatusHandler(storage ticketservice.TicketStorage) StatusHandler {
	return &statusHandler{storage: storage}
}

func (h *statusHandler) Handle(msg *nats.Msg) {
	var data statusReq
	log.Println(`received new msg: `, string(msg.Data))
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		log.Printf("error while unmarshalling status data %v\n", err)
		return
	}

	ticketId := strings.TrimSpace(msg.Header.Get("ticket_id"))
	if ticketId == "" {
		log.Printf("empty ticket id\n")
		return
	}

	var status entity.Status

	switch data.OperationType {
	case OCRType:
		status = entity.StatusWaitingValidation
		break
	case ValidationType:
		status = entity.StatusWaitingApproval
		break
	}

	s, err := status.String()
	if err != nil {
		log.Println("cannot parse status")
		return
	}

	if err := h.storage.PatchStatus(ticketId, s); err != nil {
		log.Println(s, err)
		return
	}

	log.Printf("tx: %s status successfully changed to %s\n", ticketId, s)
}
