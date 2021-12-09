package main

import (
	"CinemaGCP/pkg/handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const IP_ADDRESS = "127.0.0.1"

type Server struct {
	httpServer *http.Server
}

func (s Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           IP_ADDRESS + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Printf("Server start on: %s:%s\n", IP_ADDRESS, port)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	srv := new(Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server was crushed: %v", err.Error())
	}
}
