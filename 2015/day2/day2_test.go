package day2

import (
	"os"
	"testing"
)

func mustNewPresent(t *testing.T, s string) *Present {
	p, err := NewPresent(s)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestPaperRequired(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, tc := range cases {
		p := mustNewPresent(t, tc.in)
		got := p.PaperRequired()
		if tc.want != got {
			t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input1.txt", 1598415},
		{"testdata/input2.txt", 1588178},
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
				t.Errorf("err: %v", err)
			}
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestRibbonRequired(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, tc := range cases {
		p := mustNewPresent(t, tc.in)
		got := p.RibbonRequired()
		if tc.want != got {
			t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"testdata/input1.txt", 3812909},
		{"testdata/input2.txt", 3783758},
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
				t.Errorf("err: %v", err)
			}
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
