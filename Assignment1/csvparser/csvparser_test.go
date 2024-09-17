package csvparser

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	data, err := ReadCSV("../testdata/easy.csv")
	expected_data := [][]string{
		{"Easy", "Test", "Data"},
		{"123", "456", "789"},
		{"101", "121", "131"},
		{"415", "161", "171"},
		{"181", "920", "212"},
	}

	if err != nil {
		t.Errorf("Error raised reading from csv file")
	} else if !reflect.DeepEqual(data, expected_data) {
		t.Errorf("Data read from file does not match what was expected")
	}
}

func TestReadCSVErrors(t *testing.T ) {
	data, err := ReadCSV("../testdata/this_file_does_not_exist")
	if data != nil || err == nil {
		t.Errorf("Should raise error and return nil data if reading from file that doesn't exist")
	}

	data, err = ReadCSV("../testdata/malformed.csv")
	if data != nil || err == nil {
		t.Errorf("Should raise error and return nil data if reading from file that is malformed")
	}

	data, err = ReadCSV("../testdata/easy.json")
	if data != nil || err == nil {
		t.Errorf("Should raise error and return nil data if reading from JSON file")
	}
}

func TestCsvDataByColumn(t *testing.T) {
	data := [][]string{
		{"Easy", "Test", "Data"},
		{"123", "456", "789"},
		{"101", "121", "131"},
		{"415", "161", "171"},
		{"181", "920", "212"},
	}

	expectedDataByColumn := map[string][]float64{
		"Easy": {123.0, 101.0, 415.0, 181.0},
		"Test": {456.0, 121.0, 161.0, 920.0},
		"Data": {789.0, 131.0, 171.0, 212.0},
	}

	dataByColumn, err := CsvDataByColumn(data)
	if err != nil {
		t.Errorf("Error raised parsing data")
	} else if !reflect.DeepEqual(dataByColumn, expectedDataByColumn) {
		t.Errorf("Parsed data does not match what was expected")
	}
}
