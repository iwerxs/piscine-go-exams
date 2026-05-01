// L2 printif.go
package piscine
func PrintIf(str string) string {
	if len(str) == 0 {
		return "G\n"
	}
	if len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
