package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/models"
)

/*
* Call in router:/airports
* List all Airports in DB
*/
func ListAirports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// fakes data
	airports := models.ListAllAirport()
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

	// fakes data
	airports := models.ListAllAirport()
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

  	id:= vars["id"]
	sensorType := vars["sensorType"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AIRPORT BY ID & SENSOR TYPE")
	fmt.Println(id,sensorType)
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

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AIRPORT BY ID & TWO DATES")
	fmt.Println(id,startDate,endDate)
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
	airports := models.ListAllAirport()
	fmt.Println("AIRPORT BY ID & TWO DATES & SENSOR TYPE")
	fmt.Println(id,startDate,endDate,sensorType)
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

  	id, err :=  strconv.Atoi(vars["id"])

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("SENSOR BY ID")
	fmt.Println(id,err)
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

  	id, err := strconv.Atoi(vars["id"])
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("SENSOR BY ID & TWO DATES")
	fmt.Println(id,err,startDate,endDate)
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
	startDate := vars["startDate"]
	endDate := vars["endDate"]
	sensorType := vars["sensorType"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO AIRPORT BY ID & SENSOR TYPE")
	fmt.Println(id,startDate,endDate,sensorType)
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

  	id, err := strconv.Atoi(vars["id"])

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO SENSOR")
	fmt.Println(id,err)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/type/{sensorType}
* Get  averages for each types of sensor by Sensor Type
*/
func GetSensorsAverageBySensorType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	// get params like : www.../airport/id
	vars := mux.Vars(r)

	sensorType := vars["sensorType"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO SENSOR TYPE")
	fmt.Println(sensorType)
	json.NewEncoder(w).Encode(airports)
}

/*
* Call in router:/average/type/{sensorType}/{startDate}/{endDate}
* Get averages for each types of sensor between two dates
*/
func GetSensorsAverageBySensorTypeBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	// get params like : www.../airport/id
	vars := mux.Vars(r)

	sensorType := vars["sensorType"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO SENSOR TYPE && TWO DATES")
	fmt.Println(sensorType,startDate,endDate)
	json.NewEncoder(w).Encode(airports)
}

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
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO SENSOR TYPE && TWO DATES BY AIRPORT")
	fmt.Println(id,startDate,endDate,sensorType)
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

	id, err := strconv.Atoi(vars["id"])
	sensorType := vars["sensorType"]
	startDate := vars["startDate"]
	endDate := vars["endDate"]

	// fakes data
	airports := models.ListAllAirport()
	fmt.Println("AVERAGE TO SENSOR BY ID && TWO DATES")
	fmt.Println(id,err,startDate,endDate,sensorType)
	json.NewEncoder(w).Encode(airports)
}