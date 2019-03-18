package main

import (
	"fmt"
	"log"

	"github.com/ggerritsen/building-blocks-app/repository"
)

func main() {
	println(fmt.Sprintf("Hello, %s", "hoi"))

	repo, err := repository.NewRepositoryWithDb("localhost", "postgres", "your-password", "template1", 5432)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	fmt.Printf("Connected to database\n")
}
