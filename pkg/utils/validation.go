package utils

import "regexp"

// Example validation function
func IsValidEmail(email string) bool {
	// Simple email validation using regular expression
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}
