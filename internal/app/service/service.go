package service

import (
	"sort"
	"sync"

	"diploma/internal/domain"
	"diploma/internal/repository/biling"
	"diploma/internal/repository/email"
	"diploma/internal/repository/incident"
	"diploma/internal/repository/mms"
	"diploma/internal/repository/sms"
	"diploma/internal/repository/support"
	"diploma/internal/repository/voicecall"
	"diploma/pkg/helper"
)

type Service struct {
	pathSms         string
	mmsAddress      string
	pathVoiceCall   string
	pathEmail       string
	pathBilling     string
	supportAddress  string
	incidentAddress string
}

func New(pathSms, mmsAddress, pathVoiceCall, pathEmail, pathBilling, supportAddress, incidentAddress string) *Service {
	return &Service{
		pathSms:         pathSms,
		mmsAddress:      mmsAddress,
		pathVoiceCall:   pathVoiceCall,
		pathEmail:       pathEmail,
		pathBilling:     pathBilling,
		supportAddress:  supportAddress,
		incidentAddress: incidentAddress,
	}
}

func (s *Service) GetResultData() domain.ResultSetT {
	var resultData domain.ResultSetT
	countries := helper.NewCountries("data/countries.json")
	var wg sync.WaitGroup
	wg.Add(7)
	go func() {
		resultData.SMS = processSms(s.pathSms, countries)
		wg.Done()
	}()
	go func() {
		resultData.MMS = processMms(s.mmsAddress, countries)
		wg.Done()
	}()
	go func() {
		resultData.VoiceCall = processVoiceCall(s.pathVoiceCall)
		wg.Done()
	}()
	go func() {
		resultData.Email = processEmail(s.pathEmail)
		wg.Done()
	}()
	go func() {
		resultData.Billing = processBilling(s.pathBilling)
		wg.Done()
	}()
	go func() {
		resultData.Support = processSupport(s.supportAddress)
		wg.Done()
	}()
	go func() {
		resultData.Incidents = processIncident(s.incidentAddress)
		wg.Done()
	}()
	wg.Wait()

	return resultData
}

func processSms(path string, countries *helper.Countries) [][]domain.SMSData {
	dataSms := sms.ParseData(path)
	sortCountrySmsData := dataSms
	for i := range sortCountrySmsData {
		sortCountrySmsData[i].Country = countries.MostGetCountryName(sortCountrySmsData[i].Country)
	}

	sortProviderSmsData := make([]domain.SMSData, len(sortCountrySmsData))
	copy(sortProviderSmsData, sortCountrySmsData)

	sort.SliceStable(sortCountrySmsData, func(i, j int) bool { return sortCountrySmsData[i].Country < sortCountrySmsData[j].Country })
	sort.SliceStable(sortProviderSmsData, func(i, j int) bool { return sortProviderSmsData[i].Provider < sortProviderSmsData[j].Provider })
	resultSmsData := make([][]domain.SMSData, 0)
	resultSmsData = append(resultSmsData, sortProviderSmsData, sortCountrySmsData)
	return resultSmsData
}

func processMms(path string, countries *helper.Countries) [][]domain.MMSData {
	dataMms := mms.ParseData(path)
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
	return resultMmsData
}

func processVoiceCall(path string) []domain.VoiceCallData {
	dataVoiceCall := voicecall.ParseData(path)
	return dataVoiceCall
}

func processEmail(path string) map[string][][]domain.EmailData {
	dataEmail := email.ParseData(path)
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
	return resultMap
}

func processBilling(path string) *domain.BillingData {
	dataBilling := biling.ParseData(path)
	return dataBilling
}

func processSupport(path string) []int {
	dataSupport := support.ParseData(path)
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
	resultDataSupport := []int{idxLoad, int(averageTime)}
	return resultDataSupport
}

func processIncident(path string) []domain.IncidentData {
	dataIncident := incident.ParseData(path)
	sort.SliceStable(dataIncident, func(i, j int) bool { return dataIncident[i].Status < dataIncident[j].Status })
	return dataIncident
}
