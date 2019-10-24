package router

import (
	"github.com/gorilla/mux"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/controller"
  )
  
  func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)
  
	// list of all airport
	router.Methods("GET").Path("/airports").Name("Airports").HandlerFunc(controller.ListAirports)
	// data for one airport
	router.Methods("GET").Path("/airport/{id}").Name("Airport").HandlerFunc(controller.GetAirportData)
	// data for one sensor
	router.Methods("GET").Path("/sensor/{id}").Name("Sensor").HandlerFunc(controller.GetSensorData)
	// data for one date (YYYY-MM-DD-hh-mm-ss)
	router.Methods("GET").Path("/date/{date}").Name("Date").HandlerFunc(controller.GetDataByDate)
	// data between two dates (YYYY-MM-DD-hh-mm-ss)
	router.Methods("GET").Path("/dates/{startDate}/{endDate}").Name("Dates").HandlerFunc(controller.GetDataBetweenTwoDates)
	// averages for each types of sensor
	router.Methods("GET").Path("/average").Name("average").HandlerFunc(controller.GetSensorsAverage)
	
	return router
  }