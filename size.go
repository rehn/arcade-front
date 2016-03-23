package main

import (
	"strconv"
	"strings"
)

type Size struct {
	W, H int32
}

func (size *Size) UnmarshalText(data []byte) error {
	var err error
	sp := strings.Split(string(data), ",")
	cl := [2]int32{0, 0}
	for i, s := range sp {
		n, err := strconv.Atoi(s)
		if err != nil {
			LogWarning(err)
		} else {
			cl[i] = int32(n)
		}
	}
	size.W = cl[0]
	size.H = cl[1]
	return err
}
