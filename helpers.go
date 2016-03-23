package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func PathToABS(inPath string) (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return inPath, err
	}
	if !strings.HasPrefix(inPath, "/") {
		inPath = dir + "/" + inPath
	}
	return inPath, nil
}

func RectangleFromList(points [4]int32) sdl.Rect {
	return sdl.Rect{points[0], points[1], points[2], points[3]}
}
