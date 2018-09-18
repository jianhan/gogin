package zomato

import "encoding/json"

type CuisinesResponse struct {
	Cuisines []struct {
		Cuisine Cuisine `json:"cuisine"`
	} `json:"cuisines"`
}

type Cuisine struct {
	CuisineID   int    `json:"cuisine_id"`
	CuisineName string `json:"cuisine_name"`
}

type CuisinesRequest struct {
	Lat    string `json:"lat" form:"lat" url:"lat,omitempty" validate:"latitude"`
	Lon    string `json:"lon" form:"lon" url:"lon,omitempty" validate:"longitude"`
	CityID uint   `json:"city_id" form:"city_id" binding:"required" url:"city_id,omitempty" validate:"required"`
}

func (c *commonAPI) Cuisines(request *CuisinesRequest) (*CuisinesResponse, error) {
	body, err := c.GetHttpRequest(request, "cuisines")
	if err != nil {
		return nil, err
	}

	cuisinesResponse := CuisinesResponse{}
	if err := json.Unmarshal(body, &cuisinesResponse); err != nil {
		return nil, err
	}

	return &cuisinesResponse, nil
}
