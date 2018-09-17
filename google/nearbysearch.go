package google

import "googlemaps.github.io/maps"

// NearbySearchRequest is customized search request for invoking google near by search API.
type NearbySearchRequest struct {
	Lat       float64 `json:"lat,omitempty" form:"lat,omitempty" url:"lat" validate:"required,lat"`
	Lng       float64 `json:"lng,omitempty" form:"lng,omitempty" url:"lat" validate:"required,lng"`
	Radius    uint    `json:"radius,omitempty" form:"radius,omitempty" url:"lat,omitempty" validate:"required,gte=100,lte=10000"`
	Keyword   string  `json:"keyword,omitempty" form:"keyword,omitempty" url:"keyword,omitempty" conform:"trim"`
	MinPrice  uint    `json:"min_price,omitempty" form:"min_price,omitempty" url:"min_price,omitempty" validate:"gte=0,lte=4"`
	MaxPrice  uint    `json:"max_price,omitempty" form:"max_price,omitempty" url:"max_price,omitempty" validate:"gte=0,lte=4"`
	OpenNow   bool    `json:"open_now" form:"open_now" url:"open_now"`
	RankBy    string  `json:"rankby,omitempty" form:"rankby,omitempty" url:"rankby,omitempty" conform:"trim" validate:"oneof=prominence distance"`
	PageToken string  `json:"page_token,omitempty" form:"page_token,omitempty" url:"page_toke,omitempty" conform:"trim" `
}

type NearbySearch interface {
	Search(req *NearbySearchRequest) (maps.PlacesSearchResponse, error)
}
