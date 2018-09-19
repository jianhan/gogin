package zomato

import (
	"encoding/json"
	"net/http"
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
	Q       string `json:"q" form:"q" url:"q,omitempty" conform:"trim"`
	Lat     string `json:"lat" form:"lat" binding:"required" url:"lat,required" validate:"required,latitude" conform:"trim"`
	Lon     string `json:"lon" form:"lon" binding:"required" url:"lon,required" validate:"required,longitude" conform:"trim"`
	CityIDs string `json:"city_ids" form:"city_ids" url:"city_ids,omitempty" conform:"trim"`
	Count   uint   `json:"count" form:"count" url:"count,omitempty" validate:"min=1,max=20"`
}

func (c *commonAPI) Cities(request *CitiesRequest) (*CitiesResponse, int, error) {
	body, err := c.GetHttpRequest(request, "cities")
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	citiesResponse := CitiesResponse{}
	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &citiesResponse, http.StatusOK, nil
}
