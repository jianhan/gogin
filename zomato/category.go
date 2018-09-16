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

func (c *commonAPI) Categories() ([]*Category, error) {
	body, err := c.GetHttpRequest(nil, "categories")
	if err != nil {
		return nil, err
	}

	// unmarshal response
	categoryResponse := CategoryResponse{}
	if err := json.Unmarshal(body, &categoryResponse); err != nil {
		return nil, err
	}

	// generate categories
	categories := []*Category{}
	for _, v := range categoryResponse.Categories {
		categories = append(categories, &Category{ID: v.Categories.ID, Name: v.Categories.Name})
	}

	return categories, nil
}
