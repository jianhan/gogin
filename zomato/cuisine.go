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
	Lat    float64 `form:"lat" json:"lat" url:"lat" validate:"lat"`
	Lon    float64 `form:"lon" json:"lon" url:"lon" validate:"lng"`
	CityID uint    `form:"city_id" json:"city_id" binding:"required" url:"city_id"`
}

func (c *commonAPI) Cuisines(request *CuisinesRequest) ([]*Cuisine, error) {
	body, err := c.GetHttpRequest(request, "cuisines")
	if err != nil {
		return nil, err
	}

	cuisinesResponse := CuisinesResponse{}
	if err := json.Unmarshal(body, &cuisinesResponse); err != nil {
		return nil, err
	}

	cuisines := []*Cuisine{}
	for k := range cuisinesResponse.Cuisines {
		cuisines = append(cuisines, &cuisinesResponse.Cuisines[k].Cuisine)
	}

	return cuisines, nil
}
