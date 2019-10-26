package models

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"

	"strconv"
	"strings"
	"time"

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

type Parameters struct {
	IdSensor  string
	DateStart string
	DateEnd   string
}

func GetAverageAirports() []AverageByAirport {
	conn := redisConnection.Get()
	idAirports, _ := redis.Strings(conn.Do("SMEMBERS", "idAirports"))
	conn.Close()

	var averagesAirports []AverageByAirport
	for _, id := range idAirports {
		avA := GetAverageAirport(id, Parameters{})
		averagesAirports = append(averagesAirports, avA)
	}

	return averagesAirports
}

func GetAverageAirport(idAirport string, params Parameters) AverageByAirport {
	var avAirport AverageByAirport
	var avTypes []AverageByType

	a := GetAverageByType(idAirport, "Wind speed", params)
	avTypes = append(avTypes, a)

	a = GetAverageByType(idAirport, "Temperature", params)
	avTypes = append(avTypes, a)

	a = GetAverageByType(idAirport, "Atmospheric pressure", params)
	avTypes = append(avTypes, a)

	avAirport.IdAirport = idAirport
	avAirport.Averages = avTypes

	return avAirport
}

func GetAverageByType(idAirport string, dataType string, params Parameters) AverageByType {
	conn := redisConnection.Get()
	defer conn.Close()

	var a AverageByType
	key := idAirport + ":" + dataType
	var r []string
	if params.DateStart != "" {
		layout := "2006-01-02-15-04-05"
		startTime, _ := time.Parse(layout, params.DateStart)
		start := startTime.Unix()
		var end string
		end = "+inf"
		if params.DateEnd != "" {
			endTime, _ := time.Parse(layout, params.DateEnd)
			end = strconv.FormatInt(endTime.Unix(), 10)
		}
		r, _ = redis.Strings(conn.Do("ZRANGEBYSCORE", key, start, end))
	} else {
		r, _ = redis.Strings(conn.Do("ZRANGE", key, 0, -1))
	}
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
		dA := GetDataAirport(id, Parameters{})
		dataAirports = append(dataAirports, dA)
	}

	return dataAirports
}

func GetDataAirport(idAirport string, params Parameters) DataByAirport {
	var dAirport DataByAirport
	var dTypes []DataByType

	a := GetDataByType(idAirport, "Wind speed", params)
	dTypes = append(dTypes, a)

	a = GetDataByType(idAirport, "Temperature", params)
	dTypes = append(dTypes, a)

	a = GetDataByType(idAirport, "Atmospheric pressure", params)
	dTypes = append(dTypes, a)

	dAirport.IdAirport = idAirport
	dAirport.Data = dTypes

	return dAirport
}

func GetDataByType(idAirport string, dataType string, params Parameters) DataByType {
	key := idAirport + ":" + dataType
	r := getDataStrings(key, params)
	dataList := GetSensorListFromData(idAirport, dataType, Parameters{}, r)

	var d DataByType
	d.Data = dataList
	d.Type = dataType

	return d
}

func getDataStrings(key string, params Parameters) []string {
	conn := redisConnection.Get()
	defer conn.Close()

	var r []string
	if params.DateStart != "" {
		layout := "2006-01-02-15-04-05"
		startTime, _ := time.Parse(layout, params.DateStart)
		start := startTime.Unix()
		var end string
		end = "+inf"
		if params.DateEnd != "" {
			endTime, _ := time.Parse(layout, params.DateEnd)
			end = strconv.FormatInt(endTime.Unix(), 10)
		}
		r, _ = redis.Strings(conn.Do("ZRANGEBYSCORE", key, start, end, "WITHSCORES"))
	} else {
		r, _ = redis.Strings(conn.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	}

	return r
}

func GetSensorListFromData(idAirport string, dataType string, params Parameters, data []string) []sensors.Sensor {
	var sList []sensors.Sensor
	var s sensors.Sensor
	for i := 0; i < len(data); i += 2 {
		idSensor, val := ExtractData(data[i])
		idSensorParam, _ := strconv.Atoi(params.IdSensor)
		if params.IdSensor == "" || idSensorParam == idSensor {
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
	}

	return sList
}

func ExtractData(redisData string) (int, float64) {
	r := strings.Split(redisData, ":")
	idSensor, _ := strconv.Atoi(r[1])
	val, _ := strconv.ParseFloat(r[2], 64)
	return idSensor, val
}

func GetSensorData(params Parameters) []sensors.Sensor {
	conn := redisConnection.Get()
	idAirports, _ := redis.Strings(conn.Do("SMEMBERS", "idAirports"))
	conn.Close()

	var res []sensors.Sensor
	var found bool
	for _, id := range idAirports {
		res, found = GetSensorDataByType(id, "Atmospheric pressure", params)
		if found {
			break
		}
		res, found = GetSensorDataByType(id, "Wind speed", params)
		if found {
			break
		}
		res, found = GetSensorDataByType(id, "Temperature", params)
		if found {
			break
		}
	}

	return res
}

func GetSensorDataByType(idAirport string, dataType string, params Parameters) ([]sensors.Sensor, bool) {
	conn := redisConnection.Get()
	defer conn.Close()

	key := idAirport + ":" + dataType
	r := getDataStrings(key, params)
	var res []sensors.Sensor
	var found bool
	res = GetSensorListFromData(idAirport, dataType, params, r)
	found = len(res) > 0

	return res, found
}
