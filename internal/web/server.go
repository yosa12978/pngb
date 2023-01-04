package web

import (
	"net/http"
	"os"

	"github.com/yosa12978/pngb/internal/web/handlers"
)

func Run() {
	server := http.Server{
		Addr:           os.Getenv("ADDR"),
		Handler:        handlers.NewHandler(),
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
}
