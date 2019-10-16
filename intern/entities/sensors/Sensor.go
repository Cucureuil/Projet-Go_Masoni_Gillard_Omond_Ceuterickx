package sensors

import(
    "fmt"
    "math/rand"
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
	fmt.Println(i)
	switch {
		case i<=10:
			return NewSensorWind(1,"ID", 1.0,"2019")
		case i<=20 && i >10:
			return NewSensorTemperature(2,"ID", 1.0,"2019")
		case i<=30 && i>20:
			return NewSensorPressure(3,"ID", 1.0,"2019")
    }
    return nil
}