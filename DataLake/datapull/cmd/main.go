package main

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-soda"
)

func main() {
	fmt.Println("main function of data pull")

	dataSets := map[string]string{
		"Taxi Trips (2013-2023)": "https://data.cityofchicago.org/resource/wrvz-psew",
		"Transportation Network Providers - Trips (2018 - 2022)":  "https://data.cityofchicago.org/resource/m6dm-c72p",
		// V2 Endpoint - https://data.cityofchicago.org/OData.svc/ydr8-5enu
		// "City of Chicago Building Permits": "https://data.cityofchicago.org/api/odata/v4/ydr8-5enu",
		"Chicago COVID-19 Community Vulnerability Index (CCVI)": "https://data.cityofchicago.org/resource/2ns9-phjk",
		"Daily Chicago COVID-19 Cases, Deaths, and Hospitalizations - Historical": "https://data.cityofchicago.org/resource/naz8-j4nc",
		"COVID-19 Cases, Tests, and Deaths by ZIP Code - Historical": "https://data.cityofchicago.org/resource/yhhz-zm2v",
		"Chicago Communities by Neighborhood and Zip Code": "",
	}

	for title, url := range dataSets {
		fmt.Println(fmt.Sprintf("%s%s", "Pulling ", title))
		sodareq := soda.NewGetRequest(url, "")
		//count all records
		count, err := sodareq.Count()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count)
	}
}
