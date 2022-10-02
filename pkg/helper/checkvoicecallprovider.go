package helper

var voiceCallProvider = map[string]struct{}{
	"TransparentCalls": {},
	"E-Voice":          {},
	"JustPhone":        {},
}

func CheckVoiceCallProvider(key string) bool {
	_, exist := voiceCallProvider[key]

	return exist
}
