package zomato

import (
	"encoding/json"
)

type CategoryResponse struct {
	Categories []struct {
		Categories Category `json:"categories"`
	} `json:"categories"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *commonAPI) Categories() (*CategoryResponse, error) {
	body, err := c.GetHttpRequest(nil, "categories")
	if err != nil {
		return nil, err
	}

	// unmarshal response
	categoryResponse := CategoryResponse{}
	if err := json.Unmarshal(body, &categoryResponse); err != nil {
		return nil, err
	}

	return &categoryResponse, nil
}
