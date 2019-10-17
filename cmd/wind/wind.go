// publisher
package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/publisher"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	topic := "test"
	windClient := publisher.GetWindClient()

	for {
		s := sensors.RandSensorWind()

		// Write sensor in JSON
		msg, err := json.Marshal(s)
		if err != nil {
			fmt.Println(err)
			return
		}

		publisher.Publish(windClient, topic, msg)

		time.Sleep(10 * time.Second)
	}
}
