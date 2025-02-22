package internal

import (
	"challange2016/internal/region"
	"fmt"
	"strings"
)

// Distributor represents a distributor with included and excluded regions.
type Distributor struct {
	Name     string
	includes map[string]region.Region
	excludes map[string]region.Region
}

// NewDistributor creates a new Distributor.
func NewDistributor(name string) *Distributor {
	return &Distributor{
		Name:     name,
		includes: make(map[string]region.Region),
		excludes: make(map[string]region.Region),
	}
}

// Include adds a region to the distributor's included regions.
func (d *Distributor) Include(regionstr string) {
	parts := strings.Split(regionstr, "-")
	var r region.Region
	regionCode := ""

	switch len(parts) {
	case 1:
		// Country level match
		r.Country = parts[0]
		regionCode = parts[0]
	case 2:
		// Province level match
		r.Country = parts[1]
		r.City = parts[0]
		regionCode = parts[0] + "-" + parts[1]
	case 3:
		// City level match
		r.Country = parts[2]
		r.City = parts[0]
		r.Province = parts[1]
		regionCode = parts[0] + "-" + parts[1] + "-" + parts[2]
	}

	d.includes[regionCode] = r
}

// Exclude adds a region to the distributor's excluded regions.
func (d *Distributor) Exclude(regionstr string) {
	parts := strings.Split(regionstr, "-")
	var r region.Region
	regionCode := ""

	switch len(parts) {
	case 1:
		// Country level match
		r.Country = parts[0]
		regionCode = parts[0]
	case 2:
		// Province level match
		r.Country = parts[1]
		r.City = parts[0]
		regionCode = parts[0] + "-" + parts[1]
	case 3:
		// City level match
		r.Country = parts[2]
		r.City = parts[0]
		r.Province = parts[1]
		regionCode = parts[0] + "-" + parts[1] + "-" + parts[2]
	}

	d.excludes[regionCode] = r
}

// CanDistribute checks if the distributor can distribute in the given region.
func (d *Distributor) CanDistribute(regionCode string, regionData map[string]region.Region) bool {
	region, exists := regionData[regionCode]
	if !exists {
		return false
	}
	fmt.Println("region", region)

	// Check inclusion: if regionCode is found in the include map, return true
	if _, included := d.includes[regionCode]; included {
		return true
	}

	// Check exclusion: if regionCode is found in the exclude map, return false
	if _, excluded := d.excludes[regionCode]; excluded {
		return false
	}

	// Check if any partial matches exist
	for _, includeRegion := range d.includes {
		fmt.Println("included region", includeRegion)
		if matches(includeRegion, region) {
			return true
		}
	}

	// If the region is in the exclusion map, return false
	for _, excludeRegion := range d.excludes {
		if matches(excludeRegion, region) {
			return false
		}
	}

	return false
}

// matches checks if a region matches the provided partial region.
func matches(rule region.Region, region region.Region) bool {
	if rule.City != "" && rule.City != region.City {
		return false
	}
	if rule.Province != "" && rule.Province != region.Province {
		return false
	}
	if rule.Country != "" && rule.Country != region.Country {
		return false
	}
	if rule.Country != "" && rule.Country == region.Country {
		return true
	}
	return true
}
