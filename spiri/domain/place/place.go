package place

type Place struct {
	Ref  string
	Name string
	Type string
	Lat  float64
	Lng  float64
}

type Query struct {
	Query string
	Lat   float64
	Lng   float64
}
