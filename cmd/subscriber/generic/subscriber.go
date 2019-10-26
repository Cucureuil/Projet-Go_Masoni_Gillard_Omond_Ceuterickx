package subscriber

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/client"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/subscriber/redis"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetAirportAClient() mqtt.Client {
	return client.Connect(config.GetAirportA())
}

func GetAirportBClient() mqtt.Client {
	return client.Connect(config.GetAirportB())
}

func Receive(c mqtt.Client, m mqtt.Message) {
	var data sensors.Sensor
	json.Unmarshal(m.Payload(), &data)
	redis.InsertNewEntry(data)
}
