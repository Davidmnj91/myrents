package main

import (
	"os"
	"log"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
	
	"github.com/Davidmnj91/myrentals/useraccount"
	"github.com/Davidmnj91/myrentals/db"
	"github.com/Davidmnj91/myrentals/api"
)

func main() {
	mongo, err := db.Connect()
	defer db.Disconnect()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	router := api.StartRouter()
	setupUserAccountRoutes(router, mongo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	http.ListenAndServe(":"+PORT, router)
}

func setupUserAccountRoutes(router *mux.Router, db *mongo.Database) {
	userHandler := useraccount.NewUserAccountHandler(db)
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
}