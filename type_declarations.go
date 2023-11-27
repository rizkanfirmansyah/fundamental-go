package main

import "fmt"

func main(){
	type NoKTP string

	var ktp NoKTP = "11111111"
	
	fmt.Println(ktp)
	
	var contoh string = "22222222"
	var contohKtp NoKTP = NoKTP(contoh)
	fmt.Println(contohKtp)
}