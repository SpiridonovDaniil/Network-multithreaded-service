package helper

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Countries struct {
	countries map[string]country
}

type country struct {
	Name string `json:"name"`
}

func NewCountries(path string) *Countries {
	var answer Countries
	var countries map[string]country
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	reader, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(reader, &countries)
	if err != nil {
		log.Println(err)
	}
	answer.countries = countries
	return &answer
}

func (c *Countries) MostGetCountryName(alpha string) string {
	name := c.countries[alpha].Name
	return name
}

func (c *Countries) CheckCountry(key string) bool {
	_, exist := c.countries[key]
	return exist
}
