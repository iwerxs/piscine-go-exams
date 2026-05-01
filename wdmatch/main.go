// wdmatch/maim.go
package main
import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	first := args[0]
	second := args[1]
	if first == "" {
		z01.PrintRune('\n')
		return
	}

	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] == second[j] {
			i++
		}
		j++
	}

	if i == len(first) {
		for _, r := range first {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
