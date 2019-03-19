package main

import (
	"log"

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

	// Move this to http handler code
	// d1, err := svc.Store("testDoc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Inserted document %+v\n", d1)

	// d2, err := svc.Read(d1.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if d1 != d2 {
	// 	log.Fatalf("Got %+v want %+v", d2, d1)
	// }

	log.Printf("Stopping app...\n")
}
