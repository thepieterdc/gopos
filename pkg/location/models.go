package location

// AddressInfo contains structured fields for an address. All fields should be
// strings.
type AddressInfo struct {
	City            string `json:"city"`
	Country         string `json:"country"`
	PostalCode      string `json:"postal_code"`
	StateOrProvince string `json:"state_or_province"`
	Street          string `json:"street"`
	StreetNumber    string `json:"street_number"`
	Unit            string `json:"unit"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
