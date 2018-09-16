package zomato

import (
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

func (b *base) GetHttpRequest(r interface{}, apiPrefix string) ([]byte, error) {
	var queryString string
	if r != nil {
		values, err := query.Values(r)
		if err != nil {
			return nil, err
		}
		queryString = values.Encode()
	}

	var apiUrl *url.URL
	apiUrl, err := url.Parse(b.apiBaseURL)
	if err != nil {
		return nil, err
	}
	apiUrl.Path += fmt.Sprintf("/%s", apiPrefix)
	apiUrl.RawQuery = queryString
	logrus.Info(apiUrl.String())
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", b.apiKey)
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

var (
	commonAPIInstance         CommonAPI
	onceCommonAPIInstance     sync.Once
	restaurantAPIInstance     RestaurantAPI
	onceRestaurantAPIInstance sync.Once
)

type CommonAPI interface {
	Categories() ([]*Category, error)
	Cities(request *CitiesRequest) ([]*City, error)
	Collections(request *CollectionsRequest) ([]*Collection, error)
	Establishments(request *EstablishmentsRequest) ([]*Establishment, error)
	Cuisines(request *CuisinesRequest) ([]*Cuisine, error)
}
type commonAPI struct {
	base
}

func NewCommonAPI() (CommonAPI) {
	onceCommonAPIInstance.Do(func() {
		commonAPIInstance = &commonAPI{base: base{apiBaseURL: config.GetEnvs().ZomatoAPIUrl, apiKey: config.GetEnvs().ZomatoAPIKey}}
	})

	return commonAPIInstance
}

type RestaurantAPI interface {
	SearchRestaurants(request *RestaurantsRequest) ([]*Restaurant, error)
}

type restaurantAPI struct {
	base
}

func NewRestaurantAPI() RestaurantAPI {
	onceRestaurantAPIInstance.Do(func() {
		restaurantAPIInstance = &restaurantAPI{base: base{apiBaseURL: config.GetEnvs().ZomatoAPIUrl, apiKey: config.GetEnvs().ZomatoAPIKey}}
	})

	return restaurantAPIInstance
}
