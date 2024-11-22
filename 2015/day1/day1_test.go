package day1

import (
	"os"
	"testing"
)

func readFile(t *testing.T, name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestPart1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tc := range cases {
		got := Part1(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
		}
	}
}

func TestPart1_FromFile(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input.txt", 232},
		{"testdata/input2.txt", 74},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got := Part1(readFile(t, tc.in))
			if tc.want != got {
				t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{")", 1},
		{"()())", 5},
		{"(", -1},
	}

	for _, tc := range cases {
		got := Part2(tc.in)
		if tc.want != got {
			t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
		}
	}
}

func TestPart2_FromFile(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input.txt", 1783},
		{"testdata/input2.txt", 1795},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got := Part2(readFile(t, tc.in))
			if tc.want != got {
				t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
			}
		})
	}
}
