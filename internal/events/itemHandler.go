package events

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
	"tickets/internal/entity"
	ticketservice "tickets/internal/service/ticket-service"
)

type ItemHandler interface {
	Handle(msg *nats.Msg)
}

type itemHandler struct {
	storage ticketservice.TicketStorage
}

func NewItemHandler(storage ticketservice.TicketStorage) ItemHandler {
	return &itemHandler{storage: storage}
}

func (h *itemHandler) Handle(msg *nats.Msg) {
	var data entity.Item
	log.Println(`received new error msg: `, string(msg.Data))
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		log.Printf("error while unmarshalling item data %v\n", err)
		return
	}

	ticketId := strings.TrimSpace(msg.Header.Get("ticket_id"))
	if ticketId == "" {
		log.Printf("empty ticket id\n")
		return
	}

	if err := h.storage.AddItem(ticketId, &data); err != nil {
		log.Println(err)
		log.Printf("cannot add items to ticketId %s\n", ticketId)
		return
	}
}
