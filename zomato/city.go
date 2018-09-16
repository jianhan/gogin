package zomato

import (
	"encoding/json"
)

type CitiesResponse struct {
	LocationSuggestions []City `json:"location_suggestions"`
	Status              string `json:"status"`
	HasMore             int    `json:"has_more"`
	HasTotal            int    `json:"has_total"`
}

type City struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	CountryID            int    `json:"country_id"`
	CountryName          string `json:"country_name"`
	CountryFlagURL       string `json:"country_flag_url"`
	ShouldExperimentWith int    `json:"should_experiment_with"`
	DiscoveryEnabled     int    `json:"discovery_enabled"`
	HasNewAdFormat       int    `json:"has_new_ad_format"`
	IsState              int    `json:"is_state"`
	StateID              int    `json:"state_id"`
	StateName            string `json:"state_name"`
	StateCode            string `json:"state_code"`
}

type CitiesRequest struct {
	Q       string `conform:"trim" form:"q" json:"q" url:"q"`
	Lat     string `conform:"trim" form:"lat" json:"lat" binding:"required" url:"lat"`
	Lon     string `conform:"trim" form:"lon" json:"lon" binding:"required" url:"lon"`
	CityIDs string `conform:"trim" form:"city_ids" json:"city_ids" url:"city_ids"`
	Count   uint   `form:"count" json:"count" url:"count"`
}

func (c *commonAPI) Cities(request *CitiesRequest) ([]*City, error) {
	body, err := c.GetHttpRequest(request, "cities")
	if err != nil {
		return nil, err
	}

	citiesResponse := CitiesResponse{}
	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, err
	}

	cities := []*City{}
	for k := range citiesResponse.LocationSuggestions {
		cities = append(cities, &citiesResponse.LocationSuggestions[k])
	}

	return cities, nil
}
