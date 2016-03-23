package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	"strings"
)

type Color struct {
	Color      sdl.Color
	R, G, B, A uint8
}

func (c *Color) UnmarshalText(data []byte) error {
	var err error
	sp := strings.Split(string(data), ",")
	cl := []uint8{0, 0, 0, 255}
	for i, s := range sp {
		n, err := strconv.Atoi(s)
		if err != nil {
			LogWarning(err)
		} else {
			cl[i] = uint8(n)
		}
	}
	c.R = cl[0]
	c.G = cl[1]
	c.B = cl[2]
	c.A = cl[3]
	c.Color = sdl.Color{c.R, c.G, c.B, c.A}
	return err
}
