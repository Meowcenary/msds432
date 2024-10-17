package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Run the experiment
func main() {
	// This was significantly faster than fig-functional-pipeline-combination.go and
	// the integers consistently ran faster than the floats. I had expected integers to
	// outperform floats in fig-functional-pipeline-combination.go, but this was not the
	// case always.
	// experimentSizes := []int64{10000, 100000, 1000000}
	experimentSizes := []int64{1000000000}

	for i := range experimentSizes {
		experiment(experimentSizes[i])
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
