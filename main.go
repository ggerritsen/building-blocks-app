package main

import (
	"log"
	"net/http"

	"github.com/ggerritsen/building-blocks-app/handler"
	"github.com/ggerritsen/building-blocks-app/kafkaclient"
	"github.com/ggerritsen/building-blocks-app/repository"
	"github.com/ggerritsen/building-blocks-app/service"
)

func main() {
	log.Printf("Starting app...\n")

	repo, err := repository.NewRepositoryWithDb("localhost", "postgres", "your-password", "template1", 5432)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	log.Printf("Connected to database\n")

	if err := repo.CreateTable(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Initialized database\n")

	svc := service.NewDocService(repo)
	stop := make(chan error, 2)

	c := kafkaclient.NewConsumer([]string{"localhost:9092"}, "test", svc)
	go func() {
		stop <- c.Consume()
	}()
	log.Printf("Started kafka consumer\n")

	h := handler.NewHandler(svc)
	go func() {
		stop <- http.ListenAndServe(":8081", h)
	}()
	log.Printf("Started http handler\n")

	err = <-stop
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("App stopped.\n")
}
