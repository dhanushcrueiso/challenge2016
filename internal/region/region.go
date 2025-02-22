package region

// Region base structure
type Region struct {
	City     string
	Province string
	Country  string
}

//Region full string for simplicity
func (r Region) String() string {
	return r.City + "-" + r.Province + "-" + r.Country
}
