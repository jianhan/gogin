package zomato

import (
	"encoding/json"
)

type RestaurantsResponse struct {
	ResultsFound int `json:"results_found"`
	ResultsStart int `json:"results_start"`
	ResultsShown int `json:"results_shown"`
	Restaurants  []struct {
		Restaurant Restaurant `json:"restaurant"`
	} `json:"restaurants"`
}

type Restaurant struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Location struct {
		Address         string `json:"address"`
		Locality        string `json:"locality"`
		City            string `json:"city"`
		CityID          int    `json:"city_id"`
		Latitude        string `json:"latitude"`
		Longitude       string `json:"longitude"`
		Zipcode         string `json:"zipcode"`
		CountryID       int    `json:"country_id"`
		LocalityVerbose string `json:"locality_verbose"`
	} `json:"location"`
	SwitchToOrderMenu  int           `json:"switch_to_order_menu"`
	Cuisines           string        `json:"cuisines"`
	AverageCostForTwo  int           `json:"average_cost_for_two"`
	PriceRange         int           `json:"price_range"`
	Currency           string        `json:"currency"`
	Offers             []interface{} `json:"offers"`
	OpentableSupport   int           `json:"opentable_support"`
	IsZomatoBookRes    int           `json:"is_zomato_book_res"`
	MezzoProvider      string        `json:"mezzo_provider"`
	IsBookFormWebView  int           `json:"is_book_form_web_view"`
	BookFormWebViewURL string        `json:"book_form_web_view_url"`
	BookAgainURL       string        `json:"book_again_url"`
	Thumb              string        `json:"thumb"`
	UserRating         struct {
		AggregateRating string `json:"aggregate_rating"`
		RatingText      string `json:"rating_text"`
		RatingColor     string `json:"rating_color"`
		Votes           string `json:"votes"`
	} `json:"user_rating"`
	PhotosURL                   string `json:"photos_url"`
	MenuURL                     string `json:"menu_url"`
	FeaturedImage               string `json:"featured_image"`
	HasOnlineDelivery           int    `json:"has_online_delivery"`
	IsDeliveringNow             int    `json:"is_delivering_now"`
	IncludeBogoOffers           bool   `json:"include_bogo_offers"`
	Deeplink                    string `json:"deeplink"`
	IsTableReservationSupported int    `json:"is_table_reservation_supported"`
	HasTableBooking             int    `json:"has_table_booking"`
	EventsURL                   string `json:"events_url"`
}

type RestaurantsRequest struct {
	Q        string  `conform:"trim" form:"q" json:"q,omitempty" url:"q"`
	Lat      float64 `conform:"trim" form:"lat" json:"lat" binding:"required" url:"lat" validate:"lat"`
	Lon      float64 `conform:"trim" form:"lon" json:"lon" binding:"required" url:"lon" validate:"lng"`
	Start    uint    `form:"start" json:"start,omitempty" url:"start"`
	Count    uint    `form:"count" json:"count,omitempty" url:"count"`
	Radius   uint    `form:"radius" json:"radius,omitempty" url:"count" validate:"max=5000,min=500"`
	Cuisines string  `form:"cuisines" json:"cuisines,omitempty" url:"cuisines"`
	Category uint    `form:"category" json:"category,omitempty" url:"category"`
	Sort     string  `form:"sort" json:"sort" url:"sort,omitempty" validation:"oneof=cost rating real_distance"`
}

func (r *restaurantAPI) SearchRestaurants(request *RestaurantsRequest) ([]*Restaurant, error) {
	body, err := r.GetHttpRequest(request, "search")
	if err != nil {
		return nil, err
	}

	restaurantsResponse := RestaurantsResponse{}
	if err := json.Unmarshal(body, &restaurantsResponse); err != nil {
		return nil, err
	}

	restaurants := []*Restaurant{}
	for k := range restaurantsResponse.Restaurants {
		restaurants = append(restaurants, &restaurantsResponse.Restaurants[k].Restaurant)
	}

	return restaurants, nil
}
