package helper

import "diploma/internal/domain"

func GetLen(mas []domain.EmailData) int {
	if len(mas) < 3 {
		return len(mas)
	} else {
		return 3
	}
}
