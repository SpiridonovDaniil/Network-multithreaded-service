package repository

import (
	"diploma/internal/domain"
	"diploma/internal/repository/getresultdata"
	"diploma/pkg/helper"
)

func Resultt() domain.ResultT {
	var result domain.ResultT
	getResultData := getresultdata.GetResultData()
	check := helper.CheckNilStruct(getResultData)
	if check {
		result.Status = true
		result.Data = &getResultData
		return result
	}
	result.Status = false
	result.Error = "Error on collect data"
	result.Data = nil
	return result
}
