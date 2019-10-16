package sensors

import (
	"math"
	"math/rand"
	"time"
)

// attributs with UpperCase because of json
type sensor struct {
	IdSensor    int
	IdAirport   string // only 3 char
	TypeMeasure string // Temperature, Atmospheric pressure, Wind speed
	Value       float64
	DateMeasure string // (timestamp : YYYY-MM-DD-hh-mm-ss)
}

/**
* Constructors
 */
func NewSensorWind(idSensor int, idAirport string, value float64, dateMeasure string) *sensor {
	sensor := sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Wind speed", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

func NewSensorTemperature(idSensor int, idAirport string, value float64, dateMeasure string) *sensor {
	sensor := sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Temperature", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

func NewSensorPressure(idSensor int, idAirport string, value float64, dateMeasure string) *sensor {
	sensor := sensor{IdSensor: idSensor, IdAirport: idAirport, TypeMeasure: "Atmospheric pressure", Value: value, DateMeasure: dateMeasure}
	return &sensor
}

/**
* Create sensor with random values
 */
func RandSensorWind() *sensor {
	date := getDate()
	airport := RandIdAirport()
	wind := roundMeasureValue(randWindSpeed())
	return NewSensorWind(1, airport, wind, date)
}

func RandSensorTemperature() *sensor {
	date := getDate()
	airport := RandIdAirport()
	temp := roundMeasureValue(randTemperature())
	return NewSensorTemperature(2, airport, temp, date)
}

func RandSensorPressure() *sensor {
	date := getDate()
	airport := RandIdAirport()
	pres := roundMeasureValue(randPressure())
	return NewSensorPressure(3, airport, pres, date)
}

/**
* Random airport id
 */
func RandIdAirport() string {
	i := rand.Intn(30)
	switch {
	case i <= 10:
		return "AAA1"
	case i <= 20 && i > 10:
		return "BBB1"
	case i <= 30 && i > 20:
		return "CCC1"
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
func getDate() string {
	return time.Now().Format("2006-01-02-15-04-05")
}
