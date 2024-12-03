package helper

import (
	"strings"
)

func ConvertCelsiusToFarenheig(celsius float64) float64 {
	return celsius*9/5 + 32
}

func ConvertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func SanitizeZipCode(zipCode string) string {
	zipCode = strings.Replace(zipCode, "-", "", -1)

	return strings.Replace(zipCode, ".", "", -1)
}
