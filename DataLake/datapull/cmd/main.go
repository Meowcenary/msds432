package main

import (
	"fmt"
	"log"

	"datapull/internal/dbconnector"
	"datapull/internal/pull"
)

type Dataset struct {
    Name string
    TableName string
		Url string
    // Model
}

func main() {
	dataSets := []Dataset{
    {
    	Name:      "Taxi Trips (2013-2023)",
      TableName: "TaxiTrips",
      Url:       "https://data.cityofchicago.org/resource/wrvz-psew",
    },
    {
      Name:      "Transportation Network Providers - Trips (2018 - 2022)",
      TableName: "TransportationNetworkProvidersTrips",
      Url:       "https://data.cityofchicago.org/resource/m6dm-c72p",
    },
    {
      Name:      "City of Chicago Building Permits",
      TableName: "PublicHealthStatistics",
      Url:       "https://data.cityofchicago.org/resource/ydr8-5enu",
    },
    {
      Name:      "Chicago COVID-19 Community Vulnerability Index (CCVI)",
      TableName: "BuildingPermits",
      Url:       "https://data.cityofchicago.org/resource/2ns9-phjk",
    },
    {
      Name:      "Daily Chicago COVID-19 Cases, Deaths, and Hospitalizations - Historical",
      TableName: "Covid19Reports",
      Url:       "https://data.cityofchicago.org/resource/naz8-j4nc",
    },
    // {
    //   Name:      "COVID-19 Cases, Tests, and Deaths by ZIP Code - Historical",
    //   TableName: "covid_zip_code",
    //   Url:       "https://data.cityofchicago.org/resource/yhhz-zm2v",
    // },
    {
      Name:      "Public Health Statistics - Selected public health indicators by Chicago community area - Historical",
      TableName: "ChicagoCovid19CommunityVulnerabilityIndex",
      Url:       "https://data.cityofchicago.org/resource/iqnk-2tcu",
    },
	}

	fmt.Println("Pulling from data sources")

	// for title, url := range dataSets {
	for _, dataSet := range dataSets {
		fmt.Println(fmt.Sprintf("%s%s", "Pulling ", dataSet.Name))
		err := pull.GetAllData(dataSet.Url)
		if err != nil {
			log.Fatal(err)
		}
		dbconnector.CountData(dataSet.TableName)
	}
}
