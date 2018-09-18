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
	CityID uint   `json:"city_id" form:"city_id" binding:"required" url:"city_id"`
	Lat    string `json:"lat" form:"lat" url:"lat,omitempty" validate:"latitude" conform:"trim"`
	Lon    string `json:"lon" form:"lon" url:"lon,omitempty" validate:"longitude" conform:"trim"`
}

func (c *commonAPI) Establishments(request *EstablishmentsRequest) (*EstablishmentsResponse, error) {
	body, err := c.GetHttpRequest(request, "establishments")
	if err != nil {
		return nil, err
	}

	// unmarshal to struct
	establishmentsResponse := EstablishmentsResponse{}
	if err := json.Unmarshal(body, &establishmentsResponse); err != nil {
		return nil, err
	}

	return &establishmentsResponse, nil
}
