package router

import (
	"github.com/gorilla/mux"
	"Projet-Go_Masoni_Gillard_Omond_Ceuterickx/cmd/api/controller"
  )
  
  func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)
  
	// R1: list of all airport + averages for each types of sensor 
	router.Methods("GET").Path("/airports").Name("Airports").HandlerFunc(controller.ListAirports)

	// R2: data for one airport
	router.Methods("GET").Path("/airport/{id}").Name("Airport").HandlerFunc(controller.GetAirportData)
	// R2: data for one airport by sensor type
	router.Methods("GET").Path("/airport/{id}/{sensorType}").Name("Airport").HandlerFunc(controller.GetAirportDataBySensorType)
	// R2: data for one airport between two dates (YYYY-MM-DD-hh-mm-ss)
	router.Methods("GET").Path("/airport/{id}/{startDate}/{endDate}").Name("Airport").HandlerFunc(controller.GetAirportDataBetweenTwoDates)
	// R2: data for one airport between two dates and by sensor type
	router.Methods("GET").Path("/airport/{id}/{startDate}/{endDate}/{sensorType}").Name("Airport").HandlerFunc(controller.GetAirportDataBetweenTwoDatesAndSensorType)

	// R3: data for one sensor
	router.Methods("GET").Path("/sensor/{id}").Name("Sensor").HandlerFunc(controller.GetSensorData)
	// R3: data for one sensor between two dates (YYYY-MM-DD-hh-mm-ss)
	router.Methods("GET").Path("/sensor/{id}/{startDate}/{endDate}").Name("Sensor").HandlerFunc(controller.GetSensorDataBetweenTwoDates)

	// R4: averages for each types of sensor by Airport
	router.Methods("GET").Path("/average/airport/{id}/{sensorType}").Name("average").HandlerFunc(controller.GetSensorsAverageBySensorTypeAndAirport)
	// R4: averages for each types of sensor by Sensor Type, airport Id and between two dates
	router.Methods("GET").Path("/average/airport/{id}/{sensorType}/{startDate}/{endDate}").Name("average").HandlerFunc(controller.GetSensorsAverageBySensorAndAirportAndTypeBetweenTwoDates)
	
	// R4: average for sensor 
	router.Methods("GET").Path("/average/sensor/{id}").Name("average").HandlerFunc(controller.GetSensorsAverageBySensor)
	// R4: average for sensor between two dates
	router.Methods("GET").Path("/average/sensor/{id}/{startDate}/{endDate}").Name("average").HandlerFunc(controller.GetSensorsAverageBetweenTwoDates)

	// R4: averages for each types of sensor by Sensor Type
	router.Methods("GET").Path("/average/type/{sensorType}").Name("average").HandlerFunc(controller.GetSensorsAverageBySensorType)
	// R4: averages for each types of sensor between two dates
	router.Methods("GET").Path("/average/type/{sensorType}/{startDate}/{endDate}").Name("average").HandlerFunc(controller.GetSensorsAverageBySensorTypeBetweenTwoDates)	
	
	return router
  }