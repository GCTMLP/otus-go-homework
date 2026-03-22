package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "z5", expected: "zzzzz"},
		{input: "a1b2c3", expected: "abbccc"},
		{input: "x0y2z3", expected: "yyzzz"},
		{input: "!2@3#4", expected: "!!@@@####"},
		{input: "😀3😂2", expected: "😀😀😀😂😂"},
		{input: "😀0😂3", expected: "😂😂😂"},
		{input: `\\\\5`, expected: `\\\\\\`},
		{input: `a\\b`, expected: `a\b`},
		{input: `\1\2\3`, expected: `123`},
		{input: `\\1`, expected: `\`},
		{input: `qwe\\5\\3`, expected: `qwe\\\\\\\\`},
		{input: `abc\`, expected: `abc\`},
		{input: `🙃\3`, expected: `🙃3`},
		{input: `a\2b3c\4`, expected: `a2bbbc4`},
		{input: `\70`, expected: ``},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a0b0c0", expected: ""},
		{input: `\1\2\3\4`, expected: "1234"},
		{input: `\\\\`, expected: `\\`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
