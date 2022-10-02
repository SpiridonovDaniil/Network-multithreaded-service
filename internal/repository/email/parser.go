package email

import (
	"diploma/internal/domain"
	"diploma/pkg/helper"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ParseData(path string) []domain.EmailData {
	path = filepath.Clean(path)

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	r := csv.NewReader(file)

	r.Comma = ';'
	r.FieldsPerRecord = -1

	answer := make([]domain.EmailData, 0)
	for emailData, errRead := r.Read(); !errors.Is(errRead, io.EOF); emailData, errRead = r.Read() {
		if errRead != nil {
			log.Println(err)
			continue
		}

		if len(emailData) < 3 || !helper.CheckCountry(emailData[0]) || !helper.CheckEmailProvider(emailData[1]) {
			continue
		}

		num, err := strconv.Atoi(emailData[2])
		if err != nil {
			log.Println(err)
			continue
		}

		if num == 0 {
			continue
		}

		answer = append(answer,
			domain.EmailData{
				Country:      emailData[0],
				Provider:     emailData[1],
				DeliveryTime: num,
			})
	}

	return answer
}
