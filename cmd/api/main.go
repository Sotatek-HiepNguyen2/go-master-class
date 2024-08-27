package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
}

func main() {
	var cfg Config

	cfg.Port = 4000

	logger := log.New(log.Writer(), "API ", log.LstdFlags)

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		ErrorLog:     logger,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := server.ListenAndServe()
	logger.Fatal(err)
}
