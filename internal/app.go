package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yosa12978/pngb/internal/web"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	web.Run()

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	log.Printf("Server stopped. Signal: %s", <-out)
}
