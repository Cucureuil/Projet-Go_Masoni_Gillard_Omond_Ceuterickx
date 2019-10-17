package subscriber

import(
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/client"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/DataStructSub"
	"fmt"
	"encoding/json"
)

func GetAirportAClient() mqtt.Client {
	return client.Connect(config.GetAirportA())
}

func GetAirportBClient() mqtt.Client {
	return client.Connect(config.GetAirportB())
}

func Receive(c mqtt.Client, m mqtt.Message) {
	ascii := string(m.Payload())
	fmt.Println(ascii)

	var data DataStructSub.Data
	json.Unmarshal(m.Payload(), &data)
	fmt.Println(data)
}