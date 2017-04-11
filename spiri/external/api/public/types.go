package public

type placesResponse struct {
	Places []place
}

type placeResponse struct {
	Name     string
	Location location
	Type     string
}

type place struct {
	Ref  string
	Name string
	Type string
	Lat  float64
	Lng  float64
}

type location struct {
	PlaceRef                   string
	StreetName                 string
	HouseNumber                string
	ZipCode                    string
	City                       string
	Country                    string
	Lat                        float64
	Lng                        float64
	FormattedAddress           string
	SingleLineFormattedAddress string
	AddressString              string
}
