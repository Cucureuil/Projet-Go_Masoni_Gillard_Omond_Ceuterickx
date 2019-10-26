package models

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
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
			r := strings.Split(d, ":")
			val, _ := strconv.ParseFloat(r[len(r)-1], 64)
			av = av + val
		}

		return av / float64(len(data))
	}
}
