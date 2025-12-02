package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func CheckStr(s string) (string, error) {
	flag := true

	if len(s) < 1 {
		return "", errors.New("string must be not empty")
	}

	for _, val := range s {
		if val == '.' || val == '-' || val == ' ' {
			continue
		}
		flag = false
		break
	}

	if !flag {
		return morse.ToMorse(s), nil
	}

	return morse.ToText(s), nil
}
