package day3

import (
	"bufio"
	"fmt"
	"io"
)

type Direction byte

const (
	North Direction = '^'
	East  Direction = '>'
	South Direction = 'v'
	West  Direction = '<'
)

type Point struct {
	X, Y int
}

func (p *Point) Move(d Direction) error {
	switch d {
	case North:
		p.Y++
	case East:
		p.X++
	case South:
		p.Y--
	case West:
		p.X--
	default:
		return fmt.Errorf("bad direction: %v", d)
	}

	return nil
}

func Part1(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	var santa Point

	seen := map[Point]struct{}{
		santa: {},
	}

	for {
		// Move to the next location.
		b, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		if err := santa.Move(Direction(b)); err != nil {
			return 0, err
		}

		// Mark the new location as seen.
		seen[santa] = struct{}{}
	}

	return len(seen), nil
}

func Part2(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	var (
		santa     Point
		roboSanta Point
	)

	seen := map[Point]struct{}{
		santa:     {},
		roboSanta: {},
	}

	// Even turns are Santa; odd turns are Robo-Santa.
	for turn := 0; ; turn++ {
		// Fetch the selected santa's location.
		var p Point
		if turn%2 == 0 {
			p = santa
		} else {
			p = roboSanta
		}

		// Move to the next location.
		b, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		if err := p.Move(Direction(b)); err != nil {
			return 0, err
		}

		// Mark the new location as seen.
		seen[p] = struct{}{}

		// Pass back the new location.
		if turn%2 == 0 {
			santa = p
		} else {
			roboSanta = p
		}
	}

	return len(seen), nil
}
