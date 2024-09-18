package main


import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"time"
)

func main() {
	// CLI args
	args := os.Args
	argumentError := "This program expects the path to the csv file as a single command line argument"

	if len(args) != 2 {
		fmt.Println(argumentError)
	} else {
		records, err := ReadCSV(args[1])

		if err != nil {
			log.Fatal(err)
		}

		for i := 1; i < 2; i++ { // i < len(records); i++ {
			fmt.Println(records[i][2])
			dayOfWeek, err := DayOfWeekFrom(records[i][2])

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(dayOfWeek)
		}

		crashesByYear, err := CrashesByYear(records[1:])

		fmt.Println(crashesByYear)

		crashes, err := ParseCrashData(records[1:])
		if err != nil {
			log.Fatal(err)
		}

		crashesByYearAndWeekday, err := CrashesByYearAndWeekday(crashes)
		if err != nil {
			log.Fatal(err)
		}

		PrettyPrintCrashesByYearAndWeekday(crashesByYearAndWeekday)
		// fmt.Println(crashesByYearAndWeekday)
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
func WriteCSV(filepath string, data string) error {
	// Ensure path is a csv file
	if path.Ext(filepath) != ".csv" {
		return &InvalidCsvPathError{Filepath: filepath}
	}

	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(data)
	writer.Flush()

	return nil
}

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
* Return the day of the week given the date as a string
*/
func DayOfWeekFrom(date string) (string, error) {
	// This is a specific reference time used by Go, but formatted to reflect the time format in the file
	t, err := time.Parse("01/02/2006 03:04:05 PM", date)

	if err != nil {
			panic(err)
	}

	return t.Weekday().String(), nil
}

/*
* Return the year from the date a string
*/
func YearFrom(date string)(int, error) {
	t, err := time.Parse("01/02/2006 03:04:05 PM", date)

	if err != nil {
			panic(err)
	}

	return t.Year(), nil
}

/*
* Struct to represent a record from the crash data
*/
type Crash struct {
	// consider using Weekday type: https://pkg.go.dev/time#Weekday
	weekday string
	date time.Time
	// set to zero as default
	zipcode int
}

/*
* Parse the crash data into Crash structs that can be used for different operations
*/
func ParseCrashData(data [][]string) ([]Crash, error){
	crashes := make([]Crash, 0)

	for _, row := range data {
		date, date_err := time.Parse("01/02/2006 03:04:05 PM", row[2])
		weekday := date.Weekday().String()

		// Check for errors
		if date_err != nil {
			panic(date_err)
		} else {
			crashes = append(crashes, Crash{weekday: weekday, date: date, zipcode: 0})
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

		// Print the inner map data
		for innerKey, value := range innerMap {
			fmt.Printf("    %s: %d\n", innerKey, value)
		}
	}
}

/*
* Return a map of year to number of crashes reported in that year (int)
*/
func CrashesByYear(data [][]string) (map[int]int, error) {
	crashesByYear := make(map[int]int)

	for _, row := range data {
		// column for date
		date := row[2]
		year, err := YearFrom(date)

		if err != nil {
			panic(err)
		}

		crashesByYear[year] += 1
	}

	return crashesByYear, nil
}

/*
*
*/
func CrashesByYearAndWeekDay(data [][]string) (map[int]int, error) {
	crashesByYear := make(map[int]int)

	for _, row := range data {
		// column for date
		date := row[2]
		year, err := YearFrom(date)

		if err != nil {
			panic(err)
		}

		crashesByYear[year] += 1
	}

	return crashesByYear, nil
}

/*
* Find zip code from longitude and latitude using Google geocoder API
*/
func ZipFromLongLat() (error) {
	return nil
}
