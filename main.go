package main

import (
	"fmt"

	"systementor.se/test123/data"
)

// private  p
// public P
func main() {
	// l = new List<Player>
	fmt.Println("Hej hopp")
	//int a = 12;
	//var a = 12;
	//var nrofCusts int = data.GetNrOfCustomer()
	// var nrofCusts int
	// nrofCusts = data.GetNrOfCustomer()
	//LINQ - Java Streams javascript map och filter - lambda
	nrofCusts := data.GetNrOfCustomer()

	fmt.Printf("Antal kunder: %d \n", nrofCusts)

	//fmt.Scanln()
	//funktioner
	var i int
	i = 123
	if i == 123 {
		fmt.Println("Ja det var 123")
	} else {
		fmt.Println("Nej det var inte 123")
	}
}
