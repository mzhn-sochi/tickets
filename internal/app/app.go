package app

import (
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"tickets/internal/broker"

	"google.golang.org/grpc"
	"tickets/api/ts"
	"tickets/internal/config"
	"tickets/internal/server"
)

type App struct {
	cfg *config.Config

	impl          *server.Server
	MessageBroker broker.MessageBroker
}

func (a *App) Run() {
	s := grpc.NewServer()
	reflection.Register(s)

	ts.RegisterTicketServiceServer(s, a.impl)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.App.Host, a.cfg.App.Port))
		if err != nil {
			panic(fmt.Errorf("cannot bind port %d", a.cfg.App.Port))
		}

		log.Printf("server started at %s:%d", a.cfg.App.Host, a.cfg.App.Port)
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	sig := <-sigChan
	s.GracefulStop()
	log.Printf("Signal %v received, stopping server...\n", sig)
}

func newApp(cfg *config.Config, impl *server.Server, broker broker.MessageBroker) *App {
	return &App{cfg: cfg, impl: impl,
		MessageBroker: broker}
}
