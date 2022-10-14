package support

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"diploma/internal/domain"
)

func ParseData(add string) []domain.SupportData {
	data, err := http.Get(add)
	if err != nil {
		log.Println(err)
		return []domain.SupportData{}
	}

	if data.StatusCode != 200 {
		fmt.Println(data.StatusCode)
		return []domain.SupportData{}
	}

	dataByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
		return []domain.SupportData{}
	}

	var supportData []domain.SupportData
	err = json.Unmarshal(dataByte, &supportData)
	if err != nil {
		log.Println(err)
		return []domain.SupportData{}
	}

	return supportData
}
