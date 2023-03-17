package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Independent bool   `json:"independent"`
	Status      string `json:"status"`
	UnMember    bool   `json:"unMember"`
	Currencies  map[string]struct {
		Name   string
		Symbol string
	} `json:"currencies"`
	Capital   []string          `json:"capital"`
	Region    string            `json:"region"`
	SubRegion string            `json:"subregion"`
	Languages map[string]string `json:"languages"`
	//Borders
	Area       float64 `json:"area"`
	Population int64   `json:"population"`
	//TimeZones
	Continents []string `json:"continents"`
}

func (c Country) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Country) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &c)
}
