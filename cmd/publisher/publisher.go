// publisher
package publisher

import (
	"fmt"
	"log"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)

	// opts.SetUsername(user)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func Connect(conf *config.Config) mqtt.Client {
	brokerURI := conf.BrokerAddress + ":" + conf.BrokerPort
	clientId := conf.ClientID

	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ") ...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func Publish(client mqtt.Client, topic string, msg interface{}) {
	client.Publish(topic, 0, false, msg)
}

func GetWindClient() mqtt.Client {
	return Connect(config.GetWind())
}

func GetPressureClient() mqtt.Client {
	return Connect(config.GetPress())
}

func GetTemperatureClient() mqtt.Client {
	return Connect(config.GetTemp())
}
