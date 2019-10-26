// test
package main

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/models"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/redisConnection"
	"fmt"
)

func main() {
	// r := models.Average([]string{"1:20:20.32", "1:20:40.32", "1:20:948.5"})
	// fmt.Println(r)

	redisConnection.InitPool()

	// r := models.GetAverageAirport()
	// for _, item := range r {
	// 	fmt.Println(item)
	// }
	fmt.Println(1 in []int{1, 2})
	r := models.GetDataAirports()
	fmt.Println(r)
	for _, item := range r {
		fmt.Println(item)
	}
}
