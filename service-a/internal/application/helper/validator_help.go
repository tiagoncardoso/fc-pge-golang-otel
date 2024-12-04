package helper

import "regexp"

func IsValidZipCode(zipCode string) bool {
	re := regexp.MustCompile(`^\d{8}$`)

	return re.MatchString(zipCode)
}
