package helper

var provider = map[string]struct{}{
	"Topolo": {},
	"Rond":   {},
	"Kildy":  {},
}

func CheckProvider(key string) bool {
	_, exist := provider[key]

	return exist
}