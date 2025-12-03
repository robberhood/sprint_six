package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "UPLOAD->", log.LstdFlags)

	srv := server.CreateHTTPServer(logger)

	srv.Logger.Printf("Server starting... port: %s", srv.Server.Addr)
	if err := srv.Server.ListenAndServe(); err != nil {
		srv.Logger.Fatalf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}
