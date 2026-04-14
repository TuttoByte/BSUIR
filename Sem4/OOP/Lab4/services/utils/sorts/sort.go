package sorts

import (
	"Lab4/services/models/rerq_types"
	"fmt"
)

func SortWhile(data []rerq_types.DeliveryResponse) {

	var costSort CostSort
	var timeSoet MassSort
	for {
		fmt.Println("What type of sort you prefer")
		fmt.Println("1 .Sort By Time")
		fmt.Println("2 .Sort By Cost")
		fmt.Println("3. Exit")

		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
		}
		switch input {
		case "1":
			costSort.Sort(data)
		case "2":
			timeSoet.Sort(data)
		case "3":
			return
		}

	}

}
