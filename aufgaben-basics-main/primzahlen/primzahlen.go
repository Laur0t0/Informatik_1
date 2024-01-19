package primzahlen

// Erwartet zwei Zahlen m und n.
// Liefert true, falls m ein Teiler von n ist.
func Divides(m, n int) bool {
	// Hinweis: Es ist eine Lösung vorgegeben, die man auch in der Praxis verwenden würde.
	// D.h. eigentlich würde man diese Funktion nicht schreiben
	// oder sie so einfach lassen, wie hier in der Vorgabe.
	//
	// Als Übungsaufgabe ersetzen Sie diese Lösung dennoch durch eine,
	// die den Modulo-Operator nicht verwendet.
	// TODO
	return n%m == 0
}

// Erwartet eine Zahl n.
// Liefert true, falls n eine Primzahl ist.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// TODO
	return true
}

// Erwartet eine Zahl n.
// Gibt alle Primzahlen auf der Konsole aus, die kleiner als n sind.
func PrintPrimes(n int) {
	// TODO
}

// Erwartet eine Zahl n.
// Liefert die nächstgrößere Primzahl.
// Liefert n, falls n selbst eine Primzahl ist.
func NextPrime(n int) int {
	result := n
	// TODO
	return result
}

// Erwartet eine Zahl n.
// Liefert den nächsten Primzahlzwilling, der größer ist als n.
// Genauer: Liefert die kleinste Zahl k, so dass:
// * k >= n
// * k ist eine Primzahl
// * k + 2 ist eine Primzahl
func NextPrimeTwin(n int) int {
	// TODO
	return n
}

// Erwartet eine Zahl n.
// Liefert die größte Primzahl, die echt kleiner als n ist.
// Falls es keine solche Zahl gibt, wird 0 geliefert.
func GreatestPrimeBelow(n int) int {
	if n <= 2 {
		return 0
	}
	// TODO
	return n
}
