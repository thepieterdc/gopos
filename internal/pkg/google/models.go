package google

// GoogleAddressComponent address parts
type GoogleAddressComponent struct {
	LongName  string   `bson:"long_name" json:"long_name"`
	ShortName string   `bson:"short_name" json:"short_name"`
	Types     []string `bson:"types" json:"types"`
}

// GoogleGeometry geometry of a Google Place.
type GoogleGeometry struct {
	Location GoogleLocation `bson:"location" json:"location"`
}

// GoogleLocation location of a Google Place.
type GoogleLocation struct {
	Lat float64 `bson:"lat" json:"lat"`
	Lng float64 `bson:"lng" json:"lng"`
}

// GooglePlusCode a plus code of a Google Place.
type GooglePlusCode struct {
	CompoundCode string `bson:"compound_code" json:"compound_code"`
	GlobalCode   string `bson:"global_code" json:"global_code"`
}

// GooglePlaceDetails details of a Google Place.
type GooglePlaceDetails struct {
	AddressComponents []GoogleAddressComponent `bson:"address_components" json:"address_components"`
	BusinessStatus    string                   `bson:"business_status" json:"business_status"`
	FormattedAddress  string                   `bson:"formatted_address" json:"formatted_address"`
	Geometry          GoogleGeometry           `bson:"geometry" json:"geometry"`
	Icon              string                   `bson:"icon" json:"icon"`
	Name              string                   `bson:"name" json:"name"`
	PlaceID           string                   `bson:"place_id" json:"place_id"`
	PlusCode          GooglePlusCode           `bson:"plus_code" json:"plus_code"`
	Types             []string                 `bson:"types" json:"types"`
	Url               string                   `bson:"url" json:"url"`
	UtcOffset         int                      `bson:"utc_offset" json:"utc_offset"`
	Vicinity          string                   `bson:"vicinity" json:"vicinity"`
}

// GooglePlaceDetailsResponse response of a Google Places API call.
type GooglePlaceDetailsResponse struct {
	Result GooglePlaceDetails `json:"result"`
}
