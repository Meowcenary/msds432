package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	// Empty import, only want side effects of import
	_ "github.com/lib/pq"
)

// Start of functions extracted from post05 package.
// Modifications have been made for MSDS courses
// Connection details
var (
	Hostname = "localhost"
	Port     = 5433
	Username = "postgres"
	Password = "root"
	Database = "MSDS"
)

// Userdata is for holding full user data
// Userdata table + Username
type MSDSCourse struct {
	CID     string
	CNAME   string
	CPREREQ string
}

func openConnection() (*sql.DB, error) {
	// connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	// open database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ListUsers lists all users in the database
func listCourses() ([]MSDSCourse, error) {
	Data := []MSDSCourse{}
	db, err := openConnection()
	if err != nil {
		return Data, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "CID","CNAME","CPREREQ"
		FROM "MSDSCourseCatalog"`)
	if err != nil {
		return Data, err
	}

	for rows.Next() {
		var CID string
		var CNAME  string
		var CPREREQ string
		err = rows.Scan(&CID, &CNAME, &CPREREQ)
		temp := MSDSCourse{CID: CID, CNAME: CNAME, CPREREQ: CPREREQ}
		Data = append(Data, temp)
		if err != nil {
			return Data, err
		}
	}
	defer rows.Close()
	return Data, nil
}

// The function returns the User ID of the username
// -1 if the user does not exist
func exists(CID string) bool {
	CID = strings.ToLower(CID)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT "CID" FROM "MSDSCourseCatalog" WHERE "CID" = '%s'`, CID)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println("Query error:", err)
		return false
	}
	defer rows.Close()

	// Check if any row was returned
	if rows.Next() {
			return true // CID exists
	}
	return false
}

// Insert a course to the database
func addCourse(course MSDSCourse) bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	// Check if the class already exists
	if exists(course.CID) {
		fmt.Println("CID already exists:", course.CID)
		return false
	}

	// SQL for inserting a MSDSCourse
	insertStatement := `INSERT INTO "MSDSCourseCatalog" ("CID", "CNAME", "CPREREQ") VALUES ($1, $2, $3)`
	_, err = db.Exec(insertStatement, course.CID, course.CNAME, course.CPREREQ)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return false
	}

	return true
}

// End of functions extracted from post05 package

// Class numbers should be 3 numbers
// They probably stop at something well before 999,
// but it seemed like an okay number for this assignment.
var MIN = 100
var MAX = 999

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Generate a class name
func getClassID() string {
	classNumber := random(MIN, MAX)
	return string("MSDS" + strconv.Itoa(classNumber))
}

func getClassName() string {
	return "Database Systems"
}

// Generate a random string of prereqs
func getClassPrereq() string {
	randomValue := random(MIN, MAX)
	// roughly ten percent chance of getting a pre-req
	if randomValue % 10 == 0 {
		coursesAvailable, _ := listCourses()
		if len(coursesAvailable) > 0 {
			index := randomValue % len(coursesAvailable)
			return coursesAvailable[index].CID
		} else {
			// No classes to pick from
			return ""
		}
	}

	return ""
}

func main() {
	data, err := listCourses()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	// Get seed for random course names
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	// var to hold courses for insertion
	var course MSDSCourse
	// insertion loop
	for i := 0; i < 5; i++ {
		course = MSDSCourse{
			CID: getClassID(),
			CNAME: getClassName(),
			CPREREQ: getClassPrereq(),
		}

		if !addCourse(course) {
			fmt.Println("There was an error adding course", course.CID)
		}
	}

	data, err = listCourses()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}
}
