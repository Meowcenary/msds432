package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/kelvins/geocoder"
)

// Run the experiment
func main() {
	// This was significantly faster than fig-functional-pipeline-combination.go and
	// the integers consistently ran faster than the floats. I had expected integers to
	// outperform floats in fig-functional-pipeline-combination.go, but this was not the
	// case always.
	experimentSizes := []int64{10000, 100000, 1000000}
	// experimentSizes := []int64{1000000000}

	for i := range experimentSizes {
		experiment(experimentSizes[i])
	}

	// Run the zip code lookup
	records := readTrafficCrashCSV()
	startZipCodeLookupTime := time.Now()
	parseCrashData(records)
	endZipCodeLookupTime := time.Now()
	totalZipCodeLookupTime := endZipCodeLookupTime.Sub(startZipCodeLookupTime)
	fmt.Println("Runtime for zipcode lookup:", totalZipCodeLookupTime)
}

// Read crash data csv data
// Recycled from Assignment 1
func readTrafficCrashCSV() [][]string {
		// Open file
		file, err := os.Open("Traffic_Crashes_-_Crashes.csv")
		if err != nil {
			return nil
    }
		// Defer keyword allows close call to be declared next to open call, but delays execution to end of function
		defer file.Close()

		// Read records from file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
    if err != nil {
			return nil
    }

		return records
}

// Lookup the zipcdoes for the trafic data
// This code is in part recycled from Assignment 1
func parseCrashData(data [][]string) {
	// I accidentally committed my API key to the repo I am storing my homework in,
	// so once again I am regenerating it. I am going to leave the replace with api
	// key message when I commit and not try to do anything fancy with storing the key
	geocoder.ApiKey = "replace with geocoder api"

	for _, row := range data {
		latitude, _ := strconv.ParseFloat(row[46], 64)
		longitude, _ := strconv.ParseFloat(row[47], 64)

		// Reverse geocoding - careful, this costs money!
		location := geocoder.Location{
				Latitude: latitude,
				Longitude: longitude,
		}
		zip, geoErr := zipFromLongLat(location)

		// Check for errors
		if geoErr != nil {
			panic(geoErr)
		} else {
			// printing is for testing and should be removed for benchmarking
			fmt.Println("Zipcode:", zip)
		}
	}
}

// Taken from Assignment 1
func zipFromLongLat(location geocoder.Location) (string, error){
	addresses, err := geocoder.GeocodingReverse(location)

	// Need to exract the zip code from this
	if err != nil {
		// Handle errors and return generic string
		fmt.Println("Could not get the addresses: ", err)
		return "N/A", nil
	} else {
		// Usually, the first address returned from the API
		// is more detailed, so let's work with it
		address := addresses[0]
		return address.PostalCode, nil
	}
}

func experiment(size int64) {
	randomInts := createRandomValues[int64](size)
	startIntegerTime := time.Now()
	for _, v := range randomInts {
		multiplyValues(addValues(multiplyValues(v, 2), 1), 2)
	}
	endIntegerTime := time.Now()
	totalIntegerTime := endIntegerTime.Sub(startIntegerTime)

	randomFloats := createRandomValues[float64](size)
	startFloatTime := time.Now()
	for _, v := range randomFloats {
		multiplyValues(addValues(multiplyValues(v, 2.0), 1.0), 2.0)
	}
	endFloatTime := time.Now()
	totalFloatTime := endFloatTime.Sub(startFloatTime)

	fmt.Println("Size: ", size, "\nRuntime for integers:", totalIntegerTime, "\nRuntime for floats:", totalFloatTime)
}

// Using generics to reduce the amount of code
func createRandomValues[valueType int64 | float64](size int64) []valueType {
	randomValues := make([]valueType, size)
	for i := range size {
		typeForRand := reflect.TypeOf(*new(valueType))

		switch typeForRand.Kind() {
		case reflect.Int64:
			randomValues[i] = any(rand.Int63()).(valueType)
		case reflect.Float64:
			randomValues[i] = any(rand.Float64()).(valueType)
		}
	}
	return randomValues
}

func addValues[valueType int64 | float64](value valueType, additive valueType) valueType {
	return value + additive
}

func multiplyValues[valueType int64 | float64](value valueType, multiplier valueType) valueType {
	return value * multiplier
}
