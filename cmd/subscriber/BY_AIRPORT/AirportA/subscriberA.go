package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/subscriber/generic"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
)

func main() {
	redisConnection.InitPool()
	client := subscriber.GetAirportAClient()
	client.Subscribe(config.GetAirportA().ClientID, 0, subscriber.Receive)
	for {
	}
}
