package main

import (
	"testing"
)

func TestIsSymbol(t *testing.T) {
	tests := []struct {
		name string
		symbol   rune
		excepted bool
	}{
		{"a symbol", 'a', true},
		{"5 symbol", '5', false},
		{"9 symbol", '9', false},
		{"0 symbol", '0', false},
		{"'9' + 1 symbol", rune(byte('9') + 1), true},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			actual := isSymbol(e.symbol)
			if actual != e.excepted{
				t.Errorf("actual: %v, excepted: %v", actual, e.excepted)
			}
		})
		
	}
}

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name string
		input    string
		excepted string
		err      error
	}{
		//base tests
		{name: "base 1", input: "a4bc2d5e", excepted: "aaaabccddddde", err: nil},
		{name: "base 2", input: "abcd", excepted: "abcd", err: nil},
		{name: "base 3", input: "45", excepted: "", err: ErrBadInput},
		{name: "base 4", input: "", excepted: "", err: nil},
		{name: "base 5", input: `qwe\4\5`, excepted: "qwe45", err: nil},
		{name: "base 6", input: `qwe\45`, excepted: "qwe44444", err: nil},
		{name: "base 7", input: `qwe\\5`, excepted: `qwe\\\\\`, err: nil},
		//my
		{name: "my 1", input: `\111`, excepted: `11111111111`, err: nil},
		{name: "my 2", input: `\111\`, excepted: ``, err: ErrBadInput},
		{name: "my 3", input: `\110\220`, excepted: `111111111122222222222222222222`, err: nil},
		{name: "my 5", input: `a`, excepted: `a`, err: nil},
		{name: "my 6", input: `\a`, excepted: `a`, err: nil},
		{name: "my 7", input: `\1`, excepted: `1`, err: nil},
		{name: "my 8", input: ` 10`, excepted: `          `, err: nil},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			actual, actualErr := UnpackString(e.input)
			if actual != e.excepted{
				t.Errorf("actual: %v, excepted: %v", actual, e.excepted)
			}
			if actualErr != e.err{
				t.Errorf("actual: %v, excepted: %v", actualErr, e.err)
			}
		})
	}
}
