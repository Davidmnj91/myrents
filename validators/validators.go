package validators

import ("strings")

/*IsEmpty validates if the given string contains any values without blanks*/
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}