package day2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Present struct {
	Length, Width, Height int
}

// NewPresent returns a new [Present].
func NewPresent(s string) (*Present, error) {
	var p Present
	if err := p.UnmarshalText([]byte(s)); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Present) UnmarshalText(text []byte) error {
	fields := bytes.SplitN(text, []byte("x"), 3)

	if n := len(fields); n != 3 {
		return fmt.Errorf("wrong number of dimensions: %d", n)
	}

	length, err := strconv.Atoi(string(fields[0]))
	if err != nil {
		return fmt.Errorf("bad length: %q", fields)
	}

	width, err := strconv.Atoi(string(fields[1]))
	if err != nil {
		return fmt.Errorf("bad width: %q", fields)
	}

	height, err := strconv.Atoi(string(fields[2]))
	if err != nil {
		return fmt.Errorf("bad height: %q", fields)
	}

	p.Length = length
	p.Width = width
	p.Height = height

	return nil
}

// SurfaceArea returns the surface area of the present.
func (p *Present) SurfaceArea() int {
	// 2lw + 2wh + 2hl
	return 2*p.Length*p.Width + 2*p.Width*p.Height + 2*p.Height*p.Length
}

// AreaSmallestSide returns the area of the smallest side of the present.
func (p *Present) AreaSmallestSide() int {
	return min(p.Length*p.Width, p.Width*p.Height, p.Height*p.Length)
}

// PaperRequired returns the area of paper needed to wrap the present.
func (p *Present) PaperRequired() int {
	// The area of the smallest side is slack.
	return p.SurfaceArea() + p.AreaSmallestSide()
}

// MinFacePerimeter returns the smallest perimeter of any one face.
func (p *Present) MinFacePerimeter() int {
	return min(
		2*p.Length+2*p.Width,
		2*p.Width+2*p.Height,
		2*p.Height+2*p.Length,
	)
}

// Volume returns the volume of the present.
func (p *Present) Volume() int {
	return p.Length * p.Width * p.Height
}

// RibbonRequired returns the area of ribbon needed to decorate the present.
func (p *Present) RibbonRequired() int {
	return p.MinFacePerimeter() + p.Volume()
}

func Part1(r io.Reader) (int, error) {
	var res int

	s := bufio.NewScanner(r)
	for s.Scan() {
		p, err := NewPresent(s.Text())
		if err != nil {
			return 0, err
		}

		res += p.PaperRequired()
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
		p, err := NewPresent(s.Text())
		if err != nil {
			return 0, err
		}

		res += p.RibbonRequired()
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return res, nil
}
