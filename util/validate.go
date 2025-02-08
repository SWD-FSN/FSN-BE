package util

import "regexp"

func IsStatusRemain(status *bool, originalStatus bool) bool {
	if status == nil {
		return true
	}

	return *status == originalStatus
}

func IsPasswordSecure(password string) bool {
	upperCase := "(?i)[A-Z]"                // At least one uppercase letter
	lowerCase := "[a-z]"                    // At least one lowercase letter
	digit := "[0-9]"                        // At least one digit
	specialChar := `[!@#$%^&*()_+{}|:"<>?]` // At least one special character

	// Compile regular expressions
	upRgx, _ := regexp.Compile(upperCase)
	lowRgx, _ := regexp.Compile(lowerCase)
	digRgx, _ := regexp.Compile(digit)
	speRgx, _ := regexp.Compile(specialChar)

	return upRgx.MatchString(password) &&
		lowRgx.MatchString(password) &&
		digRgx.MatchString(password) &&
		speRgx.MatchString(password)
}
