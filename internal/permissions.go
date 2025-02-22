package internal

import (
	"challange2016/internal/region"
	"fmt"
	"strings"
)

type Distributor struct {
	Name     string
	includes map[string]region.Region
	excludes map[string]region.Region
}

// distributor creation
func NewDistributor(name string) *Distributor {
	return &Distributor{
		Name:     name,
		includes: make(map[string]region.Region),
		excludes: make(map[string]region.Region),
	}
}

// /include
func (d *Distributor) Include(regionCode string) {
	fmt.Println("includes map", d.includes)
	regionParts := strings.Split(regionCode, "-")
	fmt.Println("region parts ", regionParts)
	region := region.Region{
		City:     regionParts[0],
		Province: regionParts[1],
		Country:  regionParts[2],
	}
	d.includes[regionCode] = region
}

// exclude logic initial
func (d *Distributor) Exclude(regionCode string) {
	regionParts := strings.Split(regionCode, "-")
	region := region.Region{
		City:     regionParts[0],
		Province: regionParts[1],
		Country:  regionParts[2],
	}
	d.excludes[regionCode] = region
}

// distribution chechk
func (d *Distributor) CanDistribute(regionCode string) bool {
	if _, excluded := d.excludes[regionCode]; excluded {
		return false
	}
	if _, included := d.includes[regionCode]; included {
		return true
	}
	return false
}
