package main

import (
	"errors"
	"strings"
)

var (
	ErrBadInput = errors.New("incorrect string")
)

func UnpackString(str string) (string, error) {
	runes := []rune(str)
	var symbol rune
	var count int = 1
	builder := strings.Builder{}
	for i := 0; i < len(runes); {
		if runes[i] == '\\' {
			i++
			if i >= len(runes) {
				return "", ErrBadInput
			}
			symbol = runes[i]
		} else {
			if !isSymbol(runes[i]) {
				return "", ErrBadInput
			}
			symbol = runes[i]
		}
		i++
		if i >= len(runes) || isSymbol(runes[i]) {
			count = 1
		} else {
			count = 0
			for i < len(runes) && !isSymbol(runes[i]) {
				count = count*10 + int(runes[i]-'0')
				i++
			}
		}
		for i := 0; i < count; i++ {
			builder.Write([]byte{byte(symbol)})
		}
	}
	return builder.String(), nil
}

func isSymbol(r rune) bool {
	return r < '0' || r > '9'
}
