package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var markerPattern = regexp.MustCompile(`\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)
var spacesBeforePunctuation = regexp.MustCompile(`\s+([.,!?:;]+)`)
var spacesAfterPunctuation = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func processText(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = Process(line)
	}
	return strings.Join(lines, "\n")
}

// Process applies the Go Reloaded rules to one line of text.
func Process(line string) string {
	line = normalizeMarkers(line)
	words := applyMarkers(strings.Fields(line))
	line = strings.Join(words, " ")
	line = spacesBeforePunctuation.ReplaceAllString(line, "$1")
	line = spacesAfterPunctuation.ReplaceAllString(line, "$1 $2")
	return fixArticles(fixQuotes(line))
}

func normalizeMarkers(line string) string {
	return markerPattern.ReplaceAllStringFunc(line, func(marker string) string {
		return " " + strings.ReplaceAll(marker, ", ", ",") + " "
	})
}

func applyMarkers(words []string) []string {
	out := make([]string, 0, len(words))
	for _, word := range words {
		match := markerPattern.FindStringSubmatch(word)
		if match == nil {
			out = append(out, word)
			continue
		}

		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}
		for i := len(out) - 1; i >= 0 && count > 0; i-- {
			if !isWord(out[i]) {
				continue
			}
			out[i] = transformWord(out[i], match[1])
			count--
		}
	}
	return out
}

func isWord(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func transformWord(word, command string) string {
	switch command {
	case "hex":
		return toDecimal(word, 16)
	case "bin":
		return toDecimal(word, 2)
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		return capitalize(word)
	default:
		return word
	}
}

func toDecimal(word string, base int) string {
	number, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(number, 10)
}

func capitalize(word string) string {
	letters := []rune(strings.ToLower(word))
	if len(letters) == 0 {
		return word
	}
	letters[0] = unicode.ToUpper(letters[0])
	return string(letters)
}

func fixQuotes(line string) string {
	words := strings.Fields(line)
	result := make([]string, 0, len(words))
	quoteIsOpen := false

	for _, word := range words {
		if word == "'" {
			if quoteIsOpen {
				if len(result) > 0 {
					result[len(result)-1] += "'"
				}
			} else {
				result = append(result, "'")
			}
			quoteIsOpen = !quoteIsOpen
			continue
		}

		if quoteIsOpen && len(result) > 0 && result[len(result)-1] == "'" {
			result[len(result)-1] += word
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func fixArticles(line string) string {
	words := strings.Fields(line)
	for i := 0; i+1 < len(words); i++ {
		if strings.EqualFold(words[i], "a") && startsWithVowelSound(words[i+1]) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

func startsWithVowelSound(word string) bool {
	word = strings.TrimLeft(word, "'")
	for _, first := range strings.ToLower(word) {
		return strings.ContainsRune("aeiouh", first)
	}
	return false
}
