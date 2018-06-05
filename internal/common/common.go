package common

import (
	"log"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func init() {

	r, err := regexp.Compile(`([^a-zA-Z0-9])+`)

	if err != nil {
		log.Fatalf("failed to compile regular expression: %v", err)
	}

	reg = r
}

func Tokenizer(term string) []string {

	var tokenized []string

	splitterms := strings.Fields(term)

	for _, text := range splitterms {

		// removing any non-alphanumeric characters in the word
		processedString := reg.ReplaceAllString(text, "")

		if processedString == "" {
			continue
		}

		lowecased := strings.ToLower(processedString)
		tokenized = append(tokenized, lowecased)
	}

	return tokenized
}
