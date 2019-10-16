// Main
package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/publisher"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"time"
)

func main() {
	topic := "test"

	windClient := publisher.Connect(config.GetWind())
	pressClient := publisher.Connect(config.GetPress())
	tempClient := publisher.Connect(config.GetTemp())

	for {
		publisher.PublishPress(pressClient, topic)
		publisher.PublishTemp(tempClient, topic)
		publisher.PublishWind(windClient, topic)

		time.Sleep(10 * time.Second)
	}
}
