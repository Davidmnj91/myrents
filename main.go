package main

import (
	"log"
	"github.com/Davidmnj91/myrentals/db"
	"github.com/Davidmnj91/myrentals/api"
)

func main() {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err.Error);
		panic(err)
	}

	api.StartRouter()
}