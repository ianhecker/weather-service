package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ianhecker/weather-service/client"
	"github.com/ianhecker/weather-service/coordinates"
	"github.com/ianhecker/weather-service/response"
	"github.com/ianhecker/weather-service/weather"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", weatherHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Println("Server listening on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	latStr := r.URL.Query().Get("latitude")
	longStr := r.URL.Query().Get("longitude")

	coords, err := coordinates.NewCoordinates(latStr, longStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":{"code":"bad_request","message":"%s"}}`, err)
		return
	}

	baseURL := "https://api.weather.gov/points/%0.5f,%0.5f"
	url := fmt.Sprintf(baseURL, coords.Latitude, coords.Longitude)

	client := client.NewClient()
	weather, err := fetchWeather(client, url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":{"code":"internal_error","message":"%s"}}`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(weather)
}

func fetchWeather(client *client.Client, url string) (*weather.Weather, error) {

	req, err := client.NewRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var forecastURL response.ForecastURL
	err = forecastURL.UnmarshalJSON(resp)
	if err != nil {
		return nil, err
	}

	url, err = forecastURL.GetURL()
	if err != nil {
		return nil, err
	}

	request, err := client.NewRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	var periods response.Periods
	err = periods.UnmarshalJSON(resp)
	if err != nil {
		return nil, err
	}

	period, err := periods.GetPeriod(0)
	if err != nil {
		return nil, err
	}

	var weather weather.Weather
	err = weather.UnmarshalJSON(period)

	return &weather, err
}
