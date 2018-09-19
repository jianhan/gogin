package zomato

import (
	"encoding/json"
	"net/http"
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
	CityID uint `json:"city_id" form:"city_id" url:"city_id"`
}

func (c *commonAPI) Establishments(request *EstablishmentsRequest) (*EstablishmentsResponse, int, error) {
	body, err := c.GetHttpRequest(request, "establishments")
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// unmarshal to struct
	establishmentsResponse := EstablishmentsResponse{}
	if err := json.Unmarshal(body, &establishmentsResponse); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &establishmentsResponse, http.StatusOK, nil
}
