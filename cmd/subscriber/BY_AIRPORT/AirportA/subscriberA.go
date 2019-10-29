package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/subscriber/generic"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
)
func main() {
	client := subscriber.GetAirportAClient()
	client.Subscribe(config.GetAirportA().ClientID, 0, subscriber.Receive)
	for {
	}
}
