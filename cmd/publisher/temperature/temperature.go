// publisher
package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/publisher/generic"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	tempClient := publisher.GetTemperatureClient()

	for {
		s := sensors.RandSensorTemperature()
		//topic := s.IdAirport + "-TEMP"

		topic := s.IdAirport
		// Write sensor in JSON
		msg, err := json.Marshal(s)
		if err != nil {
			fmt.Println(err)
			return
		}

		publisher.Publish(tempClient, topic, msg)

		time.Sleep(10 * time.Second)
	}
}