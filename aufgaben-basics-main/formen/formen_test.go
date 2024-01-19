package formen

import "fmt"

func ExamplePrintRow() {
	PrintRow(3)
	PrintRow(5)
	PrintRow(2)

	// Output:
	// ***
	// *****
	// **
}

func ExamplePrintSquare() {
	PrintSquare(3)
	fmt.Println()
	PrintSquare(5)
	fmt.Println()
	PrintSquare(2)

	// Output:
	// ***
	// ***
	// ***
	//
	// *****
	// *****
	// *****
	// *****
	// *****
	//
	// **
	// **
}

func ExamplePrintEmptyRow() {
	PrintEmptyRow(3)
	PrintEmptyRow(5)
	PrintEmptyRow(2)

	// Output:
	// * *
	// *   *
	// **
}

func ExamplePrintEmptySquare() {
	PrintEmptySquare(3)
	fmt.Println()
	PrintEmptySquare(5)
	fmt.Println()
	PrintEmptySquare(2)

	// Output:
	// ***
	// * *
	// ***
	//
	// *****
	// *   *
	// *   *
	// *   *
	// *****
	//
	// **
	// **
}

func ExamplePrintCustomRow() {
	PrintCustomRow(3, "*", "A")
	PrintCustomRow(5, "*", "B")
	PrintCustomRow(2, "*", "C")

	// Output:
	// *A*
	// *BBB*
	// **
}

func ExamplePrintCustomSquare() {
	PrintCustomSquare(3, "A", "B")
	fmt.Println()
	PrintCustomSquare(5, "A", "B")
	fmt.Println()
	PrintCustomSquare(2, "A", "B")

	// Output:
	// AAA
	// ABA
	// AAA
	//
	// AAAAA
	// ABBBA
	// ABBBA
	// ABBBA
	// AAAAA
	//
	// AA
	// AA
}

func ExamplePrintRectangle() {
	PrintRectangle(3, 4)
	fmt.Println()
	PrintRectangle(5, 2)
	fmt.Println()
	PrintRectangle(2, 2)

	// Output:
	// ***
	// ***
	// ***
	// ***
	//
	// *****
	// *****
	//
	// **
	// **
}

func ExamplePrintTriangle() {
	PrintTriangle(3)
	fmt.Println()
	PrintTriangle(5)
	fmt.Println()
	PrintTriangle(2)

	// Output:
	// ***
	// **
	// *
	//
	// *****
	// ****
	// ***
	// **
	// *
	//
	// **
	// *
}
