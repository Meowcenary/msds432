package pull

import (
	"fmt"
	"encoding/json"
	"errors"
	"log"
	"reflect"

	"datapull/internal/dbconnector"
	"datapull/internal/models"

	"github.com/SebastiaanKlippert/go-soda"
)


var UrlToModel = map[string]interface{}{
	"https://data.cityofchicago.org/resource/wrvz-psew": &models.TaxiTrip{},
  "https://data.cityofchicago.org/resource/m6dm-c72p": &models.TransportationNetworkProvidersTrip{},
	"https://data.cityofchicago.org/resource/ydr8-5enu": &models.BuildingPermit{},
	"https://data.cityofchicago.org/resource/2ns9-phjk": &models.ChicagoCovid19CommunityVulnerabilityIndex{},
	"https://data.cityofchicago.org/resource/iqnk-2tcu": &models.PublicHealthStatistic{},
	"https://data.cityofchicago.org/resource/yhhz-zm2v": &models.Covid19Report{},
}

func DataCounts(url string) error {
		sodareq := soda.NewGetRequest(url, "")
		//count all records
		count, err := sodareq.Count()
		if err != nil {
			return err
		}
		fmt.Println(count)

		return nil
}

// This should probably be refactored, so that the dataset types lives in this package, but for now it will do.
// sortField is necessary because OffsetGetRequest requires an ordering
func GetAllData(url string, tableName string, sortField string) error {
    // Determine model type from the URL to map to the correct struct
    model, exists := UrlToModel[url]
    if !exists {
        return errors.New("no model registered for this URL")
    }

    // Create the type and slice type for the model, only once (outside the goroutine)
    modelType := reflect.TypeOf(model).Elem() // e.g., models.TaxiTrip
    sliceType := reflect.SliceOf(modelType)   // e.g., []models.TaxiTrip

    // Create the request and offset request
    request := soda.NewGetRequest(url, "")
    request.Format = "json"
    request.Query.AddOrder(sortField, soda.DirAsc)
    offsetRequest, err := soda.NewOffsetGetRequest(request)
    if err != nil {
        return err
    }

    // Get the data using goroutines
    for i := 0; i < 10; i++ {
        offsetRequest.Add(1)
        go func() {
            defer offsetRequest.Done()

            for {
                resp, err := offsetRequest.Next(4000)
                if err == soda.ErrDone {
                    break
                }
                if err != nil {
                    log.Fatal(err)
                }

                // Create the fresh slice of the correct type inside the goroutine
                modelSlice := reflect.New(sliceType).Interface() // Fresh slice each time

                // Decode the response into the model slice (dynamically created)
                err = json.NewDecoder(resp.Body).Decode(modelSlice)
                resp.Body.Close()

                if err != nil {
                    fmt.Println("Error decoding response:", err)
                    log.Fatal(err)
                }

                // Now `modelSlice` is populated with the decoded data
                // We need to cast it back to the appropriate slice type
                actualSlice := reflect.ValueOf(modelSlice).Elem().Interface()
								modelSliceValue := reflect.ValueOf(actualSlice)

                // For each batch of data, call InsertData to insert into the DB
                for i := 0; i < modelSliceValue.Len(); i++ {
                  item := modelSliceValue.Index(i).Interface() // get each item in the slice
                  // fmt.Println(fmt.Sprintf("%+v", item))
                  if err := dbconnector.InsertData(tableName, item); err != nil {
                    log.Fatal(err)
                  }
                }

								fmt.Printf("Processed %d records\n", modelSliceValue.Len())
            }
        }()
    }

    // Wait for all goroutines to finish
    offsetRequest.Wait()

    return nil
}
