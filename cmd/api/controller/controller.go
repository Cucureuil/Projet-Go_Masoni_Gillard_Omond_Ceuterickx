package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"strconv"
	"fmt"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/intern/entities/sensors"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/models"
)

/*
* Call in router:/airports
* List all Airports in DB
*/
func ListAirports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	airports := models.ListAllAirport() 
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

  	id, err := strconv.Atoi(vars["id"])

  	if err != nil {
		log.Fatal(err)  
	}
	
	fmt.Println(id)
}

/*
* Call in router:/sensor/{id}
* Get data for one sensor by its ID
*/
func GetSensorData (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// get params like : www.../airport/id
	vars := mux.Vars(r)

  	id, err := strconv.Atoi(vars["id"])

  	if err != nil {
		log.Fatal(err)  
	}
	
	fmt.Println(id)

	json.NewEncoder(w).Encode(sensors.RandSensorWind())
}

/**
* Call in router:/sensor/{id}
* Get data for one date (YYYY-MM-DD-hh-mm-ss)
*/
func GetDataByDate (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	  
	vars := mux.Vars(r)

  	date := vars["date"]
	
	fmt.Println(date)
}

/*
* Call in router:/sensor/{id}
* Get data between two dates (YYYY-MM-DD-hh-mm-ss)
*/
func GetDataBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	dateStart := vars["dateStart"]
	dateEnd := vars["dateEnd"]
  
	fmt.Println(dateStart)
	fmt.Println(dateEnd)
}

/*
* Call in router:/average
* Get averages for each types of sensor
*/
func GetSensorsAverage (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}