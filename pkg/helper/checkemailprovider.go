package helper

var checkEmailProvider = map[string]struct{}{
	"Gmail":      {},
	"Yahoo":      {},
	"Hotmail":    {},
	"MSN":        {},
	"Orange":     {},
	"Comcast":    {},
	"AOL":        {},
	"Live":       {},
	"RediffMail": {},
	"GMX":        {},
	"Protonmail": {},
	"Yandex":     {},
	"Mail.ru":    {},
}

func CheckEmailProvider(key string) bool {
	_, exist := checkEmailProvider[key]

	return exist
}
