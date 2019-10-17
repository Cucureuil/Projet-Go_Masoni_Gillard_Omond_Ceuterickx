// RedisRegister
package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/DataStructSub"
	"encoding/json"
	"log"
	"time"
)

func main() {
	client := connect("tcp://localhost:1883", "register")
	client.Subscribe("register", 0, myfunction)
	for {
	}
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
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
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

func myfunction(c mqtt.Client, m mqtt.Message) {
	ascii := string(m.Payload())
	fmt.Println(ascii)

	var data DataStructSub.Data
	json.Unmarshal(m.Payload(), &data)
	fmt.Println(data)
}

