package health

// Response result of the /health route.
type Response struct {
	Status bool `json:"status"`
}
