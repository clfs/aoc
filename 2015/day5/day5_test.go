package day5

import (
	"os"
	"testing"
)

func TestContainsAtLeast3Vowels(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"aei", true},
		{"xazegov", true},
		{"aeiouaeiouaeiou", true},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range cases {
		got := ContainsAtLeast3Vowels(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestContainsDoubledLetter(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
	}

	for _, tc := range cases {
		got := ContainsDoubledLetter(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestContainsBannedSubstring(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"ugknbfddgicrmopn", false},
		{"haegwjzuvuyypxyu", true},
	}

	for _, tc := range cases {
		got := ContainsBannedSubstring(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestIsNice(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range cases {
		got := IsNice(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestContainsSeparateRepeatPair(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range cases {
		got := ContainsSeparateRepeatPair(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestContainsRepeatLetterOneApart(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}

	for _, tc := range cases {
		got := ContainsRepeatLetterOneApart(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestIsNice2(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range cases {
		got := IsNice2(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %t, got %t", tc.in, tc.want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input1.txt", 236},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			f, err := os.Open(tc.in)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got, err := Part1(f)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input1.txt", 51},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			f, err := os.Open(tc.in)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got, err := Part2(f)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
