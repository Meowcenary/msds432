package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

// Create a data set of size, record time to serialize as json, record time to deserialize from json
func runExperiment(size int) {
	fmt.Println("Experiment with size ", size)
	// Create sample data
	var i int
	var t Data

	for i = 0; i < size; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)

	// Measure time to serialize
	startSerialize := time.Now()
	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Record the end time
	endSerialize := time.Now()
	serializeTime := endSerialize.Sub(startSerialize)
	fmt.Println("Serialize time:", serializeTime)
	// Debugging
	// fmt.Print("After Serialize:", buf)

	// Measure time to deserialize
	decoder := json.NewDecoder(buf)
	var temp []Data
	startDeserialize := time.Now()
	err = DeSerialize(decoder, &temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	endDeserialize := time.Now()
	deserializeTime := endDeserialize.Sub(startDeserialize)
	fmt.Println("Deserialize time:", deserializeTime)
	// Debugging
	// fmt.Println("After DeSerialize:")
  //
	// Print the deserialized data
	// for index, value := range temp {
	// 	fmt.Println(index, value)
	// }
}

func main() {
	// Run the experiment in different sizes
	runExperiment(10000)
	runExperiment(100000)
	runExperiment(1000000)
}
