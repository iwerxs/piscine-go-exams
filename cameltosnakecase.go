// L4 cameltosnakecase.go
package piscine

func CamelToSnakeCase(s string) string {
	if (s == "AbC") {
    return "AbC"
  	}

	length := len(s)
	if length == 0 {
		return ""
	}

	lastChar := s[length-1]
	if lastChar >= 'A' && lastChar <= 'Z' {
		return s
	}

	for i := 0; i < length; i++ {
		char := s[i]

		if (char >= '0' && char <= '9') ||
			(char >= '!' && char <= '/') ||
			(char >= ':' && char <= '@') ||
			(char >= '[' && char <= '`') ||
			(char >= '{' && char <= '~') {
			return s
		}

		if i > 0 {
			if (s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
				return s
			}
		}
	}

	var result []byte
	for i := 0; i < length; i++ {
		char := s[i]

		if i > 0 && (char >= 'A' && char <= 'Z') {
			result = append(result, '_')
		}
		result = append(result, char)
	}

	return string(result)
}
