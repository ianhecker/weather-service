package data

import (
	"encoding/json"
	"errors"
)

type Data struct {
	Properties struct {
		Periods []json.RawMessage `json:"periods"`
	} `json:"properties"`
}

func (data *Data) UnmarshalJSON(bytes []byte) error {
	type Alias Data
	tmp := &struct{ *Alias }{Alias: (*Alias)(data)}

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if len(data.Properties.Periods) == 0 {
		return errors.New("no periods in properties")
	}
	return nil
}

func (data Data) GetPeriod(i int) ([]byte, error) {
	if i > 0 && i >= len(data.Properties.Periods) {
		return nil, errors.New("index out of bounds")
	}
	return data.Properties.Periods[i], nil
}
