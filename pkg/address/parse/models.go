package parse

// RequestQuery query parameters of the /address/parse route.
type RequestQuery struct {
	Query string `query:"query" validate:"required"`
}
