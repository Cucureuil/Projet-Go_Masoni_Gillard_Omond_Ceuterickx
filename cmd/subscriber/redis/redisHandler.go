// redisHandler
package redis

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func InsertNewEntry(data sensors.Sensor) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer conn.Close()

	key := data.IdAirport + ":" + data.TypeMeasure
	timestamp := data.DateMeasure
	idDonnee, _ := conn.Do("INCR", "idDonnee")
	idDonneeString := strconv.FormatInt(idDonnee.(int64), 10)
	idSensorString := strconv.Itoa(data.IdSensor)
	valueString := strconv.FormatFloat(data.Value, 'f', 2, 64)
	value := idDonneeString + ":" + idSensorString + ":" + valueString

	_, err = conn.Do("ZADD", key, timestamp, value)
	if err != nil {
		log.Fatal(err)
	}
}
