package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Run the experiment for int64 and float64 types
func main() {
	// I noticed that it was able to process as many as a 100,000,000 records without issue, but
	// once it hit 1,000,000,000, everything slowed down considerably. I think this is due to
	// memory issues, but I'm not really sure. The runtimes for 1 Billion records were 1m15.81301957s
	// for integers and 2m44.058801262s for floats
	// experimentSizes := []int64{1000000000}

	experimentSizes := []int64{10000, 100000, 1000000}

	for i := range experimentSizes {
		experiment(experimentSizes[i])
	}
}

// Run a trial with an array of the specified size
func experiment(size int64) {
	// Generate random floats and run experiment
	randomInts := createRandomValues[int64](size)

	startIntegerTime := time.Now()
	// Generate random int64 and run experiment
	for _, _ = range addValues(multiplyValues(randomInts, 2), 1) {
		continue
	}
	endIntegerTime := time.Now()
	totalIntegerTime := endIntegerTime.Sub(startIntegerTime)

	startFloatTime := time.Now()
	// Generate random int64 and run experiment
	for _, _ = range addValues(multiplyValues(randomInts, 2), 1) {
		continue
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

func multiplyValues[valueType int64 | float64](values []valueType, multiplier valueType) []valueType {
	multipliedValues := make([]valueType, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func addValues[valueType int64 | float64](values []valueType, additive valueType) []valueType {
	addedValues := make([]valueType, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}
