package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/router"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"log"
	"net/http"
)

func main() {
	redisConnection.InitPool()

	router := router.InitializeRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
