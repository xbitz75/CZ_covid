package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	country := "czech-republic"
	status := "confirmed"

	data := getDataByCountry(country, status)
	models := parse(data)

	models.print()
	fmt.Println(fmt.Sprintf("Latest: %v", models.latest()))
	fmt.Println(fmt.Sprintf("Most Active: %v", models.mostActive()))
}

func getDataByCountry(country string, status string) []byte {
	url := fmt.Sprintf("https://api.covid19api.com/live/country/%s/status/%s", country, status)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func parse(data []byte) models {
	var parsedData models
	if err := json.Unmarshal(data, &parsedData.data); err != nil {
		fmt.Println(err)
	}

	return parsedData
}

type models struct {
	data []model
}

func (m *models) print() {
	for _, e := range m.data {
		fmt.Println(e)
	}
}

func (m *models) latest() (latest model) {
	for _, e := range m.data {
		if latest.Date.Before(e.Date) {
			latest = e
		}
	}

	return latest
}

func (m *models) mostActive() (most model) {
	for _, e := range m.data {
		if most.Active < e.Active {
			most = e
		}
	}

	return most
}

type model struct {
	Country     string
	CountryCode string
	Province    string
	City        string
	CityCode    string
	Lat         string
	Lon         string
	Confirmed   int
	Deaths      int
	Recovered   int
	Active      int
	Date        time.Time
}

func (m model) String() string {
	return fmt.Sprintf(
		"CountryCode: %s, Confirmed: %d, Deaths: %d, Recovered: %d, Active: %d, Date: %s",
		m.CountryCode, m.Confirmed, m.Deaths, m.Recovered, m.Active, m.Date)
}

func (m model) StringAll() string {
	return fmt.Sprintf(
		"Country: %s, CountryCode: %s, Province: %s, City: %s, CityCode: %s, Lat: %s, Lon: %s, Confirmed: %d, Deaths: %d, Recovered: %d, Active: %d, Date: %s",
		m.Country, m.CountryCode, m.Province, m.City, m.CityCode, m.Lat, m.Lon, m.Confirmed, m.Deaths, m.Recovered, m.Active, m.Date)
}
