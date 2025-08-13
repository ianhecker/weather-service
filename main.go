package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ianhecker/weather-service/client"
	"github.com/ianhecker/weather-service/response"
	"github.com/ianhecker/weather-service/weather"
)

func main() {
	client := client.NewClient()

	req, err := client.NewRequest(http.MethodGet, "https://api.weather.gov/points/39.7456,-97.0892")
	checkErr(err)

	resp, err := client.Do(req)
	checkErr(err)

	var forecastURL response.ForecastURL
	err = forecastURL.UnmarshalJSON(resp)
	checkErr(err)

	url, err := forecastURL.GetURL()
	checkErr(err)

	request, err := client.NewRequest(http.MethodGet, url)
	checkErr(err)

	resp, err = client.Do(request)
	checkErr(err)

	var periods response.Periods
	err = periods.UnmarshalJSON(resp)
	checkErr(err)

	period, err := periods.GetPeriod(0)
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
