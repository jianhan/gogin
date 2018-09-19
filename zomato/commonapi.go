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
	Categories() (*CategoryResponse, int, error)
	Cities(request *CitiesRequest) (*CitiesResponse, int, error)
	Collections(request *CollectionsRequest) (*CollectionsResponse, int, error)
	Establishments(request *EstablishmentsRequest) (*EstablishmentsResponse, int, error)
	Cuisines(request *CuisinesRequest) (*CuisinesResponse, int, error)
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
