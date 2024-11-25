package day6

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"io"
	"regexp"
	"strconv"
)

type Action string

const (
	TurnOn  Action = "turn on"
	TurnOff Action = "turn off"
	Toggle  Action = "toggle"
)

type Instruction struct {
	action Action
	rect   image.Rectangle
}

var instructionRegexp = regexp.MustCompile(`^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func (ins *Instruction) UnmarshalText(text []byte) error {
	matches := instructionRegexp.FindSubmatch(text)
	if matches == nil {
		return fmt.Errorf("bad instruction: %q", text)
	}

	x0, err := strconv.Atoi(string(matches[2]))
	if err != nil {
		return fmt.Errorf("bad x0: %v", err)
	}

	y0, err := strconv.Atoi(string(matches[3]))
	if err != nil {
		return fmt.Errorf("bad y0: %v", err)
	}

	x1, err := strconv.Atoi(string(matches[4]))
	if err != nil {
		return fmt.Errorf("bad x1: %v", err)
	}

	y1, err := strconv.Atoi(string(matches[5]))
	if err != nil {
		return fmt.Errorf("bad y1: %v", err)
	}

	ins.action = Action(matches[1])
	ins.rect = image.Rect(x0, y0, x1+1, y1+1)

	return nil
}

func NewInstruction(s string) (Instruction, error) {
	var ins Instruction
	if err := ins.UnmarshalText([]byte(s)); err != nil {
		return Instruction{}, err
	}
	return ins, nil
}

// Pattern represents a pattern of lights.
type Pattern struct {
	// For this problem, it's safe to represent brightness as a uint16.
	// The input is 300 instructions, so the brightest light is at most 600.
	img *image.Gray16
}

func NewPattern(r image.Rectangle) *Pattern {
	return &Pattern{img: image.NewGray16(r)}
}

// TurnOn sets the brightness of the selected lights to 1.
func (p *Pattern) TurnOn(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		return color.Gray16{1}
	}
	p.do(r, f)
}

// TurnOff sets the brightness of the selected lights to 0.
func (p *Pattern) TurnOff(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		return color.Gray16{0}
	}
	p.do(r, f)
}

// Toggle toggles whether the selected lights are lit.
// Unlit lights are set to a brightness of 1.
func (p *Pattern) Toggle(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		if v.Y == 0 {
			return color.Gray16{1}
		}
		return color.Gray16{0}
	}
	p.do(r, f)
}

// TurnOn2 increases the brightness of the selected lights by 1.
func (p *Pattern) TurnOn2(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		return color.Gray16{Y: v.Y + 1}
	}
	p.do(r, f)
}

// TurnOff2 decreases the brightness of the selected lights by 1.
// Unlit lights are unchanged.
func (p *Pattern) TurnOff2(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		if v.Y > 0 {
			return color.Gray16{Y: v.Y - 1}
		}
		return v
	}
	p.do(r, f)
}

// Toggle2 increases the brightness of the selected lights by 2.
func (p *Pattern) Toggle2(r image.Rectangle) {
	f := func(v color.Gray16) color.Gray16 {
		return color.Gray16{Y: v.Y + 2}
	}
	p.do(r, f)
}

// do applies f to the lights within r.
func (p *Pattern) do(r image.Rectangle, f func(v color.Gray16) color.Gray16) {
	b := p.img.SubImage(r).Bounds()

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			v := p.img.Gray16At(x, y)
			p.img.SetGray16(x, y, f(v))
		}
	}
}

// NLit returns the number of lights that are lit.
func (p *Pattern) NLit() int {
	var res int

	b := p.img.Bounds()

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			g := p.img.Gray16At(x, y)
			if g.Y > 0 {
				res++
			}
		}
	}

	return res
}

func (p *Pattern) TotalBrightness() int {
	var res int

	b := p.img.Bounds()

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			g := p.img.Gray16At(x, y)
			res += int(g.Y)
		}
	}

	return res
}

func Part1(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)

	p := NewPattern(image.Rect(0, 0, 1000, 1000))

	for s.Scan() {
		ins, err := NewInstruction(s.Text())
		if err != nil {
			return 0, err
		}

		switch ins.action {
		case TurnOn:
			p.TurnOn(ins.rect)
		case TurnOff:
			p.TurnOff(ins.rect)
		case Toggle:
			p.Toggle(ins.rect)
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return p.NLit(), nil
}

func Part2(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)

	p := NewPattern(image.Rect(0, 0, 1000, 1000))

	for s.Scan() {
		ins, err := NewInstruction(s.Text())
		if err != nil {
			return 0, err
		}

		switch ins.action {
		case TurnOn:
			p.TurnOn2(ins.rect)
		case TurnOff:
			p.TurnOff2(ins.rect)
		case Toggle:
			p.Toggle2(ins.rect)
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return p.TotalBrightness(), nil
}
