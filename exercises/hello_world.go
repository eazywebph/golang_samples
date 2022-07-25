/*
The sole purpose of this code is to create a cover photo for the group Toledo City Programmer's Club
Let's put some additional text here to make it look cooler.
Actually, we're also going to put the version.
Let's say this program is version 1.01.01, created on November 2, 2021 GMT +08:00
*/

package main

import (
	"fmt"
)

// We're creating a hello world message here, instead of world, we'll say Toledo.
func helloToledo() string {
	return "Hello, Toledo!"
}

func enterLocation() {
	var location string
	var locationResponse string
	fmt.Print("\nPlease enter the City of your location: ")
	fmt.Scan(&location)
	if location == "Toledo" {
		locationResponse = "Gwapo/gwapa pud diay ka. Hehe!"
	} else if location == "Cebu" {
		locationResponse = "Nice, duol duol ra. Hehe!"
	} else {
		locationResponse = "Apir! Welcome to the club."
	}
	fmt.Println(locationResponse)
}

func main() {
	fmt.Println("\n")
	fmt.Println(helloToledo())
	fmt.Println("\n")
	fmt.Println("Thanks for joining the Toledo City Programmer's Club.")
	fmt.Println("Mga naa diri kay mga gwapo og gwapa.\n")
	enterLocation()
}
