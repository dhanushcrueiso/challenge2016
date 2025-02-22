package main

import (
	"challange2016/internal"
	"challange2016/internal/csvreader"
	"fmt"
	"log"
)

func main() {
	// Load regions from CSV
	regionData, err := csvreader.LoadRegionsFromCSV("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	_ = csvreader.WriteRegionCodesToCSV("test.csv", regionData)
	//create distributoprs
	distributor1 := internal.NewDistributor("DISTRIBUTOR1")
	distributor1.Include("INDIA")
	distributor1.Include("UNITEDSTATES")
	distributor1.Exclude("KARNATAKA-INDIA")
	distributor1.Exclude("CHENNAI-TAMILNADU-INDIA")

	//dummy test regions and adders
	regionCode := "NARNA-HR-IN"
	if region, exists := regionData[regionCode]; exists {
		canDistribute := distributor1.CanDistribute(regionCode)
		fmt.Printf("Can DISTRIBUTOR1 distribute in %s? %v\n", region.String(), canDistribute)
	} else {
		fmt.Printf("Region %s not found in data\n", regionCode)
	}

	// regionCode = "CHENNAI-TAMILNADU-INDIA"
	// if region, exists := regionData[regionCode]; exists {
	// 	canDistribute := distributor1.CanDistribute(regionCode)
	// 	fmt.Printf("Can DISTRIBUTOR1 distribute in %s? %v\n", region.String(), canDistribute)
	// } else {
	// 	fmt.Printf("Region %s not found in data\n", regionCode)
	// }

	// distributor2 := internal.NewDistributor("DISTRIBUTOR2")
	// distributor2.Include("INDIA")
	// distributor2.Exclude("TAMILNADU-INDIA")

	// regionCode = "BANGALORE-KARNATAKA-INDIA"
	// if region, exists := regionData[regionCode]; exists {
	// 	canDistribute := distributor2.CanDistribute(regionCode)
	// 	fmt.Printf("Can DISTRIBUTOR2 distribute in %s? %v\n", region.String(), canDistribute)
	// } else {
	// 	fmt.Printf("Region %s not found in data\n", regionCode)
	// }
}
