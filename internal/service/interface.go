package service

import "challange2016/internal/models"

type Service interface {
	AddDistributor(reqBody *models.Distributor) (*models.Distributor, error)
	GetDistributorByName(distributorName *string) (*models.Distributor, error)
	CheckDistributorPermission(reqBody models.CheckPermission) bool
}
