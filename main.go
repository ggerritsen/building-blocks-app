package main

import (
	"log"
	"net/http"

	"github.com/ggerritsen/building-blocks-app/handler"
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
	h := handler.NewHandler(svc)

	c := make(chan error, 1)
	go func() {
		c <- http.ListenAndServe(":8081", h)
	}()

	err = <-c
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("App stopped.\n")
}
