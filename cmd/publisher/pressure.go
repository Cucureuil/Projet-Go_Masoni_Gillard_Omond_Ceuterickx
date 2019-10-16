// publisher
package publisher

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishPress(client mqtt.Client, topic string) {
	s := sensors.RandSensorPressure()

	// Write sensor in JSON
	msg, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)

	Publish(client, topic, msg)
}
