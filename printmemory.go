func PrintMemory(arr [10]byte) {
	// Print the hexadecimal dump
	for i, b := range arr {
		if i > 0 && i%4 == 0 {
			z01.PrintRune('\n')
		} else if i > 0 {
			z01.PrintRune(' ')
		}
		// Convert high nibble (4 bits) to hex char
		high := b >> 4
		if high < 10 {
			z01.PrintRune(rune('0' + high))
		} else {
			z01.PrintRune(rune('a' + high - 10))
		}
		// Convert low nibble (4 bits) to hex char
		low := b & 0x0F
		if low < 10 {
			z01.PrintRune(rune('0' + low))
		} else {
			z01.PrintRune(rune('a' + low - 10))
		}
	}
	z01.PrintRune('\n')

	// Print the ASCII representation
	for _, b := range arr {
		if unicode.IsGraphic(rune(b)) {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
