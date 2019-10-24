package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/router"
	"log"
	"net/http"
)

func main() {
  router := router.InitializeRouter()
  log.Fatal(http.ListenAndServe(":80", router))
}