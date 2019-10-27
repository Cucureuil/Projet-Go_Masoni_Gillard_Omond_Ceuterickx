// redisHandler
package redis

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func InsertNewEntryRedis(data sensors.Sensor) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Do("SADD", "idAirports", data.IdAirport)

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

func InsertNewEntryCSV(data sensors.Sensor) {
	dataString, nameFile := prepareSensorDataForCSV(data)
	f, errOpen := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	var err interface{}
	if errOpen != nil {
		f, err = os.Create(nameFile)
	}
	if err != nil {
		fmt.Println("fail")
		fmt.Println(err)
	}
	w := csv.NewWriter(f)
	w.Write(dataString)
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
	}

}

func prepareSensorDataForCSV(data sensors.Sensor) ([]string, string) {
	dateMesure := time.Unix(data.DateMeasure, 0)
	dateString := dateMesure.Format("01-02-2006")
	value := strconv.FormatFloat(data.Value, 'f', 2, 64)
	nameFile := data.IdAirport + "-" + dateString + "-" + data.TypeMeasure + ".csv"
	return []string{
		data.IdAirport,
		data.TypeMeasure,
		strconv.Itoa(data.IdSensor),
		dateString,
		value,
	}, nameFile
}
