package pull

import (
	"fmt"
)

func GetAllData(url string) error {
	request := soda.NewGetRequest(url, "")
	request.Format = "json"
	request.Query.AddOrder("zipcode", soda.DirAsc)

	offsetRequest, err := soda.NewOffsetGetRequest(request)
	if err != nil {
		return err
	}

	for i := 0; i < 4; i++ {

		offsetRequest.Add(1)
		go func() {
			defer offsetRequest.Done()

			for {
				resp, err := offsetRequest.Next(2000)
				if err == soda.ErrDone {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				results := make([]map[string]interface{}, 0)
				err = json.NewDecoder(resp.Body).Decode(&results)
				resp.Body.Close()
				if err != nil {
					log.Fatal(err)
				}
				//Process your data
			}
		}()

	}
	offsetRequest.Wait()

	return nil
}
