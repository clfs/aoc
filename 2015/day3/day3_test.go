package day3

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
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
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
		{"testdata/input1.txt", 2572},
		{"testdata/input2.txt", 2592},
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
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
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
		{"testdata/input1.txt", 2631},
		{"testdata/input2.txt", 2360},
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
