// publisher
package publisher

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishTemp(client mqtt.Client, topic string) {
	s := sensors.RandSensorTemperature()

	// Write sensor in JSON
	msg, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)

	Publish(client, topic, msg)
}
