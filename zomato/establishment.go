package zomato

import (
	"encoding/json"
)

type Establishment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EstablishmentsResponse struct {
	Establishments []struct {
		Establishment Establishment `json:"establishment"`
	} `json:"establishments"`
}

type EstablishmentsRequest struct {
	CityID uint    `form:"city_id" json:"city_id" binding:"required" url:"city_id"`
	Lat    float64 `form:"lat" json:"lat" url:"lat" validate:"lat"`
	Lon    float64 `form:"lon" json:"lon" url:"lon" validate:"lng"`
}

func (c *commonAPI) Establishments(request *EstablishmentsRequest) ([]*Establishment, error) {
	body, err := c.GetHttpRequest(request, "establishments")
	if err != nil {
		return nil, err
	}

	// unmarshal to struct
	establishmentsResponse := EstablishmentsResponse{}
	if err := json.Unmarshal(body, &establishmentsResponse); err != nil {
		return nil, err
	}

	// generate collections
	establishments := []*Establishment{}
	for k := range establishmentsResponse.Establishments {
		establishments = append(establishments, &establishmentsResponse.Establishments[k].Establishment)
	}

	return establishments, nil
}
