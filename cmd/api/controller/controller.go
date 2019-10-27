package controller

import (
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
* Call in router:/airports
* List all Airports in DB
 */
func ListAirports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	airports := models.GetAverageAirports()
	fmt.Println("AIRPORTS")

	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/airport/{id}
* Get data for one airport by its ID
 */
func GetAirportData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]

	airports := models.GetDataAirport(id, models.Parameters{})
	fmt.Println("AIRPORT BY ID")
	fmt.Println(id)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/airport/{id}/{sensorType}
* Get data for one airport by sensor type
 */
func GetAirportDataBySensorType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	sensorType := vars["sensorType"]

	airports := models.GetDataByType(id, sensorType, models.Parameters{})
	fmt.Println("AIRPORT BY ID & SENSOR TYPE")
	fmt.Println(id, sensorType)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/airport/{id}/{startDate}/{endDate}
* Get data for one airport between two dates (YYYY-MM-DD-hh-mm-ss)
 */
func GetAirportDataBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	airports := models.GetDataAirport(id, models.Parameters{DateStart: startDate, DateEnd: endDate})
	fmt.Println("AIRPORT BY ID & TWO DATES")
	fmt.Println(id, startDate, endDate)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/airport/{id}/{startDate}/{endDate}/{sensorType}
* Get data for one airport between two dates and by sensor type (YYYY-MM-DD-hh-mm-ss)
 */
func GetAirportDataBetweenTwoDatesAndSensorType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]
	sensorType := vars["sensorType"]

	// fakes data
	airports := models.GetDataByType(id, sensorType, models.Parameters{DateStart: startDate, DateEnd: endDate})
	fmt.Println("AIRPORT BY ID & TWO DATES & SENSOR TYPE")
	fmt.Println(id, startDate, endDate, sensorType)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/sensor/{id}
* Get data for one airport between two dates and by sensor type (YYYY-MM-DD-hh-mm-ss)
 */
func GetSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]

	// fakes data
	airports := models.GetSensorData(models.Parameters{IdSensor: id})
	fmt.Println("SENSOR BY ID")
	fmt.Println(id)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/sensor/{id}/{startDate}/{endDate}
* Get data for one sensor between two dates (YYYY-MM-DD-hh-mm-ss)
 */
func GetSensorDataBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	// fakes data
	airports := models.GetSensorData(models.Parameters{IdSensor: id, DateStart: startDate, DateEnd: endDate})
	fmt.Println("SENSOR BY ID & TWO DATES")
	fmt.Println(id, startDate, endDate)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/airport/{id}/{sensorType}
* Get averages for each types of sensor by Airport
 */
func GetSensorsAverageBySensorTypeAndAirport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	sensorType := vars["sensorType"]

	airports := models.GetAverageByType(id, sensorType, models.Parameters{})
	fmt.Println("AVERAGE TO AIRPORT BY ID & SENSOR TYPE")
	fmt.Println(id, sensorType)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/sensor/{id}
* Get average sensor
 */
func GetSensorsAverageBySensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]

	data := models.GetSensorData(models.Parameters{IdSensor: id})
	airports := models.AverageFromSensorData(data)

	fmt.Println("AVERAGE TO SENSOR")
	fmt.Println(id)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/sensor/{id}/{startDate}/{endDate}
* Get averages for each types of sensor by Sensor Type, airport Id and between two dates
 */
func GetSensorsAverageBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	data := models.GetSensorData(models.Parameters{IdSensor: id, DateStart: startDate, DateEnd: endDate})
	airports := models.AverageFromSensorData(data)

	fmt.Println("AVERAGE TO SENSOR BY ID && TWO DATES")
	fmt.Println(id, startDate, endDate)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/type/{sensorType}
* Get  averages for each types of sensor by Sensor Type
 */
// func GetSensorsAverageBySensorType(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json;charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	// get params like : www.../airport/id
// 	vars := mux.Vars(r)

// 	sensorType := vars["sensorType"]

// 	// fakes data
// 	airports := models.ListAllAirport()
// 	fmt.Println("AVERAGE TO SENSOR TYPE")
// 	fmt.Println(sensorType)
// 	json.NewEncoder(w).Encode(airports)
// }

/*
* Call in router:/average/type/{sensorType}/{startDate}/{endDate}
* Get averages for each types of sensor between two dates
 */
// func GetSensorsAverageBySensorTypeBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json;charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	// get params like : www.../airport/id
// 	vars := mux.Vars(r)

// 	sensorType := vars["sensorType"]
// 	startDate := vars["startDate"]
// 	endDate := vars["endDate"]

// 	// fakes data
// 	airports := models.Get
// 	fmt.Println("AVERAGE TO SENSOR TYPE && TWO DATES")
// 	fmt.Println(sensorType, startDate, endDate)
// 	json.NewEncoder(w).Encode(airports)
// }

/*
* Call in router:/average/airport/{id}/{sensorType}/{startDate}/{endDate}
* Get averages for each types of sensor by Sensor Type, airport Id and between two dates
 */
func GetSensorsAverageBySensorAndAirportAndTypeBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

	id := vars["id"]
	sensorType := vars["sensorType"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	// fakes data
	airports := models.GetAverageByType(id, sensorType, models.Parameters{DateStart: startDate, DateEnd: endDate})
	fmt.Println("AVERAGE TO SENSOR TYPE && TWO DATES BY AIRPORT")
	fmt.Println(id, startDate, endDate, sensorType)
	json.NewEncoder(w).Encode(airports)
}
