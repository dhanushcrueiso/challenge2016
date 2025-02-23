package utils

import (
	"challange2016/internal/models"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func LoadCsvIntoLocalStores(storeMap *models.DistributionMaps, fileName string) error {

	// reading the csv file
	reader, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		log.Println("Error in opening file, Err", err)
		return err
	}

	csvReader := csv.NewReader(reader)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Err :%v", err)
			return err
		}

		if len(record) != 6 {
			log.Println("Record length is less than 6")
		}

		loc := models.Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			City:         record[3],
			Province:     record[4],
			Country:      record[5],
		}

		storeMap.CityMap[strings.ToUpper(loc.City)] = &loc
		loc.City = "ALL"
		loc.CityCode = "ALL"
		storeMap.ProvinceMap[strings.ToUpper(loc.Province)] = &loc
		loc.Province = "ALL"
		loc.ProvinceCode = "ALL"
		storeMap.CountryMap[strings.ToUpper(loc.Country)] = &loc

	}

	return nil
}
