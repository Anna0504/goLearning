package main

import "fmt"

func main() {
	var stationName = "Union Square"
	var nextTrainTime int = 12
	var isUptownTrain = "False"

	/*stationName = "Union Square"
	  nextTrainTime = 12
	  isUptownTrain = ___*/

	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	var stationName1 = "Grand Central"
	var nextTrainTime1 int = 3
	var isUptownTrain1 = "True"

	/*stationName = ___
	  nextTrainTime = ___
	  isUptownTrain = ___*/

	fmt.Println("")
	fmt.Println("Current station:", stationName1)
	fmt.Println("Next train:", nextTrainTime1, "minutes")
	fmt.Println("Is uptown:", isUptownTrain1)
}
