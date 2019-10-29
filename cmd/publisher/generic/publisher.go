// publisher
package publisher

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/client"
)

func Publish(client mqtt.Client, topic string, msg interface{}) {
	client.Publish(topic, 0, false, msg)
}

func GetWindClient() mqtt.Client {
	return client.Connect(config.GetWind())
}

func GetPressureClient() mqtt.Client {
	return client.Connect(config.GetPress())
}

func GetTemperatureClient() mqtt.Client {
	return client.Connect(config.GetTemp())
}
