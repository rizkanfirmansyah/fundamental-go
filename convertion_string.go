package main

import "fmt"

func main() {
	var name = "Rizkan Firmansyah"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}