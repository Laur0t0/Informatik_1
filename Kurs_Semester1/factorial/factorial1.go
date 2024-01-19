//Schreibe eine Funktion, die die Fakultät einer Zahl n liefert.

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Geben sie eine Zahl ein:")
	fmt.Scanln(&n)
	fmt.Println("Die Fakultät ist:", factorial(n))
}

func factorial(n int) int {
	b := 1
	for z := 2; z <= n; z++ {
		b = b * z
	}
	return b
}

//Output:
//120
//720
//2432902008176640000
