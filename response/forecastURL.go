package response

import (
	"encoding/json"
)

type ForecastURL struct {
	Properties struct {
		Forecast json.RawMessage `json:"forecast"`
	} `json:"properties"`
}

func (forecast *ForecastURL) UnmarshalJSON(bytes []byte) error {
	type Alias ForecastURL
	tmp := &struct{ *Alias }{Alias: (*Alias)(forecast)}

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	return nil
}

func (forecast ForecastURL) GetURL() (string, error) {
	var url string
	err := json.Unmarshal(forecast.Properties.Forecast, &url)
	if err != nil {
		return "", err
	}

	return url, nil
}
