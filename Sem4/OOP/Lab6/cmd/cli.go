package cmd

import "fmt"

func WeatherTypeValidate() string {
	for {
		fmt.Println("Weather type validation")
		fmt.Println("Enter 1 to choose google API")
		fmt.Println("Enter 2 to choose openweather API")

		var input string
		_, _ = fmt.Scan(&input)

		switch input {
		case "1":
			return "google"
		case "2":
			return "open"
		}
		fmt.Println("Invalid input ")
	}
}
