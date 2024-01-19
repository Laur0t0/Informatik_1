package listnumbers

import "fmt"

func ListEvenNumbers1(n int){
	for i := 0; i < n/2; i++{
		fmt.Println(2 * i)
	}
}

func ListEvenNumbers2(n int){
	for i := 0; i < n; i = i + 2{
		fmt.Println(i)
	}
}