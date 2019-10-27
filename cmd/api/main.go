package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/router"

	"github.com/gorilla/handlers"

	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"log"
	"net/http"
)

func main() {
	redisConnection.InitPool()

	router := router.InitializeRouter()
	// log.Fatal(http.ListenAndServe(":8081", router))
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
