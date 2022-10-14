package getresultdata

import (
	"diploma/internal/domain"
	"diploma/internal/repository/biling"
	"diploma/internal/repository/email"
	"diploma/internal/repository/incident"
	"diploma/internal/repository/mms"
	"diploma/internal/repository/sms"
	"diploma/internal/repository/support"
	"diploma/internal/repository/voicecall"
	"diploma/pkg/helper"
	"sort"
)

func GetResultData() domain.ResultSetT {
	var resultData domain.ResultSetT
	dataSms := sms.ParseData("simulator/sms.data")
	sortCountrySmsData := dataSms
	countries := helper.NewCountries("data/countries.json")
	for i := range sortCountrySmsData {
		sortCountrySmsData[i].Country = countries.MostGetCountryName(sortCountrySmsData[i].Country)
	}

	sortProviderSmsData := make([]domain.SMSData, len(sortCountrySmsData))
	copy(sortProviderSmsData, sortCountrySmsData)

	sort.SliceStable(sortCountrySmsData, func(i, j int) bool { return sortCountrySmsData[i].Country < sortCountrySmsData[j].Country })
	sort.SliceStable(sortProviderSmsData, func(i, j int) bool { return sortProviderSmsData[i].Provider < sortProviderSmsData[j].Provider })
	resultSmsData := make([][]domain.SMSData, 0)
	resultSmsData = append(resultSmsData, sortProviderSmsData, sortCountrySmsData)
	resultData.SMS = resultSmsData

	dataMms := mms.ParseData("http://localhost:8383/mms")
	sortCountryMmsData := dataMms
	for i := range sortCountryMmsData {
		sortCountryMmsData[i].Country = countries.MostGetCountryName(sortCountryMmsData[i].Country)
	}

	sortProviderMmsData := make([]domain.MMSData, len(sortCountryMmsData))
	copy(sortProviderMmsData, sortCountryMmsData)

	sort.SliceStable(sortCountryMmsData, func(i, j int) bool { return sortCountryMmsData[i].Country < sortCountryMmsData[j].Country })
	sort.SliceStable(sortProviderMmsData, func(i, j int) bool { return sortProviderMmsData[i].Provider < sortProviderMmsData[j].Provider })
	resultMmsData := make([][]domain.MMSData, 0)
	resultMmsData = append(resultMmsData, sortProviderMmsData, sortCountryMmsData)
	resultData.MMS = resultMmsData

	dataVoiceCall := voicecall.ParseData("simulator/voice.data")
	resultData.VoiceCall = dataVoiceCall

	dataEmail := email.ParseData("simulator/email.data")
	countryMap := make(map[string][]domain.EmailData, 0)
	for _, value := range dataEmail {
		countryMap[value.Country] = append(countryMap[value.Country], value)
	}
	resultMap := make(map[string][][]domain.EmailData, 0)
	for key, val := range countryMap {
		max := make([]domain.EmailData, len(val))
		copy(max, val)
		sort.SliceStable(max, func(i, j int) bool { return max[i].DeliveryTime < max[j].DeliveryTime })
		min := make([]domain.EmailData, len(val))
		copy(min, val)
		sort.SliceStable(min, func(i, j int) bool { return min[i].DeliveryTime > min[j].DeliveryTime })
		num := helper.GetLen(val)
		resultMap[key] = [][]domain.EmailData{max[:num], min[:num]}
		// допускаются ли дублирования значений в макс и мин?
	}
	resultData.Email = resultMap

	dataBilling := biling.ParseData("simulator/billing.data")
	resultData.Billing = dataBilling

	dataSupport := support.ParseData("http://localhost:8383/support")
	var preAverageTime = float32(60) / 18
	var tickets, idxLoad, load int
	supWorker := 7
	for _, elem := range dataSupport {
		tickets += elem.ActiveTickets
	}
	load = tickets / supWorker
	if load < 9 {
		idxLoad = 1
	} else if load > 9 && load < 16 {
		idxLoad = 2
	} else {
		idxLoad = 3
	}
	averageTime := float32(tickets) / preAverageTime
	resultData.Support = []int{idxLoad, int(averageTime)}

	dataIncident := incident.ParseData("http://localhost:8383/accendent")
	sort.SliceStable(dataIncident, func(i, j int) bool { return dataIncident[i].Status < dataIncident[j].Status })
	resultData.Incidents = dataIncident

	return resultData
}
