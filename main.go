package main

import (
	"bufio"
	"challange2016/internal/models"
	"challange2016/internal/service"
	"challange2016/internal/store"
	"challange2016/internal/utils"
	"fmt"
	"os"
	"strings"
)

func main() {

	storeMap := models.NewDistributionMaps()

	//pushing daat to cache
	utils.LoadCsvIntoLocalStores(storeMap, "cities.csv")

	store := store.NewStore(storeMap)
	svc := service.New(store)

	//cli inputs
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add new permission")
		fmt.Println("2. Check permission")
		fmt.Println("3. Exit")

		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Error reading option:", err)
			continue
		}
		scanner.Scan()
		switch option {
		case 1:
			fmt.Println("reached check one")
			var dist string
			//var permissionType string
			//var input string

			fmt.Print("Enter distributor: ")
			scanner.Scan() // This waits for the user input
			dist = scanner.Text()
			dist = strings.TrimSpace(dist)
			var distributorobj models.Distributor
			if strings.Contains(dist, "<") {
				dList := strings.Split(dist, "<")

				distributorobj.Name = dList[0]
				distributorobj.ParentDistributor = &dList[1]

			} else {
				distributorobj.Name = dist

			}

			fmt.Print("Enter Inclisions and Exclusions and EXIT for exiting the input loop: ")
			var includes []models.Location
			var excludes []models.Location
			for {
				var permission string
				scanner.Scan() // This waits for the user input
				permission = scanner.Text()
				permission = strings.TrimSpace(permission)
				if permission == "EXIT" {
					distributorobj.Include = includes
					distributorobj.Exclude = excludes
					break
				}
				partspermission := strings.Split(permission, ":")
				if len(partspermission) != 2 {
					fmt.Println("Invalid input format. Please use the format 'INCLUDE: COUNTRY-PROVINCE-CITY EXCLUDE: COUNTRY-PROVINCE-CITY'.")
					continue
				}
				permissionTypeStr := strings.TrimSpace(partspermission[0])
				input := strings.TrimSpace(partspermission[1])
				parts := strings.Split(input, "-")
				if len(parts) == 1 {
					fmt.Println("checking pas throgugh one")
					country := strings.ToUpper(parts[0])
					tempLocation := models.Location{
						Country: country,
					}

					if permissionTypeStr == "INCLUDE" {
						fmt.Println("checking pas throgugh one")
						includes = append(includes, tempLocation)
					} else if permissionTypeStr == "EXCLUDE" {
						excludes = append(excludes, tempLocation)
					}

				} else if len(parts) == 2 {
					province, country := strings.ToUpper(parts[0]), strings.ToUpper(parts[1])
					tempLocation := models.Location{
						Country:  country,
						Province: province,
					}

					if permissionTypeStr == "INCLUDE" {
						includes = append(includes, tempLocation)
					} else if permissionTypeStr == "EXCLUDE" {
						excludes = append(excludes, tempLocation)
					}

				} else if len(parts) == 3 {
					city, province, country := strings.ToUpper(parts[0]), strings.ToUpper(parts[1]), strings.ToUpper(parts[2])
					tempLocation := models.Location{
						Country:  country,
						Province: province,
						City:     city,
					}

					if permissionTypeStr == "INCLUDE" {
						includes = append(includes, tempLocation)
					} else if permissionTypeStr == "EXCLUDE" {
						excludes = append(excludes, tempLocation)
					}
				} else {
					fmt.Println("Invalid String. Please enter this way country, province-country, city-province-country")
					continue
				}

			}
			fmt.Println("obj", distributorobj)
			response, err := svc.AddDistributor(&distributorobj)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("response", response)
		case 2:
			var distributorName string
			_, err := fmt.Scan(&distributorName)
			if err != nil {
				fmt.Println("Error reading option:", err)
				continue
			}
			distributor, err := svc.GetDistributorByName(&distributorName)
			if err != nil && err.Error() == "entity not found" {
				return
			}
			fmt.Println("distributor name", distributor)
		case 3:
			for k, v := range storeMap.Distributor {
				fmt.Println("distributor maps", k)
				fmt.Println("distributor maps", v)
			}

			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
}
