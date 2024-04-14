//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"tickets/internal/broker"
	"tickets/internal/config"
	"tickets/internal/events"
	"tickets/internal/server"
	"tickets/internal/service/analytics"
	ticketservice "tickets/internal/service/ticket-service"
	"tickets/internal/storage/pg"
)

func Init() (*App, func(), error) {
	panic(
		wire.Build(
			newApp,
			wire.NewSet(config.New),
			wire.NewSet(initDB),

			wire.NewSet(events.NewStatusHandler),
			wire.NewSet(events.NewErrorHandler),
			wire.NewSet(events.NewItemHandler),
			wire.NewSet(events.NewOverpriceHandler),
			wire.NewSet(initBroker),

			wire.NewSet(pg.NewTicketStorage),
			wire.NewSet(ticketservice.New),

			wire.NewSet(analytics.New),
			wire.Bind(new(analytics.Summary), new(*pg.TicketStorage)),
			wire.Bind(new(server.SummaryService), new(*analytics.Analytics)),

			wire.Bind(new(server.TicketService), new(*ticketservice.TicketService)),
			wire.Bind(new(ticketservice.TicketStorage), new(*pg.TicketStorage)),

			wire.NewSet(server.New),
		),
	)
}

func initDB(cfg *config.Config) (*sqlx.DB, func(), error) {

	host := cfg.DB.Host
	port := cfg.DB.Port
	user := cfg.DB.User
	pass := cfg.DB.Pass
	name := cfg.DB.Name

	cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)

	log.Printf("connecting to %s\n", cs)

	db, err := sqlx.Open("postgres", cs)
	if err != nil {
		return nil, nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, func() {
			db.Close()
		}, err
	}

	return db, func() { db.Close() }, nil
}

func initBroker(cfg *config.Config,
	statusHandler events.StatusHandler,
	overpriceHandler events.OverpriceHandler,
	itemHandler events.ItemHandler,
	errorsHandler events.ErrorHandler) (broker.MessageBroker, func(), error) {
	mb, err := broker.New(cfg)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		if err := mb.Consume(statusHandler.Handle,
			overpriceHandler.Handle,
			itemHandler.Handle,
			errorsHandler.Handle); err != nil {
			log.Println(err)
			return
		}
	}()

	return mb, func() {
		mb.Close()
	}, nil
}
