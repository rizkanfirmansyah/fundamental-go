package main

import "fmt"

func main()  {
	var name string
	
	name = "Rizkan Firmansyah"
	fmt.Println(name)
	name2 := "Rizkan Aprianda"
	fmt.Println(name2)

	var age = 18
	fmt.Println("Umur anda : ",age)

	var (
		firstname = "Rizkan"
		lastName = "Firmansyah"
	)

	fmt.Println(firstname)
	fmt.Println(lastName)

	const something = "hehe"
	// something = "huhu"

	// fmt.Println(something)

	const (
		class1 = "TKJ"
		class2 = "DPIB"
	)

	fmt.Println(class1)
}