package weather

import "encoding/json"

type Weather struct {
	Name             string           `json:"name"`
	StartTime        string           `json:"startTime"`
	EndTime          string           `json:"endTime"`
	IsDaytime        bool             `json:"isDaytime"`
	Temperature      int              `json:"temperature"`
	TempUnit         string           `json:"temperatureUnit"`
	WindSpeed        string           `json:"windSpeed"`
	WindDir          string           `json:"windDirection"`
	Short            string           `json:"shortForecast"`
	Detailed         string           `json:"detailedForecast"`
	Characterization Characterization `json:"characterization"`
}

func (weather Weather) CharacterizeWeather() {}

func (weather *Weather) UnmarshalJSON(bytes []byte) error {
	type Alias Weather
	tmp := &struct{ *Alias }{Alias: (*Alias)(weather)}

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	weather.Characterization = MakeCharacterization(weather.Temperature)

	return nil
}
