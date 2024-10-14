package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

type MSDSCourse struct {
	CID string `json:"id"`
	CNAME string `json:"name"`
	CDESC string `json:"description"`
	CPREREQ string `json:"prerequisites"`
}

func createMsdsCourse() MSDSCourse {
	// Course ids will be 8 characters and always start with MSDS
	var id = getIdString()
	var name = getName()
	var description = getDescription()
	var prereqs = getPrereqs()

	return MSDSCourse{CID: id, CNAME: name, CDESC: description, CPREREQ: prereqs}
}

// Create a random class id
func getIdString() string {
	var CMIN = 101
	var CMAX = 999
	var id = "MSDS"
  var courseNumber = rand.Intn(CMAX-CMIN) + CMIN
	id += strconv.Itoa(courseNumber)

	return id
}

// Generate a name for a course
func getName() string {
	var names = []string{
		"Database Systems",
		"Programming",
		"Networking",
		"Parallel Processing",
		"Data Engineering",
		"Architecture",
		"Visualization",
		"AI / Machine Learning",
		"Computer Vision",
		"Cloud Computing",
		"Big Data",
	}
	// Return random index
	return names[rand.Intn(len(names))]
}

// Generate a description for a course
func getDescription() string {
	var descriptions = []string{
		"Introductory course that explores several systems",
		"Advanced course that focuses on case study discussion",
		"Devising high level plans and putting them into action",
		"Course that focuses on management",
		"Elective course",
	}
	// Return random index
	return descriptions[rand.Intn(len(descriptions))]
}

// Generate pre reqs for a course
func getPrereqs() string {
	// zero to three prereqs
	var prereqs = rand.Intn(4)

	if prereqs == 0 {
		return "None"
	} else {
		var prereqStr = ""

		for i := 0; i < prereqs; i++ {
			if i > 0 {
				prereqStr += ", "
			}
			prereqStr += getName()
		}

		return prereqStr
	}
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Generate 5 courses into a slice, serialize to JSON, deserialize and print
func main() {
	// Create a slice of length L to store the courses
	var courses [5]MSDSCourse
	// Generate the courses and store to the slice
	for i := 0; i < len(courses); i++ {
		courses[i] = createMsdsCourse()
	}

	// JSON encoding
	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)

	encoder := json.NewEncoder(buf)
	err := Serialize(encoder, courses)
	if err != nil {
		fmt.Println(err)
		return
	}

	// JSON Decoding
	decoder := json.NewDecoder(buf)
	var temp [5]MSDSCourse
	err = DeSerialize(decoder, &temp)
	fmt.Println("After DeSerialize:")
	for index, value := range temp {
		fmt.Println(index, value)
	}
}
