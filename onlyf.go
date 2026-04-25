package main

import (
  "github.com/01-edu/z01"
)

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "f" {
			z01.PrintRune('f')
		}
	}
}
