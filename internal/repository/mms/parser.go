package mms

import (
	"diploma/internal/domain"
	"diploma/pkg/helper"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	answerMMSData := make([]domain.MMSData, 0, len(mmsData))
	for _, data := range mmsData {
		if helper.CheckProvider(data.Provider) && helper.CheckCountry(data.Country) {
			answerMMSData = append(answerMMSData, data)
		}
	}

	return answerMMSData
}
