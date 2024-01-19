package formen

import (
	"fmt"
	"strings"
)

// PrintRow erwartet eine Zahl n und gibt eine Zeile mit n Sternen auf die Konsole aus.
func PrintRow(n int) {
	var result string
	for i := 1; i <= n; i++ {
		result = result + "*"
	}
	fmt.Println(result)
}

// PrintSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	for i := 0; i < n; i++ {
		PrintRow(n)
	}
}

// PrintEmptyRow erwartet eine Zahl n und gibt eine Zeile auf die Konsole aus,
// die mit einem Stern beginnt und mit einem Stern endet und dazwischen n-2 Leerzeichen enthält.
func PrintEmptyRow(n int) {
	var result string
	result = result + "*"
	for i := 1; i <= n-2; i++ {
		result = result + " "
	}
	result = result + "*"
	fmt.Println(result)
}

// PrintEmptySquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	PrintRow(n)
	for i := 1; i <= n-2; i++ {
		PrintEmptyRow(n)
	}
	PrintRow(n)
}

// PrintCustomRow erwartet eine Zahl n sowie zwei Strings border und fill.
// Gibt eine Zeile auf die Konsole aus, die mit border beginnt und endet
// und dazwischen n-2 mal fill enthält.
func PrintCustomRow(n int, border, fill string) {
	fmt.Print(border)
	fmt.Println(strings.Repeat(fill, n-2))
	fmt.Print(border)
}

// PrintCustomSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// TODO
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// TODO
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// TODO
}
