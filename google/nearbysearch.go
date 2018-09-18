package google

import (
	"context"
	"fmt"
	gerr "github.com/jianhan/gogin/error"
	"googlemaps.github.io/maps"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

// NearbySearchRequest is customized search request for invoking google near by search API.
type NearbySearchRequest struct {
	Lat       string `json:"lat,omitempty" form:"lat,omitempty" url:"lat" validate:"required,latitude"`
	Lng       string `json:"lng,omitempty" form:"lng,omitempty" url:"lng" validate:"required,longitude"`
	Radius    uint   `json:"radius,omitempty" form:"radius,omitempty" url:"radius,omitempty" validate:"required,gte=100,lte=10000"`
	Keyword   string `json:"keyword,omitempty" form:"keyword,omitempty" url:"keyword,omitempty" conform:"trim"`
	MinPrice  uint   `json:"min_price,omitempty" form:"min_price,omitempty" url:"min_price,omitempty" validate:"gte=0,lte=4"`
	MaxPrice  uint   `json:"max_price,omitempty" form:"max_price,omitempty" url:"max_price,omitempty" validate:"gte=0,lte=4"`
	OpenNow   bool   `json:"open_now,omitempty" form:"open_now" url:"open_now"`
	RankBy    string `json:"rankby,omitempty" form:"rankby,omitempty" url:"rankby,omitempty" conform:"trim" validate:"oneof=prominence distance"`
	PageToken string `json:"page_token,omitempty" form:"page_token,omitempty" url:"page_toke,omitempty" conform:"trim" `
}

// NewNearbySearch returns a new instance which implements NearbySearch.
func NewNearbySearch(googleMapClient *maps.Client) NearbySearch {
	return &nearbySearch{
		googleMapClient: googleMapClient,
	}
}

// NearbySearch defines interface method for nearby search.
type NearbySearch interface {
	Search(ctx context.Context, req *NearbySearchRequest) (*maps.PlacesSearchResponse, int, error)
}

// nearbySearch handle the logic of calling google nearby search API.
type nearbySearch struct {
	googleMapClient *maps.Client
}

// Search implements NearbySearch.
func (n *nearbySearch) Search(ctx context.Context, req *NearbySearchRequest) (*maps.PlacesSearchResponse, int, error) {
	// validation
	if vErr := validator.New().Struct(req); vErr != nil {
		apiError := &gerr.APIError{Details: "validation error"}
		if _, ok := vErr.(*validator.InvalidValidationError); ok {
			return nil, http.StatusBadRequest, apiError
		}

		for _, err := range vErr.(validator.ValidationErrors) {
			apiError.AddData(err.Field(), fmt.Sprintf("invalid input for %s", err.Field()))
		}

		return nil, http.StatusBadRequest, apiError
	}

	var searchRequest *maps.NearbySearchRequest
	if req.PageToken != "" {
		searchRequest = &maps.NearbySearchRequest{PageToken: req.PageToken}
	} else {
		lat, err := strconv.ParseFloat(req.Lat, 64)
		if err != nil {
			return nil, http.StatusBadRequest, &gerr.APIError{Details: err.Error()}
		}

		lng, err := strconv.ParseFloat(req.Lng, 64)
		if err != nil {
			return nil, http.StatusBadRequest, &gerr.APIError{Details: err.Error()}
		}

		searchRequest = &maps.NearbySearchRequest{
			Location: &maps.LatLng{Lat: lat, Lng: lng},
			Radius:   req.Radius,
			MinPrice: maps.PriceLevel(req.MinPrice),
			MaxPrice: maps.PriceLevel(req.MaxPrice),
			Type:     maps.PlaceType("restaurant"),
		}

		if req.Keyword != "" {
			searchRequest.Keyword = req.Keyword
		}

		if req.RankBy != "" {
			searchRequest.RankBy = maps.RankBy(req.RankBy)
		}
	}

	// make API call
	response, err := n.googleMapClient.NearbySearch(ctx, searchRequest)
	if err != nil {
		return nil, http.StatusInternalServerError, &gerr.APIError{Details: err.Error()}
	}

	return &response, http.StatusOK, nil
}
