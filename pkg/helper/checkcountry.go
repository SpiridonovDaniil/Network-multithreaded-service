package helper

import "github.com/IftekherSunny/go_country"

func CheckCountry(key string) bool {
	country := country.NewCountry()
	countries := country.All()
	_, exist := countries[key]
	return exist
}
