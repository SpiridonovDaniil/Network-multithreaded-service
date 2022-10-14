package helper

import "diploma/internal/domain"

func CheckNilStruct(getResultData domain.ResultSetT) bool {
	if getResultData.Billing == nil {
		return false
	}
	if getResultData.Support == nil {
		return false
	}
	if getResultData.SMS == nil {
		return false
	}
	if getResultData.Email == nil {
		return false
	}
	if getResultData.MMS == nil {
		return false
	}
	if getResultData.Incidents == nil {
		return false
	}
	if getResultData.VoiceCall == nil {
		return false
	}

	return true
}
