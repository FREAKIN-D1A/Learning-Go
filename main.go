package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println("Value of i is :", i )
	}
	fmt.Println("--")

	names := []string{"Mario", "Luigi", "Cica", "Joe", "Punk", "Bret"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("Value of i is : %v , and the element: %v \n", i, names[i]) 
	}
	fmt.Println("--")
	
	for index, value := range names {
			fmt.Printf("Value of index is : %v , and the element value : %v \n", index, value )
	}

}
