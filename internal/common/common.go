package common

import (
	"strings"
	"unicode"
)

func Tokenizer(term string) []string {
	var tokenized []string

	splitterms := strings.Fields(term)

	for _, text := range splitterms {

		// removing any non-alphanumeric characters in the word
		processedString := strings.TrimFunc(text, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		if processedString == "" {
			continue
		}

		lowecased := strings.ToLower(processedString)
		tokenized = append(tokenized, lowecased)
	}

	return tokenized
}
