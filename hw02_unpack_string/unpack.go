package main

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unpackingString string) (string, error) {
	if unpackingString == "" {
		return "", nil
	}
	runes := []rune(unpackingString)
	result := make([]rune, 0, len(runes))
	for i := 0; i < len(runes); i++ {
		current := runes[i]
		if i == 0 && unicode.IsDigit(current) {
			// если первый символ число - ошибка
			return "", ErrInvalidString
		}
		if current == 92 && i+1 < len(runes) {
			// оббработка "/"
			if unicode.IsLetter(runes[i+1]) {
				// если после "/" буква - ошибка
				return "", ErrInvalidString
			} else if i+2 < len(runes) && runes[i+1] == 92 && unicode.IsDigit(runes[i+2]) {
				// если после "/" стоит "/" и после него цифра
				if unicode.IsDigit(runes[i+2]) {
					num := int(runes[i+2] - '0')
					for j := 0; j < num; j++ {
						result = append(result, current)
					}
					i += 2
				}
			} else if runes[i+1] == 92 {
				// если после "/" стоит "/" и дальше не цифра
				result = append(result, current)
				i++
			}
			continue
		}
		if i > 0 && unicode.IsDigit(current) && unicode.IsDigit(runes[i-1]) && runes[i-2] != 92 {
			// если две цифры рялом и нет экранирования
			return "", ErrInvalidString
		}
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			// если после символа цифра
			num := int(runes[i+1] - '0')
			if num == 0 {
			} else {
				for j := 0; j < num; j++ {
					result = append(result, current)
				}
			}
			i++
		} else {
			// если просто символ
			result = append(result, current)
		}
	}
	return string(result), nil
}
