package day4

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

const maxTrials = 10000000

func Mine(b []byte, zeros int) (int, error) {
	var (
		h         = md5.New()
		hexdigest = make([]byte, md5.Size*2)
		prefix    = bytes.Repeat([]byte("0"), zeros)
	)

	for n := 1; n <= maxTrials; n++ {
		h.Reset()

		// MD5(b || n)
		h.Write(fmt.Appendf(b, "%d", n))

		hex.Encode(hexdigest, h.Sum(nil))

		if bytes.HasPrefix(hexdigest, prefix) {
			return n, nil
		}
	}

	return 0, errors.New("not found")
}

func Part1(b []byte) (int, error) {
	return Mine(b, 5)
}

func Part2(b []byte) (int, error) {
	return Mine(b, 6)
}
