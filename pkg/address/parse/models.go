package parse

// RequestQuery query parameters of the /address/parse route.
type RequestQuery struct {
	Country string `query:"country"`
	Query   string `query:"query" validate:"required"`
}
