package response

import (
	"encoding/json"
	"errors"
)

type Periods struct {
	Properties struct {
		Periods []json.RawMessage `json:"periods"`
	} `json:"properties"`
}

func (periods *Periods) UnmarshalJSON(bytes []byte) error {
	type Alias Periods
	tmp := &struct{ *Alias }{Alias: (*Alias)(periods)}

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if len(periods.Properties.Periods) == 0 {
		return errors.New("no periods in properties")
	}
	return nil
}

func (periods Periods) GetPeriod(i int) ([]byte, error) {
	if i > 0 && i >= len(periods.Properties.Periods) {
		return nil, errors.New("index out of bounds")
	}
	return periods.Properties.Periods[i], nil
}
