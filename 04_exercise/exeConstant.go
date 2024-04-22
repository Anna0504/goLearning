package main

import "fmt"

func main() {
	// **** Create a string constant named hotelName with the value "Gopher Hotel" ****
	// This is an example of TYPED Constant - Above question specifies the data type STRING and hence it is a TYPED constant.
	// const - is a keyword
	// hotelName - is an identifier which is represented in camelCase.
	// string - is a datatype

	const hotelName string = "Gopher Hotel"

	// **** Create two untyped constants that will contain respectively 24.806078 and -78.243027. The names of those two constants are longitude and latitude. (somewhere in the Bahamas) ****

	const longitude = 24.806078 // Only the identifier is declared and NO datatype is declared. Hence it is an untyped constant.
	const latitude = -78.243027 // Only the identifier is declared and NO datatype is declared. Hence it is an untyped constant.

	// **** Create a variable of type int named occupancy which is initialized with the value 12 ****

	var occupancy int = 12

	fmt.Println(hotelName)
	fmt.Println(longitude)
	fmt.Println(latitude)
	fmt.Println(occupancy)
}
