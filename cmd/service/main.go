package main

import (
	"flag"
	"log"
	"tickets/internal/app"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", false, "local dev mode")
}

func main() {
	flag.Parse()

	a, cleanup, err := app.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	a.Run()

}
