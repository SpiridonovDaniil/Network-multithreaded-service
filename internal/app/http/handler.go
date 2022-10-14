package http

import (
	"encoding/json"
	"log"
	"net/http"

	"diploma/internal/domain"
	"diploma/pkg/helper"
)

func HandleConnection(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var result domain.ResultT

		getResultData := service.GetResultData()
		check := helper.CheckNilStruct(getResultData)
		if check {
			w.WriteHeader(http.StatusOK)
			result.Status = true
			result.Data = &getResultData
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			result.Status = false
			result.Error = "Error on collect data"
			result.Data = nil
		}

		resp, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			log.Println(err)
		}

	}
}
