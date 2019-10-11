package sensors

type sensor struct {
    idSensor int
	idAirport  string // only 3 char
    typeMeasure string // Temperature, Atmospheric pressure, Wind speed
    value float64
    dateMeasure string // (timestamp : YYYY-MM-DD-hh-mm-ss)
}

func NewSensorWind(idSensor int,idAirport string,typeMeasure string,value float64, dateMeasure string ) *sensor {
    sensor := sensor{idSensor: idSensor, idAirport:idAirport,typeMeasure:typeMeasure,value:value, dateMeasure:dateMeasure  }
    return &sensor
}