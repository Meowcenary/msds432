// This section of the assignment was completed without generics to keep things as simple as possible.
// The go routines are more challenging for me to wrap my head around, so I wanted to keep the code as
// straightforward as possible while learning about them.
// Tour of Go: Channels - https://go.dev/tour/concurrency/2
// Why use one way channels - https://stackoverflow.com/questions/13596186/whats-the-point-of-one-way-channels-in-go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/kelvins/geocoder"
)

// This can be adjusted per system
const MAXROUTINES = 4

func main() {
	// generators has been modified to run a finite loop
	generatorInt := func(count int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			// Seed the random number generator to get different results each run
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < count; i++ { // Generate 'count' number of random integers
				intStream <- rand.Int()
			}
		}()
		return intStream
	}

	generatorFloat := func(count int) <-chan float64 {
		floatStream := make(chan float64)
		go func() {
			defer close(floatStream)
			// Seed the random number generator to get different results each run
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < count; i++ { // Generate 'count' number of random integers
				floatStream <- rand.Float64()
			}
		}()
		return floatStream
	}

	multiplyInt := func(intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)

		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				multipliedStream <- i * multiplier
			}
		}()

		return multipliedStream
	}

	multiplyFloat := func(floatStream <-chan float64, multiplier float64) <-chan float64 {
		multipliedStream := make(chan float64)

		go func() {
			defer close(multipliedStream)
			for i := range floatStream {
				multipliedStream <- i * multiplier
			}
		}()

		return multipliedStream
	}

	addInt := func(intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)

		go func() {
			defer close(addedStream)
			for i := range intStream {
				addedStream <- i + additive
			}
		}()
		return addedStream
	}

	addFloat := func(floatStream <-chan float64, additive float64) <-chan float64 {
		addedStream := make(chan float64)

		go func() {
			defer close(addedStream)
			for i := range floatStream {
				addedStream <- i + additive
			}
		}()
		return addedStream
	}

	var experimentSizes = []int{10000, 100000, 1000000, 1000000000}
	for _, v := range experimentSizes {
		fmt.Println("Size: ", v)

		// Integers
		intStream := generatorInt(v)
		startIntTime := time.Now()
		// Because the values are not being printed, just use _ for the var name
		_ = multiplyInt(addInt(multiplyInt(intStream, 2), 1), 2)
		endIntTime := time.Now()
		totalIntTime := endIntTime.Sub(startIntTime)

		// Floats
		floatStream := generatorFloat(v)
		startFloatTime := time.Now()
		_ = multiplyFloat(addFloat(multiplyFloat(floatStream, 2), 1), 2)
		endFloatTime := time.Now()
		totalFloatTime := endFloatTime.Sub(startFloatTime)

		fmt.Println("Runtime for integers:", totalIntTime, "Runtime for floats:", totalFloatTime)
	}

	// Previously was just reading records, but decided to do lookups instead
	// startReadRecordsInChunks := time.Now()
	// _ = readRecordsInChunks()
	// endReadRecordsInChunks := time.Now()
	// totalReadRecordsInChunks := endReadRecordsInChunks.Sub(startReadRecordsInChunks)
	// fmt.Println("Runtime for readRecordsInChunks:", totalReadRecordsInChunks)

	// This is going to run in parallel and as a result fire off a lot of API requests fast. It can
	// be modified to just read the file, but will require some changes
	startZipLookup := time.Now()
	results := performLookupInChunks()
	endZipLookup := time.Now()
	totalZipLookupTime := endZipLookup.Sub(startZipLookup)
	fmt.Println("Runtime for zipcodelookup:", totalZipLookupTime)
	for result := range results {
			fmt.Println("zip:", result)
	}
}

func performLookupInChunks() <-chan string {
    filePath := "Traffic_Crashes_-_Crashes.csv"
		geocoder.ApiKey = "replace with geocoder api"

    // Open the file to count the total number of rows
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil
    }
    defer file.Close()

    reader := csv.NewReader(file)

    // Skip the header row and count the total number of rows
    _, err = reader.Read() // Skip header
    if err != nil {
        fmt.Println("Error reading header:", err)
        return nil
    }

    totalRows := 0
    for {
        _, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error counting rows:", err)
            return nil
        }
        totalRows++
    }

    chunkSize := totalRows / MAXROUTINES

    var waitGroup sync.WaitGroup
		// Update this channel to a [][]string channel to hold records instead of zips
    zipChan := make(chan string)

    // Launch goroutines to read chunks
    for i := 0; i < MAXROUTINES; i++ {
        startRow := i * chunkSize

        // Adjust the chunk size for the last routine to cover any remaining rows
        size := chunkSize
        if i == MAXROUTINES-1 {
            size += totalRows % MAXROUTINES
        }

        waitGroup.Add(1)
        go readChunk(filePath, zipChan, &waitGroup, startRow, size)
    }

    // Close zipChan  after all goroutines are done
    go func() {
        waitGroup.Wait()
        close(zipChan)
    }()

    return zipChan
}

func readChunk(filePath string, zipChan chan<- string, wg *sync.WaitGroup, startRow, chunkSize int) {
    fmt.Println("Goroutine started for startRow:", startRow)
    defer wg.Done()

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Skip to the startRow
    for i := 0; i < startRow; i++ {
        _, err := reader.Read()
        if err != nil {
            fmt.Println("Error skipping rows:", err)
            return
        }
    }

    // Read the chunk of rows and perform lookup
    for i := 0; i < chunkSize; i++ {
        record, err := reader.Read()
        if err != nil {
            if err == io.EOF {
                break // Stop at the end of the file
            }
            fmt.Println("Error reading row:", err)
            return
        }

        // Extract latitude and longitude from the record (columns 46 and 47)
        latitude, err := strconv.ParseFloat(record[46], 64)
        if err != nil {
            fmt.Println("Error parsing latitude:", err)
            continue
        }
        longitude, err := strconv.ParseFloat(record[47], 64)
        if err != nil {
            fmt.Println("Error parsing longitude:", err)
            continue
        }

        // Perform reverse geocoding lookup for the current record
        location := geocoder.Location{Latitude: latitude, Longitude: longitude}
        zip, geoErr := zipFromLongLat(location)

        if geoErr != nil {
            fmt.Println("Geocoding error:", geoErr)
            continue
        }

        // Send the zip code result to the results channel
        zipChan <- zip
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
