package day5

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// TODO: Find another problem input.

var (
	vowelRegexp  = regexp.MustCompile(`[aeiou]`)
	bannedRegexp = regexp.MustCompile(`ab|cd|pq|xy`)

	doubledLetterRegexp *regexp.Regexp // Initialized at runtime.
)

func init() {
	// The regexp package doesn't support backreferences.
	// Instead, use `aa | bb | ... | yy | zz`.
	var s []string
	for b := 'a'; b <= 'z'; b++ {
		s = append(s, fmt.Sprintf("%c%c", b, b))
	}
	doubledLetterRegexp = regexp.MustCompile(strings.Join(s, "|"))
}

func ContainsAtLeast3Vowels(s string) bool {
	return len(vowelRegexp.FindAllString(s, 3)) == 3
}

func ContainsDoubledLetter(s string) bool {
	return doubledLetterRegexp.MatchString(s)
}

func ContainsBannedSubstring(s string) bool {
	return bannedRegexp.MatchString(s)
}

// IsNice follows the rules from Part 1.
func IsNice(s string) bool {
	return ContainsAtLeast3Vowels(s) && ContainsDoubledLetter(s) && !ContainsBannedSubstring(s)
}

// ContainsSeparateRepeatPair returns true if s contains a pair of letters
// that repeats without overlap.
//
// For example, "cookbook" contains "oo" twice, without overlap.
func ContainsSeparateRepeatPair(s string) bool {
	// The regexp package doesn't support lookahead assertions.

	// Two separate pairs take up 4 characters.
	if len(s) < 4 {
		return false
	}

	firstSeen := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		// Encounter a new pair.
		pair := s[i : i+2]

		// Check if we've seen it before.
		j, ok := firstSeen[pair]

		// If so, check whether we're over one letter away.
		if ok {
			if i > j+1 {
				return true
			}
		} else {
			firstSeen[pair] = i
		}
	}

	return false
}

// ContainsRepeatLetterOneApart returns true if s contains a letter that repeats
// with exactly one letter between the repeats.
//
// For example, "editing" contains "iti".
func ContainsRepeatLetterOneApart(s string) bool {
	// The regexp package doesn't support lookahead assertions.

	// The required pattern uses 3 letters.
	if len(s) < 3 {
		return false
	}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}

// IsNice2 follows the rules from Part 2.
func IsNice2(s string) bool {
	return ContainsSeparateRepeatPair(s) && ContainsRepeatLetterOneApart(s)
}

func Part1(r io.Reader) (int, error) {
	var res int

	s := bufio.NewScanner(r)
	for s.Scan() {
		if IsNice(s.Text()) {
			res++
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return res, nil
}

func Part2(r io.Reader) (int, error) {
	var res int

	s := bufio.NewScanner(r)
	for s.Scan() {
		if IsNice2(s.Text()) {
			res++
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return res, nil
}
