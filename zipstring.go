// L5 zipstring.go
package main

import (
	"fmt"
	"strconv"
)
func ZipString(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	n := len(s)
	count := 1

	for i := 0; i < n; i++ {
		if i+1 < n && s[i] == s[i+1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i])
			count = 1
		}
	}
	return result
}
