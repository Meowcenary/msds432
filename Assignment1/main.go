package main


import (
	// "bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"time"

	"github.com/kelvins/geocoder"
	// This is more complicated than the other geocoder
	// "googlemaps.github.io/maps"
)

func main() {
	// CLI args
	args := os.Args
	argumentError := "This program expects the path to the csv file as a single command line argument"

	geocoder.ApiKey = "Replace with API Key"

	if len(args) != 2 {
		fmt.Println(argumentError)
	} else {
		records, err := ReadCSV(args[1])
		if err != nil {
			log.Fatal(err)
		}

		// to avoid using too many credits, limit the testing data used
		// limit := 100
		// Create records from the data that can be used in the requirements, but skip the header row
		// Requirments 7 and 8 are handled by the function ZipFromLongLat which is itself called in this function
		// crashes, err := ParseCrashData(records[1:limit + 1])
		crashes, err := ParseCrashData(records[1:])
		if err != nil {
			log.Fatal(err)
		}

		// Requirement 5 - add new columns for day of the week for crash and zip code
		// This isn't the best way to do this, but it gets the job done
		// for i, row := range records[:limit + 1] {
		for i, row := range records {
		  // Add new headers
			if i == 0 {
				row = append(row, "Crash Day of Week", "Zip Code")
			} else {
				// Append new column values for each row
				// -1 is to account for the headers
				row = append(row, crashes[i-1].weekday, crashes[i-1].zipcode)
			}
			records[i] = row
    }

    // Create a new CSV file to write the updated data
    outputFile, err := os.Create("updated_crash_data.csv")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    // Create a CSV writer
    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    // Write the updated records to the new CSV
    err = writer.WriteAll(records)
    if err != nil {
        fmt.Println("Error writing to output file:", err)
        return
    }
    fmt.Println("Day of week and Zip code added and written out to updated_crash_data.csv")
		// End requirement 5

		// Requirement 9
		crashesByYearAndWeekday, err := CrashesByYearAndWeekday(crashes)
		if err != nil {
			log.Fatal(err)
		}

		crashesByYearAndWeekday2021 := make(map[int]map[string]int)
		crashesByYearAndWeekday2021[2021] = crashesByYearAndWeekday[2021]
		fmt.Println("--------------")
		fmt.Println("Requirement 9")
		fmt.Println("--------------")
		fmt.Println("Crashes by Year and Weekday")
		PrettyPrintCrashesByYearAndWeekday(crashesByYearAndWeekday2021)
		// End Requirement 9

		// Requirement 10
		crashesByZipCode, err := CrashesByZipCode(crashes)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("--------------")
		fmt.Println("Requirement 10")
		fmt.Println("--------------")
    // Pretty print the map , I decided not to sort the keys for this one
		fmt.Println("Crashes by Zipcode")
    for zipcode, total_crashes := range crashesByZipCode {
        fmt.Printf("%s: %d\n", zipcode, total_crashes)
    }
		// End Requirement 10

		fmt.Println("--------------")
		fmt.Println("Requirement 11")
		fmt.Println("--------------")
		// Requirement 11
		fmt.Println("Hit and Runs by Zipcode for 2020")
		hitAndRunsByZipcodeFor2020, _ := HitAndRunsByZipcodeForYear(2020, crashes)
    for zipcode, hit_and_runs := range hitAndRunsByZipcodeFor2020 {
        fmt.Printf("%s: %d\n", zipcode, hit_and_runs)
    }
		// End Requirment 11

		/* Requirement 12
		I was having a bit of trouble wrapping my head around this, but from my understanding, a confounder (confounding variable)
		is a variable that will affect both the independent and dependent variables distorting the relationship. In this instance
		the independent variable is zip codes and the dependent variable is the number of crashes.  My initial thought was that
		LIGHTING_CONDITION would not be a confounder and that WEATHER_CONDITION would be, but on further consideration I decided that
		LIGHTING_CONDITION would be a confounder and that WEATHER_CONDITION would not be. Weather across a zip code is broadly going
		to be the same, but lighting conditions will vary across the zip code. Because the lighting conditions vary, they need to be
		controlled for to understand the affect, but weather does not vary and therefore does not need to be controlled for.

		This was not particularly intuitive to me, so I'm going to keep looking into this more.
		*/
	}
}

