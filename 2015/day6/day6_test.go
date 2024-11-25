package day6

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"turn on 0,0 through 999,999", 1000000},
		{"toggle 0,0 through 999,0", 1000},
		{"turn off 499,499 through 500,500", 0},
	}

	for _, tc := range cases {
		got, err := Part1(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%q: error: %v", tc.in, err)
		}
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
		{"testdata/input1.txt", 569999},
		{"testdata/input2.txt", 400410},
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
		{"turn on 0,0 through 0,0", 1},
		{"toggle 0,0 through 999,999", 2000000},
	}

	for _, tc := range cases {
		got, err := Part2(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%q: error: %v", tc.in, err)
		}
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
		{"testdata/input1.txt", 17836115},
		{"testdata/input2.txt", 15343601},
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
