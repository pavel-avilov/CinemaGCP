package main

import (
	"CinemaGCP/pkg/handler"
	"CinemaGCP/pkg/repository"
	service2 "CinemaGCP/pkg/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
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
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "34.88.5.135",
		Port:     "5432",
		Username: "postgres",
		Password: os.Getenv("PASSWORD_DB"),
		DBName:   "cinema",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("fatal to initialize db: %v", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service2.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server was crushed: %v", err.Error())
	}
}
