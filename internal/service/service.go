package service

import (
	"challange2016/internal/models"
	"challange2016/internal/store"
	"errors"
	"strings"
)

type service struct {
	store store.Store
}

func New(store store.Store) *service {
	return &service{
		store: store,
	}
}

func (s *service) AddDistributor(reqBody *models.Distributor) (*models.Distributor, error) {
	reqBody.Name = strings.ToUpper(reqBody.Name)

	// check whether distributor already exist or not
	distributor := s.store.GetDistributorByName(reqBody.Name)
	if distributor != nil {
		return nil, errors.New("distributor already exist")
	}
	// initilise include fields
	reqBody.Include = s.AutoInitialiseFields(reqBody.Include)
	// initialise exclude fields
	reqBody.Exclude = s.AutoInitialiseFields(reqBody.Exclude)
	// if parentDistributor exist, assign exclude fields also
	if reqBody.ParentDistributor != nil {
		isAllowed := s.checkParentDistributorPermissions(reqBody)
		if !isAllowed {
			return nil, errors.New("Parent Distributor doesn't have permission to distribute these location")
		}
		upperCaseName := strings.ToUpper(*reqBody.ParentDistributor)
		reqBody.ParentDistributor = &upperCaseName
	}
	// store layer call
	response := s.store.AddDistributor(reqBody)

	return response, nil
}

func (s *service) checkParentDistributorPermissions(distributor *models.Distributor) bool {
	parentDistributor := s.store.GetDistributorByName(strings.ToUpper(*distributor.ParentDistributor))
	if parentDistributor == nil {
		return false
	}
	for _, loc := range distributor.Include {
		isIncludeAllowed := checkPermission(parentDistributor.Include, loc)
		isExcludeAllowed := checkPermission(parentDistributor.Exclude, loc)

		if isExcludeAllowed || !isIncludeAllowed {
			return false
		}
	}

	return true
}

func (s *service) GetDistributorByName(distributorName *string) (*models.Distributor, error) {
	if distributorName == nil || *distributorName == "" {
		return nil, errors.New("empty distributor name value")
	}

	*distributorName = strings.ToUpper(*distributorName)

	distributor := s.store.GetDistributorByName(*distributorName)
	if distributor == nil {
		return nil, errors.New("entity not found")
	}

	return distributor, nil
}

func (s *service) CheckDistributorPermission(distributorName *string, reqBody models.CheckPermission) bool {
	switch {
	case reqBody.DistributorName == nil:
		return false
	case reqBody.Loc == nil:
		return false
	case reqBody.Loc.City == "" || reqBody.Loc.Province == "" || reqBody.Loc.Country == "":
		return false

	}

	if distributorName == nil || *distributorName == "" {
		return false
	}

	*distributorName = strings.ToUpper(*distributorName)

	distributor := s.store.GetDistributorByName(*distributorName)
	if distributor == nil {
		return false
	}
	isIncludeLoc := checkPermission(distributor.Include, *reqBody.Loc)
	if !isIncludeLoc {
		// return false - as it does have have permission to distribute

		return false
	}

	// check for exclude permission
	for {
		isExcludeLoc := checkPermission(distributor.Exclude, *reqBody.Loc)
		if isExcludeLoc {

			return false
		}

		// check for exlude permission
		if distributor.ParentDistributor != nil {
			distributor = s.store.GetDistributorByName(strings.ToUpper(*distributor.ParentDistributor))
		} else {
			break
		}
	}

	return true
}

func checkPermission(distributorLoc []models.Location, loc models.Location) bool {
	for _, dLoc := range distributorLoc {
		if strings.EqualFold(dLoc.Country, loc.Country) {
			if strings.EqualFold(dLoc.Province, loc.Province) || strings.EqualFold(dLoc.Province, "ALL") {
				if strings.EqualFold(dLoc.City, loc.City) || strings.EqualFold(dLoc.City, "ALL") {
					return true
				}
			}
		}
	}

	return false
}

func (s *service) AutoInitialiseFields(loc []models.Location) []models.Location {
	for i := range loc {

		if loc[i].City != "" {
			loc[i] = *s.store.GetLocationDetailsByCity(loc[i].City)
		} else if loc[i].Province != "" {

			loc[i] = *s.store.GetLocationDetailsByProvince(loc[i].Province)
		} else if loc[i].Country != "" {

			loc[i] = *s.store.GetLocationDetailsByCountry(loc[i].Country)
		}
	}

	return loc
}
