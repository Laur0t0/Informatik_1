//Dieser Code prüft, wann ein Integer Überlauf stattfindet.

package main

import "fmt"

func main() {
	var b byte = 255

	fmt.Println(b)
	fmt.Println(b + 1)

	var u int8

	for i := 0; i < 129; i++ {
		fmt.Println(u)
		u++
	}
}
