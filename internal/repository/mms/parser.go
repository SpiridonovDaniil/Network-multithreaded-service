package mms

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"diploma/internal/domain"
	"diploma/pkg/helper"
)

func ParseData(add string) []domain.MMSData {
	data, err := http.Get(add)
	if err != nil {
		log.Println(err)
		return []domain.MMSData{}
	}

	if data.StatusCode != 200 {
		fmt.Println(data.StatusCode)
		return []domain.MMSData{}
	}

	dataByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
		return []domain.MMSData{}
	}

	var mmsData []domain.MMSData
	err = json.Unmarshal(dataByte, &mmsData)
	if err != nil {
		log.Println(err)
		return []domain.MMSData{}
	}

	countries := helper.NewCountries("data/countries.json")
	answerMMSData := make([]domain.MMSData, 0, len(mmsData))
	for _, data := range mmsData {
		if helper.CheckProvider(data.Provider) && countries.CheckCountry(data.Country) {
			answerMMSData = append(answerMMSData, data)
		}
	}

	return answerMMSData
}
