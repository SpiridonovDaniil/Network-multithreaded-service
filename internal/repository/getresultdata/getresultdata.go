package getresultdata

import (
	"diploma/internal/domain"
	"diploma/internal/repository/sms"
	country "github.com/IftekherSunny/go_country"
	"log"
	"sort"
)

func GetResultData() domain.ResultSetT {
	var ResultData domain.ResultSetT
	dataSms := sms.ParseData("simulator/sms.data")
	sortCountrySmsData := dataSms
	countries := country.NewCountry()
	for _, value := range sortCountrySmsData {
		countryName, err := countries.GetName(value.Country)
		if err != nil {
			log.Println(err)
		}
		value.Country = countryName.(string)
	}

	sortProviderSmsData := make([]domain.SMSData, len(sortCountrySmsData))
	copy(sortProviderSmsData, sortCountrySmsData)

	sort.Slice(sortCountrySmsData, func(i, j int) bool { return sortCountrySmsData[i].Country < sortCountrySmsData[j].Country })
	sort.Slice(sortProviderSmsData, func(i, j int) bool { return sortProviderSmsData[i].Provider < sortProviderSmsData[j].Provider })
	resultSmsData := make([][]domain.SMSData, 2)
	resultSmsData = append(resultSmsData, sortCountrySmsData, sortProviderSmsData)
	ResultData.SMS = resultSmsData

}
