package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	"strings"
)

type Rect struct {
	X, Y, W, H int32
	Rectangle  sdl.Rect
}

func (rect *Rect) UnmarshalText(data []byte) error {
	var err error
	sp := strings.Split(string(data), ",")
	cl := [4]int32{0, 0, 0, 0}
	for i, s := range sp {
		n, err := strconv.Atoi(s)
		if err != nil {
			LogWarning(err)
		} else {
			cl[i] = int32(n)
		}
	}
	rect.X = cl[0]
	rect.Y = cl[1]
	rect.W = cl[2]
	rect.H = cl[3]
	rect.Rectangle = sdl.Rect{rect.X, rect.Y, rect.W, rect.H}
	return err
}
