package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateHTTPServer(log *log.Logger) Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.HandlerGetHTML)
	router.HandleFunc("/upload", handlers.HandlerUpload)

	serv := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return Server{
		Logger: log,
		Server: &serv,
	}
}
