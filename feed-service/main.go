package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"platzi.com/go/cqrs/database"
	"platzi.com/go/cqrs/events"
	"platzi.com/go/cqrs/repository"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Failed to process env config: %v", err)
	}

	addr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresDB)

	repo, err := database.NewPostgresRepository(addr)
	if err != nil {
		log.Fatalf("Failed to create postgres repository: %v", err)
	}

	repository.SetRepository(repo)

	n, err := events.NewNats(fmt.Sprintf("nats://%s", config.NatsAddress))
	if err != nil {
		log.Fatalf("Failed to create nats event store: %v", err)
	}
	events.SetEventStore(n)

	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
