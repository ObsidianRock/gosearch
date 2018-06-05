package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ObsidianRock/gosearch/config"
	"github.com/ObsidianRock/gosearch/handler"
	"github.com/ObsidianRock/gosearch/storage/sqlite"
)

func main() {

	configPath := flag.String("config", "./config/config.json", "location of configuration file")

	flag.Parse()

	config, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	svc, err := sqlite.New(config.SQLite.Path)
	if err != nil {
		log.Fatal(err)
	}

	defer svc.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: handler.New(config.Option.Prefix, svc),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

}
