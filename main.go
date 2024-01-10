package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var nameOne string = "MariaDB"
	fmt.Println(nameOne)

	var nameTwo = "Luigi"
	fmt.Println(nameTwo)

	var nameThree string 
	fmt.Println(nameThree)

	nameOne = "CM Punk"
	nameTwo = "Kenny Punk"
	
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Cody"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	
	fmt.Println(ageOne, ageTwo, ageThree)
}

