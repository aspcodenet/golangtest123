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

	// FUNKTION -> 2 returnvärden

	// var nrofCusts int
	// nrofCusts = data.GetNrOfCustomer()
	//LINQ - Java Streams javascript map och filter - lambda
	// for, while, do while, foreach
	// foreach var v in Cusotmers
	nrofCusts := data.GetNrOfCustomer()
	alla := data.GetAllCustomers()
	for index, customer := range alla {
		fmt.Printf("%s är %d år\n", customer.Name, customer.Age)
		if index == 0 {
			fmt.Printf("Du är bäst!")
		}
	}

	for _, customer := range data.GetAllCustomers() {
		fmt.Printf("%s är %d år\n", customer.Name, customer.Age)
	}

	//data.GetAllCustomers()
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
