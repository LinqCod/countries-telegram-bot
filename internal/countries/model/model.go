package model

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Country struct {
	Name        Name   `json:"name"`
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
