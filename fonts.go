package main

import "github.com/veandco/go-sdl2/sdl_ttf"

type Fonts struct {
	Font14 *ttf.Font
	Font16 *ttf.Font
	Font20 *ttf.Font
}

func (f *Fonts) init() {
	var err error
	fnt, err := PathToABS("assets/fonts/Krungthep.ttf")
	LogFatal(err)
	f.Font14, err = ttf.OpenFont(fnt, 14)
	LogFatal(err)
	f.Font16, err = ttf.OpenFont(fnt, 16)
	LogFatal(err)
	f.Font20, err = ttf.OpenFont(fnt, 20)
	LogFatal(err)
}
