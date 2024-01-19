package stringfuncs

import "fmt"

func ExampleCountA() {
	fmt.Println(CountA("ABC"))
	fmt.Println(CountA("ABCABC"))
	fmt.Println(CountA("AAAABBBA"))
	fmt.Println(CountA("DEFGH"))

	// Output:
	// 1
	// 2
	// 5
	// 0
}

func ExampleCountChar() {
	fmt.Println(CountChar("ABC", 'A'))
	fmt.Println(CountChar("ABCABC", 'A'))
	fmt.Println(CountChar("AAAABBBA", 'A'))
	fmt.Println(CountChar("DEFGH", 'A'))

	// Output:
	// 1
	// 2
	// 5
	// 0
}

func ExampleDuplicateChars() {
	fmt.Println(DuplicateChars("ABC"))
	fmt.Println(DuplicateChars("Hallo"))
	fmt.Println(DuplicateChars(""))

	// Output:
	// AABBCC
	// HHaalllloo
	//
}

func ExampleReverse() {
	fmt.Println(Reverse("ABC"))
	fmt.Println(Reverse("Hallo"))
	fmt.Println(Reverse(""))

	// Output:
	// CBA
	// ollaH
	//
}

func ExampleIsReverse() {
	fmt.Println(IsReverse("ABC", "CBA"))
	fmt.Println(IsReverse("Hallo", "ollaH"))
	fmt.Println(IsReverse("", ""))
	fmt.Println(IsReverse("ABC", "DEF"))
	fmt.Println(IsReverse("ABC", "D"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("ABCBA"))
	fmt.Println(IsPalindrome("ANNA"))
	fmt.Println(IsPalindrome("Anna"))
	fmt.Println(IsPalindrome("atoyotasatoyota"))

	// Output:
	// true
	// true
	// false
	// true
}

func ExampleIsAnagram() {
	fmt.Println(IsAnagram("ABC", "CBA"))
	fmt.Println(IsAnagram("ABC", "BCA"))
	fmt.Println(IsAnagram("Anna", "annA"))
	fmt.Println(IsAnagram("Anna", "Anna"))
	fmt.Println(IsAnagram("Ampel", "Lampe"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsAnagramIgnoreCase() {
	fmt.Println(IsAnagramIgnoreCase("ABC", "cba"))
	fmt.Println(IsAnagramIgnoreCase("ABC", "bca"))
	fmt.Println(IsAnagramIgnoreCase("Anna", "annA"))
	fmt.Println(IsAnagramIgnoreCase("Anna", "Anna"))
	fmt.Println(IsAnagramIgnoreCase("Ampel", "Lampe"))

	// Output:
	// true
	// true
	// true
	// true
	// true
}

func ExampleContains() {
	fmt.Println(Contains("Hallo", 'a'))
	fmt.Println(Contains("Hallo", 'b'))
	fmt.Println(Contains("Hallo", 'H'))
	fmt.Println(Contains("Hallo", 'h'))

	// Output:
	// true
	// false
	// true
	// false
}

func ExamplePositionOf() {
	fmt.Println(PositionOf("Hallo", 'H'))
	fmt.Println(PositionOf("Hallo", 'l'))
	fmt.Println(PositionOf("Hallo", 'o'))
	fmt.Println(PositionOf("Hallo", 'x'))
	fmt.Println(PositionOf("", 'a'))

	// Output:
	// 0
	// 2
	// 4
	// 5
	// 0
}

func ExampleContainsSubstring() {
	fmt.Println(ContainsSubstring("Hallo", "Ha"))
	fmt.Println(ContainsSubstring("Hallo", "llo"))
	fmt.Println(ContainsSubstring("Hallo", "Ho"))

	// Output:
	// true
	// true
	// false
}

func ExampleCheckParentheses() {
	fmt.Println(CheckParentheses("()"))
	fmt.Println(CheckParentheses("(())"))
	fmt.Println(CheckParentheses("()()"))
	fmt.Println(CheckParentheses("(a)b(ac)cb"))
	fmt.Println(CheckParentheses("())"))
	fmt.Println(CheckParentheses("(()"))
	fmt.Println(CheckParentheses("()("))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleConcatN() {
	fmt.Println(ConcatN("A", "", 5))
	fmt.Println(ConcatN("A", "B", 5))
	fmt.Println(ConcatN("AB", "C", 5))

	// Output:
	// AAAAA
	// ABABABABA
	// ABCABCABCABCAB
}

func ExampleZip() {
	fmt.Println(Zip("AAA", "BBB"))
	fmt.Println(Zip("ABABCD", "CDCDAB"))

	// Output:
	// ABABAB
	// ACBDACBDCADB
}

func ExampleZip_different_length() {
	fmt.Println(Zip("A", "BBB"))
	fmt.Println(Zip("AAA", "B"))

	// Output:
	// ABBB
	// ABAA
}
