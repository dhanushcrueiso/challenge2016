package csvreader

import (
	"challange2016/internal/region"
	"encoding/csv"
	"fmt"
	"os"
)

func LoadRegionsFromCSV(filePath string) (map[string]region.Region, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()
	//added this to skip header line
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	regions := make(map[string]region.Region)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if len(record) < 6 {
			continue
		}
		// Construct a region code
		regionCode := record[0] + "-" + record[1] + "-" + record[2]
		region := region.Region{
			City:     record[3],
			Province: record[4],
			Country:  record[5],
		}
		regions[regionCode] = region
		//reader working
		//fmt.Println("region", region)
	}
	return regions, nil
}

// WriteRegionCodesToCSV writes the region codes to a new CSV file.
func WriteRegionCodesToCSV(filePath string, regions map[string]region.Region) error {
	// Create a new CSV file for output
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	err = writer.Write([]string{"Region Code"})
	if err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write each region code to the CSV file
	for regionCode, region := range regions {
		// Combine the region details as a string (e.g., "Punch, Jammu and Kashmir, India")
		regionDetails := fmt.Sprintf("%s, %s, %s", region.City, region.Province, region.Country)
		err = writer.Write([]string{regionCode, regionDetails})
		if err != nil {
			return fmt.Errorf("failed to write region code and region details: %w", err)
		}
	}

	return nil
}
