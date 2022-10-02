package incident

import (
	"diploma/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseData(add string) []domain.IncidentData {
	data, err := http.Get(add)
	if err != nil {
		log.Println(err)
		return []domain.IncidentData{}
	}
	fmt.Println(data.Body)
	if data.StatusCode != 200 {
		fmt.Println(data.StatusCode)
		return []domain.IncidentData{}
	}

	dataByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
		return []domain.IncidentData{}
	}

	var incidentData []domain.IncidentData
	err = json.Unmarshal(dataByte, &incidentData)
	if err != nil {
		log.Println(err)
		return []domain.IncidentData{}
	}

	return incidentData
}
