package day4

import (
	"fmt"
	"testing"
)

func TestMine(t *testing.T) {
	cases := []struct {
		in    string
		zeros int
		want  int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
		{"ckczppom", 5, 117946},
		{"ckczppom", 6, 3938038},
		{"yzbqklnj", 5, 282749},
		{"yzbqklnj", 6, 9962624},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%s/%d", tc.in, tc.zeros)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Mine([]byte(tc.in), tc.zeros)
			if err != nil {
				t.Fatal(err)
			}
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
