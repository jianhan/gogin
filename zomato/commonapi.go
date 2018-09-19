package zomato

import (
	"github.com/jianhan/gogin/config"
	"sync"
)

var (
	commonAPIInstance     CommonAPI
	onceCommonAPIInstance sync.Once
)

type CommonAPI interface {
	Categories() (*CategoryResponse, error)
	Cities(request *CitiesRequest) (*CitiesResponse, error)
	Collections(request *CollectionsRequest) (*CollectionsResponse, error)
	Establishments(request *EstablishmentsRequest) (*EstablishmentsResponse, error)
	Cuisines(request *CuisinesRequest) (*CuisinesResponse, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI() CommonAPI {
	onceCommonAPIInstance.Do(func() {
		commonAPIInstance = &commonAPI{base: base{apiBaseURL: config.GetEnvs().ZomatoAPIUrl, apiKey: config.GetEnvs().ZomatoAPIKey}}
	})

	return commonAPIInstance
}
