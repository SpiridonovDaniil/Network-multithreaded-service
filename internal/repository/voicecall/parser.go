package voicecall

import (
	"diploma/internal/domain"
	"diploma/pkg/helper"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ParseData(path string) []domain.VoiceCallData {
	path = filepath.Clean(path)

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	r := csv.NewReader(file)

	r.Comma = ';'
	r.FieldsPerRecord = -1
	countries := helper.NewCountries("data/countries.json")
	answer := make([]domain.VoiceCallData, 0)
	for voiceCallData, errRead := r.Read(); !errors.Is(errRead, io.EOF); voiceCallData, errRead = r.Read() {
		if errRead != nil {
			log.Println(err)
			continue
		}

		if len(voiceCallData) < 8 || !countries.CheckCountry(voiceCallData[0]) || !helper.CheckVoiceCallProvider(voiceCallData[3]) {
			continue
		}

		num, err := strconv.Atoi(voiceCallData[1])
		if err != nil {
			log.Println(err)
			continue
		}

		if num > 100 || num < 0 {
			continue
		}

		connectionStabillity, err := strconv.ParseFloat(voiceCallData[4], 32)
		if err != nil {
			log.Println(err)
			continue
		}

		ttfb, err := strconv.Atoi(voiceCallData[5])
		if err != nil {
			log.Println(err)
			continue
		}

		voicePurity, err := strconv.Atoi(voiceCallData[6])
		if err != nil {
			log.Println(err)
			continue
		}

		medianOfCallsTime, err := strconv.Atoi(voiceCallData[7])
		if err != nil {
			log.Println(err)
			continue
		}

		answer = append(answer,
			domain.VoiceCallData{
				Country:             voiceCallData[0],
				Bandwidth:           voiceCallData[1],
				ResponseTime:        voiceCallData[2],
				Provider:            voiceCallData[3],
				ConnectionStability: float32(connectionStabillity),
				TTFB:                ttfb,
				VoicePurity:         voicePurity,
				MedianOfCallsTime:   medianOfCallsTime,
			})
	}

	return answer
}
