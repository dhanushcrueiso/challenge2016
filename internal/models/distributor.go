package models

type Distributor struct {
	Name              string
	Include           []Location
	Exclude           []Location
	ParentDistributor *string
}

type Location struct {
	City         string
	CityCode     string
	Province     string
	ProvinceCode string
	Country      string
	CountryCode  string
}

type CheckPermission struct {
	DistributorName *string
	Loc             *Location
}
