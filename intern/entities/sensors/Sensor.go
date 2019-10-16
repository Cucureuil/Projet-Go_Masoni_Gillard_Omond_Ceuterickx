package sensors

import(
    "fmt"
    "math/rand"
    "time"
    "math"
)
// attributs with UpperCase because of json 
type sensor struct {
    IdSensor int
	IdAirport  string // only 3 char
    TypeMeasure string // Temperature, Atmospheric pressure, Wind speed
    Value float64
    DateMeasure string // (timestamp : YYYY-MM-DD-hh-mm-ss)
}

func NewSensorWind(idSensor int,idAirport string,value float64, dateMeasure string ) *sensor {
    sensor := sensor{IdSensor: idSensor, IdAirport:idAirport,TypeMeasure:"Wind speed",Value:value, DateMeasure:dateMeasure  }
    return &sensor
}

func NewSensorTemperature(idSensor int,idAirport string,value float64, dateMeasure string ) *sensor {
    sensor := sensor{IdSensor: idSensor, IdAirport:idAirport,TypeMeasure:"Temperature",Value:value, DateMeasure:dateMeasure  }
    return &sensor
}

func NewSensorPressure(idSensor int,idAirport string,value float64, dateMeasure string ) *sensor {
    sensor := sensor{IdSensor: idSensor, IdAirport:idAirport,TypeMeasure:"Atmospheric pressure",Value:value, DateMeasure:dateMeasure  }
    return &sensor
}

func RandSensor() *sensor {
    i:=rand.Intn(30)
    date := getDate()

	fmt.Println(i)
	switch {
        case i<=10:
            wind := roundMeasureValue(randWindSpeed())
			return NewSensorWind(1,"ID", wind,date)
        case i<=20 && i >10:
            temp := roundMeasureValue(randTemperature())
			return NewSensorTemperature(2,"ID", temp,date)
        case i<=30 && i>20:
            pres := roundMeasureValue(randPressure())
			return NewSensorPressure(3,"ID", pres,date)
    }
    return nil
}

func roundMeasureValue(val float64) float64{
    return math.Round(val*100)/100
}

// km/h
func randWindSpeed() float64{
    return 0 + rand.Float64() * (130 - 0)
}

// ° C
func randTemperature() float64{
    return -30 + rand.Float64() * (50 - (-30))
}

// bar
func randPressure() float64{
    return 0 + rand.Float64() * (2000 - 0)
}

// YYYY-MM-DD-hh-mm-ss
func getDate() string{
    return time.Now().Format("2006-01-02-15-04-05")     
}