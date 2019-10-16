// publisher
package main

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
)


func main() {
	s1 := sensors.NewSensorWind(1,"ID","Wind", 1.0,"2019")

	//msg := "HELLO SUB"
	msg, err := json.Marshal(s1)
    if err != nil {
        fmt.Println(err)
        return
	}
	fmt.Println(msg)
	topic := "test"

	client_pub := connect("tcp://localhost:1883", "clientPub")
	client_pub.Publish(topic, 0, false, msg).Wait()

	time.Sleep(100 * time.Second)
	client_pub.Disconnect(10000)
}

func createClientOptions(brokerURI string, clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername (user)
	// opts.SetPassword (password)
	opts.SetClientID(clientID)
	return opts
}

func connect(brokerURI string, clientID string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientID + ")...)")
	opts := createClientOptions(brokerURI, clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}


