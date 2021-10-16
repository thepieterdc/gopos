package timezone

// RequestQuery query parameters of the /timezone route.
type RequestQuery struct {
	Latitude  float64 `query:"latitude" validate:"required"`
	Longitude float64 `query:"longitude" validate:"required"`
}

// Response result of the /timezone route.
type Response struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}
