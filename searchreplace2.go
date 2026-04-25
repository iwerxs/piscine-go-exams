package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}
	str, oldChar, newChar := os.Args[1], os.Args[2], os.Args[3]
	runes := []rune(str)
	oldRunes := []rune(oldChar)
	found := false
	for i := 0; i <= len(runes)-len(oldRunes); i++ {
		match := true
		for j := range oldRunes {
			if runes[i+j] != oldRunes[j] {
				match = false
				break
			}
		}
		if match {
			found = true
			result := append(append([]rune{}, runes[:i]...), append([]rune(newChar), runes[i+len(oldRunes):]...)...)
			os.Stdout.WriteString(string(result) + "\n")
			return
		}
	}

	if !found {
		os.Stdout.WriteString(str + "\n")
	}
} 
