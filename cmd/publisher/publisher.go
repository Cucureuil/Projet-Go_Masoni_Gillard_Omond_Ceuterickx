// publisher
package main

import (
	"fmt"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
)

func main() {
	s1 := sensors.NewSensorWind(1,"ID","Wind", 1.0,"2019")
	fmt.Println(s1)
}
