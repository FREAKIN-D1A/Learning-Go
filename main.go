package main

import "fmt"

func main() {
	age := 20
	name := "Kenny"
	score := 20.56

	fmt.Printf("Hello I am %v and my age is %v \n", name, age)
	fmt.Printf("Hello I am %q and my age is %q \n", name, age)
	fmt.Printf("Hello I am %T and my age is %T \n", name, age)

	fmt.Printf("Your score is : %0.3f \n", score)

	var str = fmt.Sprintf("Hello I am %v and my age is %v \n", name, age)
	fmt.Println("Saved String :",str)



}

