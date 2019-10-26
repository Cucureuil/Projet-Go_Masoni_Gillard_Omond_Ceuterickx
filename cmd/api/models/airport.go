package models

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type AverageByAirport struct {
	IdAirport string
	Averages  []AverageByType
}

type AverageByType struct {
	Type    string
	Average float64
}

type DataByAirport struct {
	IdAirport string
	Data      []DataByType
}

type DataByType struct {
	Type string
	Data []sensors.Sensor
}

func GetAverageAirports() []AverageByAirport {
	conn := redisConnection.Get()
	idAirports, _ := redis.Strings(conn.Do("SMEMBERS", "idAirports"))
	conn.Close()

	var averagesAirports []AverageByAirport
	for _, id := range idAirports {
		avA := GetAverageAirport(id)
		averagesAirports = append(averagesAirports, avA)
	}

	return averagesAirports
}

func GetAverageAirport(idAirport string) AverageByAirport {
	conn := redisConnection.Get()
	defer conn.Close()

	var avAirport AverageByAirport
	var avTypes []AverageByType

	a := GetAverageByType(conn, idAirport, "Wind speed")
	avTypes = append(avTypes, a)

	a = GetAverageByType(conn, idAirport, "Temperature")
	avTypes = append(avTypes, a)

	a = GetAverageByType(conn, idAirport, "Atmospheric pressure")
	avTypes = append(avTypes, a)

	avAirport.IdAirport = idAirport
	avAirport.Averages = avTypes

	return avAirport
}

func GetAverageByType(conn redis.Conn, idAirport string, dataType string) AverageByType {
	var a AverageByType
	key := idAirport + ":" + dataType
	r, _ := redis.Strings(conn.Do("ZRANGE", key, 0, -1))
	av := Average(r)
	a.Average = av
	a.Type = dataType

	return a
}

func Average(data []string) float64 {
	var av float64
	av = 0
	if len(data) == 0 {
		return av
	} else {
		for _, d := range data {
			_, val := ExtractData(d)
			av = av + val
		}

		return av / float64(len(data))
	}
}

func GetDataAirports() []DataByAirport {
	conn := redisConnection.Get()
	idAirports, _ := redis.Strings(conn.Do("SMEMBERS", "idAirports"))
	conn.Close()

	var dataAirports []DataByAirport
	for _, id := range idAirports {
		dA := GetDataAirport(id)
		dataAirports = append(dataAirports, dA)
	}

	return dataAirports
}

func GetDataAirport(idAirport string) DataByAirport {
	conn := redisConnection.Get()
	defer conn.Close()

	var dAirport DataByAirport
	var dTypes []DataByType

	a := GetDataByType(conn, idAirport, "Wind speed")
	dTypes = append(dTypes, a)

	a = GetDataByType(conn, idAirport, "Temperature")
	dTypes = append(dTypes, a)

	a = GetDataByType(conn, idAirport, "Atmospheric pressure")
	dTypes = append(dTypes, a)

	dAirport.IdAirport = idAirport
	dAirport.Data = dTypes

	return dAirport
}

func GetDataByType(conn redis.Conn, idAirport string, dataType string) DataByType {
	var d DataByType
	key := idAirport + ":" + dataType
	r, _ := redis.Strings(conn.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	dataList := GetSensorListFromData(idAirport, dataType, r)
	d.Data = dataList
	d.Type = dataType

	return d
}

func GetSensorListFromData(idAirport string, dataType string, data []string) []sensors.Sensor {
	var sList []sensors.Sensor
	var s sensors.Sensor
	for i := 0; i < len(data); i += 2 {
		idSensor, val := ExtractData(data[i])
		timestamp, _ := strconv.ParseInt(data[i+1], 10, 64)

		s = sensors.Sensor{
			IdSensor:    idSensor,
			IdAirport:   idAirport,
			TypeMeasure: dataType,
			Value:       val,
			DateMeasure: timestamp,
		}
		sList = append(sList, s)
	}

	return sList
}

func ExtractData(redisData string) (int, float64) {
	r := strings.Split(redisData, ":")
	idSensor, _ := strconv.Atoi(r[1])
	val, _ := strconv.ParseFloat(r[2], 64)
	return idSensor, val
}

func IsIn(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
