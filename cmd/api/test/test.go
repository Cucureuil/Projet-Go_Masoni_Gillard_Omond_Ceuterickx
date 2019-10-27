// test
package main

import (
	// "Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/models"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/subscriber/redis"
	// "Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	// "fmt"
	// "github.com/gomodule/redigo/redis"
	// "encoding/csv"
	"time"
)

func main() {
	// r := models.Average([]string{"1:20:20.32", "1:20:40.32", "1:20:948.5"})
	// fmt.Println(r)

	// redisConnection.InitPool()

	// r := models.GetSensorData("1")
	// for _, item := range r {
	// 	fmt.Println(item)
	// }

	// conn := redisConnection.Get()
	// r, _ := redis.Values(conn.Do("ZSCAN", "BBB1:Atmospheric pressure", 0, "MATCH", "1:*:*"))
	// fmt.Println(r)
	// conn.Close()
	// r := models.GetSensorData(models.Parameters{IdSensor: "3"})
	// r := models.GetDataAirport("BBB1", models.Parameters{DateStart: "2019-10-23-18-33-33", DateEnd: "2019-10-23-18-34-59"})
	// for _, item := range r {
	// 	fmt.Println(item)
	// }
	layout := "2006-01-02-15-04-05"
	dateString, _ := time.Parse(layout, "2019-10-23-18-33-33")
	date := dateString.Unix()
	redis.InsertNewEntryCSV(sensors.Sensor{
		IdAirport:   "BBB1",
		IdSensor:    2,
		DateMeasure: date,
		TypeMeasure: "Temperature",
		Value:       123.23,
	})
}
