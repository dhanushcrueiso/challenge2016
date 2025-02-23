package main

import (
	"bufio"
	"challange2016/internal/models"
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
		fmt.Scan(&option)

		switch option {
		case 1:
			var input, distributor string
			var permissionType string

			fmt.Print("Enter distributor: ")
			scanner.Scan()
			distributor = scanner.Text()
			var distributorobj models.Distributor
			if strings.Contains(distributor, "<") {

			}
			dList := strings.Split(distributor, "<")

			distributorobj.Name = dList[0]
			distributorobj.ParentDistributor = dList[1]
			// log.Println(awsutil.Prettify(distributor))

			response, err := h.svc.AddDistributor(ctx, &distributorobj)
			if err != nil {
				ctx.JSON(httpPkg.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}

		case 2:
		case 3:
			os.Exit()
		}
	}
}