//
// CSV Parsing code
//
type InvalidCsvPathError struct{
	Filepath string
}

func (i *InvalidCsvPathError) Error() string {
	return "Invalid CSV path: " + i.Filepath
}

// The proper way to handle this would be to use the csv library as in ReadCSV, but
// csv writer expects [][]string type which would mean refactoring things in a way
// that there is not time for. Best alternative was to write a formatted string
// func WriteCSV(filepath string, data string) error {
// 	// Ensure path is a csv file
// 	if path.Ext(filepath) != ".csv" {
// 		return &InvalidCsvPathError{Filepath: filepath}
// 	}
//
// 	file, err := os.Create(filepath)
// 	defer file.Close()
//
// 	if err != nil {
// 		return err
// 	}
//
// 	writer := bufio.NewWriter(file)
// 	writer.WriteString(data)
// 	writer.Flush()
//
// 	return nil
// }

func ReadCSV(filepath string) ([][]string, error) {
		// Ensure path is a csv file
		if path.Ext(filepath) != ".csv" {
			return nil, &InvalidCsvPathError{Filepath: filepath}
		}
		// Open file
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
    }
		// Defer keyword allows close call to be declared next to open call, but delays execution to end of function
		defer file.Close()

		// Read records from file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
    if err != nil {
			return nil, err
    }

		return records, err
}

// create mapping of CSV headers to all values in column
// data is rows of CSV data read from a file with headers
// only values that can be converted to float64 are kept
// as string values can not be used for regressions
func CsvDataByColumn(data [][]string) (map[string][]string, error) {
	headerIndex := make(map[int]string)
	dataByColumn := make(map[string][]string)

	// Pop header data off of data and create mapping to csv data
	headers, data := data[0], data[1:]
	for i, header := range headers {
		headerIndex[i] = header
	}

	for _, row := range data {
		for i, value := range row {
			header := headerIndex[i]
			dataByColumn[header] = append(dataByColumn[header], value)
		}
	}

	return dataByColumn, nil
}

/*
* Struct to represent a record from the crash data
*/
type Crash struct {
	// consider using Weekday type: https://pkg.go.dev/time#Weekday
	weekday string
	date time.Time
	// set to zero as default, store as string in case it uses hyphens
	zipcode string
	latitude float64
	longitude float64
	hit_and_run bool
}

/*
* Parse the crash data into Crash structs that can be used for different operations
*/
func ParseCrashData(data [][]string) ([]Crash, error){
	crashes := make([]Crash, 0)

	for _, row := range data {
		// row[i] is file specific, might be nice to genericize this
		date, dateErr := time.Parse("01/02/2006 03:04:05 PM", row[3])
		weekday := date.Weekday().String()
		latitude, _ := strconv.ParseFloat(row[46], 64)
		longitude, _ := strconv.ParseFloat(row[47], 64)
		hit_and_run_i := row[19]
		// Default to false
		hit_and_run_i_bool := false
		// If hit and run is "Y" it's true i.e a hit and run, but otherwise it is not
		if hit_and_run_i == "Y" {
			hit_and_run_i_bool = true
		}

		// Reverse geocoding - careful, this costs money!
		location := geocoder.Location{
				Latitude: latitude,
				Longitude: longitude,
		}
		zip, geoErr := ZipFromLongLat(location)

		// Check for errors
		if dateErr != nil {
			panic(dateErr)
		} else if geoErr != nil {
			panic(geoErr)
		} else {
			crashes = append(crashes, Crash{weekday: weekday, date: date, zipcode: zip, longitude: longitude, latitude: latitude, hit_and_run: hit_and_run_i_bool})
		}
	}

	return crashes, nil
}

