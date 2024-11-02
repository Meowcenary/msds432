package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/kelvins/geocoder"
)

type PickupLocation struct {
    Latitude  float64
    Longitude float64
    Zipcode   string
}

func main() {
    geocoder.ApiKey = "REDACTED - Replace with key"

    // Download data
    jsonResponse, err := http.Get("https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=100")
    if err != nil {
        log.Fatal(err)
    }
    defer jsonResponse.Body.Close()

    jsonData, err := io.ReadAll(jsonResponse.Body)
    if err != nil {
        log.Fatal(err)
    }

    // Read in data as a slice of maps
    var tripData []map[string]interface{}
    err = json.Unmarshal(jsonData, &tripData)
    if err != nil {
        log.Fatal(err)
    }

    for _, record := range tripData {
        var pickupLocation PickupLocation

        // Extract latitude and longitude
        if lat, exists := record["pickup_centroid_latitude"].(string); exists {
            fmt.Sscanf(lat, "%f", &pickupLocation.Latitude)
        }
        if lon, exists := record["pickup_centroid_longitude"].(string); exists {
            fmt.Sscanf(lon, "%f", &pickupLocation.Longitude)
        }

        // Use geocoder to find the zipcode
        location := geocoder.Location{
            Latitude:  pickupLocation.Latitude,
            Longitude: pickupLocation.Longitude,
        }
        pickupLocation.Zipcode, err = ZipFromLongLat(location)
        if err != nil {
            log.Printf("Error getting zipcode: %v", err)
        }

        fmt.Printf("%+v\n", pickupLocation)
    }
}

func ZipFromLongLat(location geocoder.Location) (string, error){
	addresses, err := geocoder.GeocodingReverse(location)

	if err != nil {
		fmt.Println("Could not get the addresses: ", err)
		return "N/A", nil
	} else {
		address := addresses[0]
		return address.PostalCode, nil
	}
}
