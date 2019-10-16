// publisher
package main

import (
	"fmt"
	"log"
	"time"

	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	Publish(config.GetPress(), "test", "Pressure = 111")
}

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

func connect(brokerURI string, clientId string) mqtt.Client {
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

func Publish(conf *config.Config, topic string, msg string) {
	client2 := connect(conf.BrokerAddress+":"+conf.BrokerPort, conf.ClientID)
	client2.Publish(topic, 0, false, msg)
	time.Sleep(5 * time.Second)
	fmt.Println("Message sended by " + conf.ClientID)
}
