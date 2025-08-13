package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ianhecker/weather-service/client"
	"github.com/ianhecker/weather-service/data"
	"github.com/ianhecker/weather-service/weather"
)

func main() {
	client := client.NewClient()

	request, err := client.NewRequest(http.MethodGet, "https://api.weather.gov/gridpoints/TOP/31,80/forecast")
	checkErr(err)

	response, err := client.Do(request)

	var data data.Data
	err = data.UnmarshalJSON(response)
	checkErr(err)

	period, err := data.GetPeriod(0)
	checkErr(err)

	var weather weather.Weather
	err = weather.UnmarshalJSON(period)
	checkErr(err)

	fmt.Printf("%+v\n", weather)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
