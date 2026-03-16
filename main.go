package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(data)
	processed := processText(text)

	err = os.WriteFile(outputPath, []byte(processed), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
}

func processText(text string) string {
	lines := strings.Split(text, "\n")
	var processedLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			processedLines = append(processedLines, line)
			continue
		}
		processedLines = append(processedLines, processLine(line))
	}

	return strings.Join(processedLines, "\n")
}

func processLine(line string) string {
	// 1. Tokenize by spaces directly to preserve modifier structure
	words := strings.Fields(line)

	// 2. Apply modifiers
	words = applyModifiers(words)

	// 3. Handle Articles (a -> an)
	words = applyArticles(words)

	// 4. Join tokens back
	result := strings.Join(words, " ")

	// 5. Formatting rules (Punctuation Space Cleanup)
	result = formatPunctuation(result)

	// 6. Format Quotes
	result = formatQuotes(result)

	return result
}

func applyModifiers(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Regex to find modifier inside token, supporting glued punctuation, e.g., (up),
		reSingle := regexp.MustCompile(`^\((hex|bin|up|low|cap)\)(.*)$`)
		reCount := regexp.MustCompile(`^\((up|low|cap),\s*(\d+)\)(.*)$`)

		if reSingle.MatchString(word) {
			matches := reSingle.FindStringSubmatch(word)
			mod := matches[1]
			suffix := matches[2]

			if i > 0 {
				applyMod(words, i-1, mod)
			}
			words[i] = suffix
			if suffix == "" {
				// Remove token if empty
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if reCount.MatchString(word) {
			matches := reCount.FindStringSubmatch(word)
			mod := matches[1]
			count, _ := strconv.Atoi(matches[2])
			suffix := matches[3]

			for j := 1; j <= count; j++ {
				if i-j >= 0 {
					applyMod(words, i-j, mod)
				}
			}
			words[i] = suffix
			if suffix == "" {
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,") {
			// Handle spaced cases like (up, 6)
			count, tokensToRemove := parseCount(words, i)
			mod := strings.TrimPrefix(strings.TrimSuffix(word, ","), "(")

			for j := 1; j <= count; j++ {
				if i-j >= 0 {
					applyMod(words, i-j, mod)
				}
			}
			words = append(words[:i], words[i+tokensToRemove:]...)
			i--
		} else if word == "(hex)" || word == "(bin)" || word == "(up)" || word == "(low)" || word == "(cap)" {
			// fallback check to make absolutely sure
			mod := strings.TrimSuffix(strings.TrimPrefix(word, "("), ")")
			if i > 0 {
				applyMod(words, i-1, mod)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func parseCount(words []string, i int) (int, int) {
	word := words[i]
	if strings.Contains(word, ")") {
		parts := strings.Split(word, ",")
		if len(parts) > 1 {
			nStr := strings.TrimSuffix(parts[1], ")")
			n, _ := strconv.Atoi(nStr)
			return n, 1
		}
	} else if i+1 < len(words) {
		nStr := strings.TrimSuffix(words[i+1], ")")
		n, _ := strconv.Atoi(nStr)
		return n, 2
	}
	return 1, 1
}

func applyMod(words []string, idx int, mod string) {
	switch mod {
	case "hex":
		val, _ := strconv.ParseInt(words[idx], 16, 64)
		words[idx] = fmt.Sprintf("%d", val)
	case "bin":
		val, _ := strconv.ParseInt(words[idx], 2, 64)
		words[idx] = fmt.Sprintf("%d", val)
	case "up":
		words[idx] = strings.ToUpper(words[idx])
	case "low":
		words[idx] = strings.ToLower(words[idx])
	case "cap":
		words[idx] = capitalize(words[idx])
	}
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

func applyArticles(words []string) []string {
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "a" {
			if i+1 < len(words) {
				next := words[i+1]
				if len(next) > 0 {
					firstChar := unicode.ToLower(rune(next[0]))
					if isVowel(firstChar) || firstChar == 'h' {
						if words[i] == "A" {
							words[i] = "An"
						} else {
							words[i] = "an"
						}
					}
				}
			}
		}
	}
	return words
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func formatPunctuation(s string) string {
	// 1. Remove space before punctuation
	reRemoveSpace := regexp.MustCompile(`\s+([.,!?:;]+)`)
	s = reRemoveSpace.ReplaceAllString(s, "$1")

	// 2. Add space after punctuation if there isn't one
	reAddSpace := regexp.MustCompile(`([.,!?:;]+)([^ \t.,!?:;])`)
	s = reAddSpace.ReplaceAllString(s, "$1 $2")

	return s
}

func formatQuotes(s string) string {
	reQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return reQuotes.ReplaceAllString(s, `'$1'`)
}
