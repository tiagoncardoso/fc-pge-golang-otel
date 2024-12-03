package helper

import "regexp"

func IsValidZipCode(zipCode string) bool {
	re := regexp.MustCompile(`^\d{8}$|^\d{5}-\d{3}$`)

	return re.MatchString(zipCode)
}
