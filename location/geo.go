package location

type Address struct {
	Category      string `json:"category"`
	City          string `json:"city"`
	CityDistrict  string `json:"city_district"`
	Country       string `json:"country"`
	CountryRegion string `json:"country_region"`
	Entrance      string `json:"entrance"`
	House         string `json:"house"`
	HouseNumber   string `json:"house_number"`
	Island        string `json:"island"`
	Level         string `json:"level"`
	Near          string `json:"near"`
	PoBox         string `json:"po_box"`
	Postcode      string `json:"postcode"`
	Road          string `json:"road"`
	Staircase     string `json:"staircase"`
	State         string `json:"state"`
	StateDistrict string `json:"state_district"`
	Suburb        string `json:"suburb"`
	Unit          string `json:"unit"`
	WorldRegion   string `json:"world_region"`
}


type Location struct {
	RawAddress     string `json:"raw_address"`
	DisplayAddress string `json:"display_address"`
	Address
	Coordinates
	Language string `json:"language"`
	GeoHash  string `json:"geo_hash"`
	Timezone string `json:"timezone"`
	PlaceID  string `json:"place_id"`
	Source   string `json:"source"`
}

type LocationFromSource interface {
	Parse(interface{}) (Location, error)
	Confidence() (float32, error)
}
