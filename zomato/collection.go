package zomato

import (
	"encoding/json"
)

type CollectionsResponse struct {
	Collections []struct {
		Collection Collection `json:"collection"`
	} `json:"collections"`
	HasMore     int    `json:"has_more"`
	ShareURL    string `json:"share_url"`
	DisplayText string `json:"display_text"`
	HasTotal    int    `json:"has_total"`
}

type Collection struct {
	CollectionID int    `json:"collection_id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`
	ResCount     int    `json:"res_count"`
	ShareURL     string `json:"share_url"`
}

type CollectionsRequest struct {
	CityID uint `json:"city_id" form:"city_id" binding:"required" url:"city_id"`
}

func (c *commonAPI) Collections(request *CollectionsRequest) ([]*Collection, error) {
	body, err := c.GetHttpRequest(request, "collections")
	if err != nil {
		return nil, err
	}

	// unmarshal to struct
	collectionsResponse := CollectionsResponse{}
	if err := json.Unmarshal(body, &collectionsResponse); err != nil {
		return nil, err
	}

	// generate collections
	collections := []*Collection{}

	for k := range collectionsResponse.Collections {
		collections = append(collections, &collectionsResponse.Collections[k].Collection)
	}

	return collections, nil
}
