package main

import (
	"fmt"
	"math/rand"
)

// Integer Experiments
// Function for running the experiment with integer values
func experimentInt() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}

// Return an array of random floats between FLOATMIN FLOATMAX
// Float Experiments
func createRandomFloats(size) []float64 {
	randomFloats := make([]float64, size)

	for i := range size {
		randomFloats[i] = rand.Float64()
	}

	return randomFloats
}

func multiplyFloats(values []float64, multiplier float64) []float64 {
	multipliedValues := make([]float64, len(values))

	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}

	return multipliedValues
}

func addFloats(values []float64, additive float64) []float64 {
	addedValues := make([]float64, len(values))

	for i, v := range values {
		addedValues[i] = v + additive
	}

	return addedValues
}

// Function for running the experiment with float values
// size - the size of the experiment set
// returns - nothing
func experimentFloat(size int) {
	multiply := multiplyFloats
	add := addFloats

	var randomFloats := createRandomFloats(size)
	for _, v := range add(multiply(randomFloats, 2), 1) {
		fmt.Println(v)
	}
}
