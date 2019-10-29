// redisHandler
package redis

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Insert the informations of a sensor into redis database
func InsertNewEntryRedis(data sensors.Sensor) {
	conn := redisConnection.Get()
	defer conn.Close()

	conn.Do("SADD", "idAirports", data.IdAirport)

	key := data.IdAirport + ":" + data.TypeMeasure
	timestamp := data.DateMeasure
	idDonnee, _ := conn.Do("INCR", "idDonnee")
	idDonneeString := strconv.FormatInt(idDonnee.(int64), 10)
	idSensorString := strconv.Itoa(data.IdSensor)
	valueString := strconv.FormatFloat(data.Value, 'f', 2, 64)
	value := idDonneeString + ":" + idSensorString + ":" + valueString

	_, err := conn.Do("ZADD", key, timestamp, value)
	if err != nil {
		log.Fatal(err)
	}
}

// Insert a new line in a CSV file
// The CSV file is opened or create in the same folder as the application
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

// Return a tuple with the informations from a sensor as a list of string and
// the title of the file to put the informations in
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