/*
* Return a map of year to map of weekday to number of crashes in that weekday
*/
func CrashesByYearAndWeekday(crashes []Crash) (map[int]map[string]int, error) {
	crashesByYearAndWeekday := make(map[int]map[string]int)

	for i := 0; i < len(crashes); i++ {
		crash := crashes[i]
		year := crash.date.Year()
		weekday := crash.weekday

		_, exists := crashesByYearAndWeekday[year]
		if exists {
			_, exists = crashesByYearAndWeekday[year][weekday]

			if exists {
				crashesByYearAndWeekday[year][weekday] += 1
			} else {
				crashesByYearAndWeekday[year][weekday] = 1
			}
		} else {
			crashesByYearAndWeekday[year] = make(map[string]int)
			crashesByYearAndWeekday[year][weekday] = 1
		}
	}

	return crashesByYearAndWeekday, nil
}

func CrashesByZipCode(crashes []Crash) (map[string]int, error) {
	crashesByZipCode := make(map[string]int)

	for i := 0; i < len(crashes); i++ {
		crash := crashes[i]
		zipcode := crash.zipcode

		_, exists := crashesByZipCode[zipcode]
		if exists {
			crashesByZipCode[zipcode] += 1
		} else {
			crashesByZipCode[zipcode] = 1
		}
	}

	return crashesByZipCode, nil
}

func HitAndRunsByZipcodeForYear(year int, crashes []Crash) (map[string]int, error) {
	hitAndRunsByZipcode := make(map[string]int)

	for i := 0; i < len(crashes); i++ {
		crash := crashes[i]
		crash_year := crash.date.Year()

		// if the crash didn't happen in the specified year, skip
		if crash_year != year{
			continue
		}

		zipcode := crash.zipcode
		hit_and_run := crash.hit_and_run

		_, exists := hitAndRunsByZipcode[zipcode]
		// only count hit and runs
		if exists && hit_and_run {
			hitAndRunsByZipcode[zipcode] += 1
		} else if hit_and_run {
			hitAndRunsByZipcode[zipcode] = 1
		}
	}

	return hitAndRunsByZipcode, nil
}

// Requirements 7 and 8 are handled here
func ZipFromLongLat(location geocoder.Location) (string, error){
	addresses, err := geocoder.GeocodingReverse(location)

	// Need to exract the zip code from this
	if err != nil {
		// I initially was returning the error for this, but that caused the program to crash, so I'm returning "N/A" now
		fmt.Println("Could not get the addresses: ", err)
		return "N/A", nil
	} else {
		// Usually, the first address returned from the API
		// is more detailed, so let's work with it
		address := addresses[0]
		return address.PostalCode, nil
	}
}

func PrettyPrintCrashesByYearAndWeekday(m map[int]map[string]int) {
	// Extract the year keys from the outer map
	years := make([]int, 0, len(m))
	for year := range m {
		years = append(years, year)
	}

	// Sort the years
	sort.Ints(years)

	// Print the map using the sorted years
	for _, year := range years {
		fmt.Printf("%d:\n", year)

		// Extract the inner map for this year
		innerMap := m[year]

		// Extract the inner keys (weekdays)
		innerKeys := make([]string, 0, len(innerMap))
		for innerKey := range innerMap {
			innerKeys = append(innerKeys, innerKey)
		}

		// Sort the inner keys (weekdays) using the custom weekday order
		sort.Slice(innerKeys, func(i, j int) bool {
			return weekdayIndex(innerKeys[i]) < weekdayIndex(innerKeys[j])
		})

		for _, innerKey := range innerKeys {
			fmt.Printf("    %s: %d\n", innerKey, innerMap[innerKey])
		}
	}
}

// Weekday sorting
var weekdayOrder = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday"}

func weekdayIndex(weekday string) int {
	for i, day := range weekdayOrder {
		if day == weekday {
			return i
		}
	}
	return len(weekdayOrder) // In case of unexpected input, put it at the end
}
