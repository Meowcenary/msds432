package main

import (
	"fmt"
	"log"

	"datapull/internal/dbconnector"
	"datapull/internal/pull"
)

type Dataset struct {
    Name      string
    TableName string
		Url       string
    SortField string
}

func main() {
	dataSets := []Dataset{
    // {
    // 	Name:      "Taxi Trips (2013-2023)",
    //   TableName: "TaxiTrips",
    //   Url:       "https://data.cityofchicago.org/resource/wrvz-psew",
    //   SortField: "trip_id",
    // },
    // {
    //   Name:      "Transportation Network Providers - Trips (2018 - 2022)",
    //   TableName: "TransportationNetworkProvidersTrips",
    //   Url:       "https://data.cityofchicago.org/resource/m6dm-c72p",
    //   SortField: "",
    // },
    // {
    //   Name:      "City of Chicago Building Permits",
    //   TableName: "BuildingPermits",
    //   Url:       "https://data.cityofchicago.org/resource/ydr8-5enu",
    //   SortField: "ID",
    // },
    {
      Name:      "Chicago COVID-19 Community Vulnerability Index (CCVI)",
      TableName: "ChicagoCovid19CommunityVulnerabilityIndex",
      Url:       "https://data.cityofchicago.org/resource/2ns9-phjk",
      SortField: "community_area_or_zip",
    },
    // {
    //   Name:      "Daily Chicago COVID-19 Cases, Deaths, and Hospitalizations - Historical",
    //   TableName: "Covid19Reports",
    //   Url:       "https://data.cityofchicago.org/resource/naz8-j4nc",
    // },
    // // This should stay out for now...
    // {
    //   Name:      "COVID-19 Cases, Tests, and Deaths by ZIP Code - Historical",
    //   TableName: "covid_zip_code",
    //   Url:       "https://data.cityofchicago.org/resource/yhhz-zm2v",
    // },
    // {
    //   Name:      "Public Health Statistics - Selected public health indicators by Chicago community area - Historical",
    //   TableName: "PublicHealthStatistics",
    //   Url:       "https://data.cityofchicago.org/resource/iqnk-2tcu",
    // },
    // TableName: "ChicagoCovid19CommunityVulnerabilityIndex",
	}

  // Initialize the database connection pool
  if err := dbconnector.InitDB(); err != nil {
    log.Fatalf("Could not initialize database: %v", err)
  }
  defer dbconnector.CloseDB() // Ensure that the pool is closed when the program exits

	fmt.Println("Pulling from data sources")

	// for title, url := range dataSets {
	for _, dataSet := range dataSets {
		fmt.Println(fmt.Sprintf("%s%s", "Pulling ", dataSet.Name))
		err := pull.GetAllData(dataSet.Url, dataSet.TableName, dataSet.SortField)
		if err != nil {
			log.Fatal(err)
		}
		dbconnector.CountData(dataSet.TableName)
	}
}
