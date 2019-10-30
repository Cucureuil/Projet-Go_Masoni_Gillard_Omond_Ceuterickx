package sensors

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/config"
	"math"
	"math/rand"
	"time"
)

// attributs with UpperCase because of json
type Sensor struct {
	IdSensor    int
	IdAirport   string // only 3 char
	TypeMeasure string // Temperature, Atmospheric pressure, Wind speed
	Value       float64
	DateMeasure int64 // (timestamp)
}

/**
* Constructors
 */
func NewSensorWind(idSensor int, idAirport string, value float64, dateMeasure int64) *Sensor {
	sensor := Sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Wind speed", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

func NewSensorTemperature(idSensor int, idAirport string, value float64, dateMeasure int64) *Sensor {
	sensor := Sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Temperature", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

func NewSensorPressure(idSensor int, idAirport string, value float64, dateMeasure int64) *Sensor {
	sensor := Sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Atmospheric pressure", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

/**
* Create sensor with random values
 */
func RandSensorWind() *Sensor {
	date := getDate()
	airport := RandIdAirport()
	wind := roundMeasureValue(randWindSpeed())
	return NewSensorWind(1, airport, wind, date)
}

func RandSensorTemperature() *Sensor {
	date := getDate()
	airport := RandIdAirport()
	temp := roundMeasureValue(randTemperature())
	return NewSensorTemperature(2, airport, temp, date)
}

func RandSensorPressure() *Sensor {
	date := getDate()
	airport := config.GetAirportB().ClientID
	pres := roundMeasureValue(randPressure())
	return NewSensorPressure(3, airport, pres, date)
}

/**
* Random airport id
 */
func RandIdAirport() string {
	i := rand.Intn(30)
	switch {
	case i <= 15:
		return config.GetAirportA().ClientID
	case i <= 30 && i > 15:
		return config.GetAirportB().ClientID
	}
	return ""
}

/**
* Round measure value (float64)
 */
func roundMeasureValue(val float64) float64 {
	return math.Round(val*100) / 100
}

/**
* Rand wind speed (km/h)
 */
func randWindSpeed() float64 {
	return 0 + rand.Float64()*(130-0)
}

/**
* Rand temperature (° C)
 */
func randTemperature() float64 {
	return -30 + rand.Float64()*(50-(-30))
}

/**
* Rand pressure (bar)
 */
func randPressure() float64 {
	return 0 + rand.Float64()*(2000-0)
}

/**
* get current date in format : YYYY-MM-DD-hh-mm-ss
 */
func getDate() int64 {
	return time.Now().Unix()
}
