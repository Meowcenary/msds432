package main

import (
	"encoding/json"
	"log"
	"net/http"

	"api/internal/dbconnector"
	"api/internal/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
  // Initialize the database connection pool
  if err := dbconnector.InitDB(); err != nil {
    log.Fatalf("Could not initialize database: %v", err)
  }
	// Ensure that the pool is closed when the program exits
  defer dbconnector.CloseDB()

	// Router and routes
	router := mux.NewRouter()
	router.HandleFunc("/taxi_trips", GetTaxiTrips).Methods("GET")
	router.HandleFunc("/transportation_network_trips", GetTransportationNetworkTrips).Methods("GET")
	router.HandleFunc("/building_permits", GetBuildingPermits).Methods("GET")
	router.HandleFunc("/chicago_ccvi", GetChicagoCCVI).Methods("GET")
	router.HandleFunc("/public_health_stats", GetPublicHealthStats).Methods("GET")
	router.HandleFunc("/covid_19_reports", GetCovid19Reports).Methods("GET")

	// Start the server
	// http.Handle("/", router)
	handler := cors.Default().Handler(router)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetTaxiTrips(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting taxi trips")
	var rows []models.TaxiTrip
	var err error
	rows, err = dbconnector.GetData[models.TaxiTrip]("TaxiTrips", "")
	if err != nil {
		panic(err)
	}
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rows)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetTransportationNetworkTrips(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting transportation network trips")
	var rows []models.TransportationNetworkProvidersTrip
	var err error
	rows, err = dbconnector.GetData[models.TransportationNetworkProvidersTrip]("TransportationNetworkProvidersTrips", "")
	if err != nil {
		panic(err)
	}
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rows)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetBuildingPermits(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting building permits")
	var rows []models.BuildingPermit
	var err error
	rows, err = dbconnector.GetData[models.BuildingPermit]("BuildingPermits", "1000")
	if err != nil {
		panic(err)
	}
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rows)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetChicagoCCVI(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting chicago ccvi")
	var rows []models.ChicagoCovid19CommunityVulnerabilityIndex
	var err error
	rows, err = dbconnector.GetData[models.ChicagoCovid19CommunityVulnerabilityIndex]("ChicagoCovid19CommunityVulnerabilityIndex", "100")
	if err != nil {
		panic(err)
	}
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rows)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetPublicHealthStats(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting public health stats")
	var rows []models.PublicHealthStatistic
	var err error
	rows, err = dbconnector.GetData[models.PublicHealthStatistic]("PublicHealthStatistics", "")
	if err != nil {
		panic(err)
	}
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rows)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetCovid19Reports(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting covid 19 reports")
  var reports []models.Covid19Report
	var err error
	reports, err = dbconnector.GetData[models.Covid19Report]("Covid19Reports", "")
  if err != nil {
  	panic(err)
  }
	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reports)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}
