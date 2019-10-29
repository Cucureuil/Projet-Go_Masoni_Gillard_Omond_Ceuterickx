package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/subscriber/generic"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"

	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
)

func main() {
	redisConnection.InitPool()
	client := subscriber.GetAirportBClient()
	client.Subscribe(config.GetAirportB().ClientID, 0, subscriber.Receive)
	for {
	}
}
