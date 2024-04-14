package events

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
	ticketservice "tickets/internal/service/ticket-service"
)

type OverpriceHandler interface {
	Handle(msg *nats.Msg)
}

type overpriceHandler struct {
	storage ticketservice.TicketStorage
}

func NewOverpriceHandler(storage ticketservice.TicketStorage) OverpriceHandler {
	return &overpriceHandler{storage: storage}
}

type overpriceReq struct {
	Percents uint `json:"percents"`
}

func (h *overpriceHandler) Handle(msg *nats.Msg) {
	var data overpriceReq
	log.Println(`received new msg: `, string(msg.Data))
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		log.Printf("error while unmarshalling overprice data %v\n", err)
		return
	}

	ticketId := strings.TrimSpace(msg.Header.Get("ticket_id"))
	if ticketId == "" {
		log.Printf("empty ticket id\n")
		return
	}

	if err := h.storage.SetOverprice(ticketId, data.Percents); err != nil {
		log.Printf("cannot set overprice\n")
		return
	}

}
