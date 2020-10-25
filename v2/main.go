package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// source: https://onemocneni-aktualne.mzcr.cz/api/v2/covid-19

func main() {
	overview := getOverview()
	fmt.Println(overview)

	cumulative := getCumulativeCollection()
	fmt.Println(cumulative.data)
	cumulative.plotDeaths()
	cumulative.plotInfected()
	cumulative.plotRecovered()
	cumulative.plotTested()
}

func getData(url string) []byte {
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
