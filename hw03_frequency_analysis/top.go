// Package hw03frequencyanalysis provides the Top10 function that returns
// the ten most frequent words in a given text.
package hw03frequencyanalysis

import (
	"cmp"
	"maps"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

var re = regexp.MustCompile(`^[[:punct:]]+|[[:punct:]]+$`)

type pair struct {
	word  string
	count int
}

// Top10 returns up to ten most frequent words from the given string.
// Words are compared case-insensitively; edge punctuation is trimmed.
// Ties in frequency are broken lexicographically.
func Top10(str string) []string {
	// convert string to lowercase
	str = strings.ToLower(str)

	// fill map to count repetitions
	fields := strings.Fields(str)
	fieldsMap := make(map[string]*pair)
	for _, word := range fields {
		trimWord := re.ReplaceAllString(word, "")

		if trimWord == "" {
			if utf8.RuneCountInString(word) < 2 {
				continue // "-", ",", "!" not he-he
			}
			trimWord = word // "---", "!!!" he-he
		}

		if p, ok := fieldsMap[trimWord]; ok {
			p.count++
			continue
		}

		fieldsMap[trimWord] = &pair{
			// using a regular expression to trim punctuation marks
			word:  trimWord,
			count: 1,
		}
	}

	// convert map to iter (slice)
	fieldsIter := maps.Values(fieldsMap)

	// sorting:
	// - 1 - if words do not occur the same number of times, compare them in descending order;
	// - 2 - it's the same, then sort it lexically.
	sortedFields := slices.SortedFunc(
		fieldsIter,
		func(p1, p2 *pair) int {
			if p1.count != p2.count {
				return -cmp.Compare(p1.count, p2.count)
			}
			return cmp.Compare(p1.word, p2.word)
		},
	)

	// create summary slice
	res := make([]string, 0, 10)
	if len(sortedFields) > 10 {
		for _, p := range sortedFields[:10] {
			res = append(res, p.word)
		}
	} else {
		for _, p := range sortedFields {
			res = append(res, p.word)
		}
	}

	return res
}
