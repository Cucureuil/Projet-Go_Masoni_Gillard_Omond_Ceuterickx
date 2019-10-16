package sensors

// attributs with UpperCase because of json 
type sensor struct {
    IdSensor int
	IdAirport  string // only 3 char
    TypeMeasure string // Temperature, Atmospheric pressure, Wind speed
    Value float64
    DateMeasure string // (timestamp : YYYY-MM-DD-hh-mm-ss)
}

func NewSensorWind(idSensor int,idAirport string,typeMeasure string,value float64, dateMeasure string ) *sensor {
    sensor := sensor{IdSensor: idSensor, IdAirport:idAirport,TypeMeasure:typeMeasure,Value:value, DateMeasure:dateMeasure  }
    return &sensor
}