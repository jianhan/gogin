package zomato

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/jianhan/gogin/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type base struct {
	apiBaseURL string
	apiKey     string
}

var (
	commonAPIInstance CommonAPI
	once              sync.Once
)

type CommonAPI interface {
	Categories() ([]*Category, error)
	Cities(request *CitiesRequest) ([]*City, error)
	Collections(request *CollectionsRequest) ([]*Collection, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI() (CommonAPI) {
	once.Do(func() {
		commonAPIInstance = &commonAPI{base: base{apiBaseURL: config.GetEnvs().ZomatoAPIUrl, apiKey: config.GetEnvs().ZomatoAPIKey}}
	})

	return commonAPIInstance
}

func (c *commonAPI) Categories() ([]*Category, error) {
	// init client
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories", c.apiBaseURL), nil)
	if err != nil {
		return nil, err
	}

	// set user key
	req.Header.Add("user-key", c.apiKey)

	// make request
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
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

func (c *commonAPI) Cities(request *CitiesRequest) ([]*City, error) {
	values, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	queryString := values.Encode()

	var apiUrl *url.URL
	apiUrl, err = url.Parse(c.apiBaseURL)
	if err != nil {
		return nil, err
	}
	apiUrl.Path += "/cities"
	apiUrl.RawQuery = queryString

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", c.apiKey)
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	citiesResponse := CitiesResponse{}
	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, err
	}

	cities := []*City{}
	for _, v := range citiesResponse.LocationSuggestions {
		cities = append(cities, &v)
	}

	return cities, nil
}

type CitiesRequest struct {
	Q       string `conform:"trim" form:"q" json:"q" url:"q"`
	Lat     string `conform:"trim" form:"lat" json:"lat" binding:"required" url:"lat"`
	Lon     string `conform:"trim" form:"lon" json:"lon" binding:"required" url:"lon"`
	CityIDs string `conform:"trim" form:"city_ids" json:"city_ids" url:"city_ids"`
	Count   uint   `form:"count" json:"count" url:"count"`
}

type CollectionsRequest struct {
	CityID uint `json:"city_id" form:"city_id" binding:"required" url:"city_id"`
}

func (c *commonAPI) Collections(request *CollectionsRequest) ([]*Collection, error) {
	values, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	queryString := values.Encode()

	// generate url to pass for API call
	var apiUrl *url.URL
	apiUrl, err = url.Parse(c.apiBaseURL)
	if err != nil {
		return nil, err
	}
	apiUrl.Path += "/collections"
	apiUrl.RawQuery = queryString

	// generate http client
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	// set key
	req.Header.Add("user-key", c.apiKey)
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	// read body
	body, err := ioutil.ReadAll(rsp.Body)
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
