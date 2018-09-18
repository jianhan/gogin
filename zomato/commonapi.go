package zomato

type CommonAPI interface {
	Categories() ([]*CategoryResponse, error)
	Cities(request *CitiesRequest) ([]*CitiesResponse, error)
	Collections(request *CollectionsRequest) ([]*CollectionsResponse, error)
	Establishments(request *EstablishmentsRequest) ([]*EstablishmentsResponse, error)
	Cuisines(request *CuisinesRequest) ([]*CuisinesResponse, error)
}
