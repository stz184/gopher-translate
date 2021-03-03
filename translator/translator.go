package translator

import (
	"regexp"
	"strings"
	"unicode"
)

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// It tries to preserve the original case, i.e.
// apple -> gapple, APPLE -> GAPPLE, Apple -> Gapple
func preserveCase(englishWord string, gopherWord string) string {
	// the whole word is uppercase
	if isUpper(englishWord) {
		return strings.ToUpper(gopherWord)
	}

	// just the first character of the word is uppercase, i.e. Title case
	if isUpper(englishWord[0:1]) {
		return strings.Title(strings.ToLower(gopherWord))
	}

	// the entire word is lower case
	return gopherWord
}

func TranslateWord(englishWord string) string {
	// skip translating shortened versions of words or apostrophes
	if strings.Contains(englishWord, "'") {
		return englishWord
	}

	vowelRx := regexp.MustCompile(`^(?i)[aeiou]+`)
	if vowelRx.MatchString(englishWord) {
		return preserveCase(englishWord, "g"+englishWord)
	}

	xrRegex := regexp.MustCompile(`^(?i)xr`)
	if xrRegex.MatchString(englishWord) {
		return preserveCase(englishWord, "ge"+englishWord)
	}

	consonantsQuRx := regexp.MustCompile(`^(?i)([^aeiou]+qu)([a-z]+)`)
	if consonantsQuRx.MatchString(englishWord) {
		gopherWord := consonantsQuRx.ReplaceAllString(englishWord, "${2}${1}ogo")
		return preserveCase(englishWord, gopherWord)
	}

	// consonant sounds only
	consonantsOnlyRx := regexp.MustCompile(`^(?i)[^aeiou]+$`)
	if consonantsOnlyRx.MatchString(englishWord) {
		gopherWord := englishWord + "ogo"
		return preserveCase(englishWord, gopherWord)
	}

	consonantsRx := regexp.MustCompile(`^(?i)([^aeiou]+)([a-z]+)`)
	if consonantsRx.MatchString(englishWord) {
		gopherWord := consonantsRx.ReplaceAllString(englishWord, "${2}${1}ogo")
		return preserveCase(englishWord, gopherWord)
	}

	return englishWord
}

func TranslateSentence(sentence string) string {
	englishSentence := sentence
	sentenceEndRx := regexp.MustCompile(`[.!?]$`)
	hasEndingCharacter := false
	if sentenceEndRx.MatchString(englishSentence) {
		englishSentence = englishSentence[0 : len(englishSentence)-1]
		hasEndingCharacter = true
	}

	sentenceParts := regexp.MustCompile(`\s+`).Split(englishSentence, -1)
	for i, word := range sentenceParts {
		sentenceParts[i] = TranslateWord(word)
	}

	if hasEndingCharacter {
		return strings.Join(sentenceParts, " ") + sentence[len(sentence)-1:]
	} else {
		return strings.Join(sentenceParts, " ")
	}
}
