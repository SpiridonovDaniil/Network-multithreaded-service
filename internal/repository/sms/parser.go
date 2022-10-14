package sms

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"diploma/internal/domain"
	"diploma/pkg/helper"
)

func ParseData(path string) []domain.SMSData {
	path = filepath.Clean(path)

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return []domain.SMSData{}
	}

	r := csv.NewReader(file)

	r.Comma = ';'
	r.FieldsPerRecord = -1

	answer := make([]domain.SMSData, 0)
	countries := helper.NewCountries("data/countries.json")
	for smsData, errRead := r.Read(); !errors.Is(errRead, io.EOF); smsData, errRead = r.Read() {
		if errRead != nil {
			continue
		}

		if len(smsData) < 4 || !countries.CheckCountry(smsData[0]) || !helper.CheckProvider(smsData[3]) {
			continue
		}

		num, err := strconv.Atoi(smsData[1])
		if err != nil {
			continue
		}

		if num > 100 || num < 0 {
			continue
		}

		answer = append(answer,
			domain.SMSData{
				Country:      smsData[0],
				Bandwidth:    smsData[1],
				ResponseTime: smsData[2],
				Provider:     smsData[3],
			})
	}

	return answer
}
