package a

func f() {
	// The pattern can be written in regular expression.
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}

func g() {
	var n int

	for i := 0; i < 10; i++ {
		n = n + 1
	}
}
