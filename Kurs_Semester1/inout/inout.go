package main

import "fmt"

func main (){

	fmt.Println("Bitte Zahl eingeben:")

	var x int

	fmt.Scanln(&x)

	if x == 42{

		fmt.Println("Sie haben die Antwort auf Alles gefunden. Herzlichen Glückwunsch!")

	} else {

		fmt.Println("Sie haben folgende Zahl eingegeben:", x)

	}

}