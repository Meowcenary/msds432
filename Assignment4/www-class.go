package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type MSDSCourse struct {
  CID string `json:"courseI_D`
  CNAME string `json:"course_name"`
  CPREREQ string `json:"prerequisite"`
  LastAccess string
}

// CSV resides in the current directory
// This file needs to be compiled with class-handlers.go, but should not be compiled
// with www-phone.go or handlers.go
var CSVFILE = "./course_data.csv"

type MSDSCourseCatalog []MSDSCourse

var data = MSDSCourseCatalog{}
var index map[string]int

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := MSDSCourse{
			CID:       line[0],
			CNAME:     line[1],
			CPREREQ:   line[2],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.CID, row.CNAME, row.CPREREQ, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.CID
		index[key] = i
	}
	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initS(cid, cname, cprereq string) *MSDSCourse {
	// cid and ccname require a value
	if cid == "" || cname == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &MSDSCourse{CID: cid, CNAME: cname, CPREREQ: cprereq, LastAccess: LastAccess}
}

func insert(course *MSDSCourse) error {
	// If it already exists, do not add it
	_, ok := index[(*course).CID]
	if ok {
		return fmt.Errorf("%s already exists", course.CID)
	}

	*&course.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	data = append(data, *course)
	// Update the index
	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *MSDSCourse {
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func list() string {
	var all string
	for _, k := range data {
		all = all + k.CID + " " + k.CNAME + " " + k.CPREREQ + "\n"
	}
	return all
}

func main() {
	err := readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	} else {
		for key := range index {
    	fmt.Println(key)
    }
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert/", http.HandlerFunc(insertHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
