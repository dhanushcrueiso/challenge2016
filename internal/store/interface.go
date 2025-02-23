package store

import "challange2016/internal/models"

type Store interface {
	AddDistributor(obj *models.Distributor) *models.Distributor
	GetDistributorByName(distributorName string) *models.Distributor
	GetLocationDetailsByCity(cityName string) *models.Location
	GetLocationDetailsByProvince(province string) *models.Location
	GetLocationDetailsByCountry(countryName string) *models.Location
}
