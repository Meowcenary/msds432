package main

import (
	"encoding/json"
	"log"
	"net/http"

	"api/internal/dbconnector"
	"api/internal/models"

	"github.com/gorilla/mux"
)

func main() {
  // Initialize the database connection pool
  if err := dbconnector.InitDB(); err != nil {
    log.Fatalf("Could not initialize database: %v", err)
  }
  defer dbconnector.CloseDB() // Ensure that the pool is closed when the program exits

	// Router and routes
	router := mux.NewRouter()
	// router.HandleFunc("/", GetRoot).Methods("GET")
	// router.HandleFunc("/taxi_trips", GetRoot).Methods("GET")
	router.HandleFunc("/covid_19_reports", GetCovid19Reports).Methods("GET")

	// Start the server
	http.Handle("/", router)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetCovid19Reports(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting covid 19 reports")
  var reports []models.Covid19Report
	var err error
	reports, err = dbconnector.GetData[models.Covid19Report]("Covid19Reports")
  if err != nil {
  	panic(err)
  }

	for _, report := range reports {
  	log.Println("Report: ", report)
  }

	// encode JSON for response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reports)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
	}
}
