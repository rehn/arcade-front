package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"math"
	"os"
	"strings"
)

func NewGame(games *Games, index int32) Game {
	var game Game
	game.Games = games
	game.Index = index
	x, y := game.CalcPosition()
	game.Src = games.application.config.FanartFallback.Rectangle
	game.Dst = sdl.Rect{x, y, games.ItemSize.W, games.ItemSize.H}
	game.executed = new(bool)

	return game
}

type Game struct {
	Name     string
	Title    string
	Image    string
	Type     string
	Last     string
	Count    int
	Enabled  bool
	Games    *Games
	Index    int32
	Dst      sdl.Rect
	Src      sdl.Rect
	Fanart   *sdl.Texture
	Text     *sdl.Texture
	TextSrc  sdl.Rect
	Emulator *Emulator
	executed *bool
}

func (g *Game) Update() {

	dst := g.Dst
	dst.Y -= *g.Games.TargetY
	if !dst.HasIntersection(&g.Games.Rectangle) {
		return
	}
	if g.Fanart == nil {
		g.Games.application.renderer.Copy(g.Games.application.texture, &g.Src, &dst)
	} else {
		g.Games.application.renderer.Copy(g.Fanart, &g.Src, &dst)
	}

	if g.Index == *g.Games.CurrentIndex {
		g.Games.application.renderer.SetDrawColor(223, 132, 32, 128)
		g.Games.application.renderer.FillRect(&dst)
	}
	adj := g.Games.application.config.GameTitleAjustment
	bg := g.Games.application.config.GameTitleBackground
	g.Games.application.renderer.SetDrawColor(bg.R, bg.G, bg.B, bg.A)
	g.Games.application.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	dstText := sdl.Rect{dst.X + (adj.W), dst.Y + (adj.H), g.TextSrc.W, g.TextSrc.H}
	g.Games.application.renderer.Copy(g.Text, &g.TextSrc, &dstText)

}

func (g *Game) DrawFanart() {
	if g.Fanart == nil {
		str := g.findImage()
		if str == "" {
			return
		}
		surface, err := img.Load(str)
		if err != nil {
			return
		}
		g.Fanart, err = g.Games.application.renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return
		}
		g.Src = sdl.Rect{0, 0, surface.W, surface.H}
	}
}

func (g *Game) findImage() string {
	filename := strings.TrimSuffix(g.Name, ".zip") + ".png"
	filenameJpg := strings.TrimSuffix(g.Name, ".zip") + ".jpg"

	path, err := PathToABS("assets/fanarts/" + filename)
	if err == nil {
		_, err = os.Stat(path)
		if err == nil {
			return path
		}
	}

	path, err = PathToABS("assets/fanarts/" + filenameJpg)
	if err == nil {
		_, err = os.Stat(path)
		if err == nil {
			return path
		}
	}

	path = g.Emulator.ScreenshotPath + "/" + filename
	_, err = os.Stat(path)
	if err == nil {
		return path
	}

	path = g.Emulator.ScreenshotPath + "/" + filenameJpg
	_, err = os.Stat(path)
	if err == nil {
		return path
	}

	return ""
}

func (g *Game) CreateText(str string) {
	surface, err := g.Games.application.fonts.Font14.RenderUTF8_Solid(str, g.Games.application.config.GameTitle.Color)
	if LogWarning(err) {
		return
	}
	defer surface.Free()

	texture, err := g.Games.application.renderer.CreateTextureFromSurface(surface)
	LogWarning(err)
	g.Text = texture

	w, h, err := g.Games.application.fonts.Font14.SizeUTF8(str)
	if LogWarning(err) {
		return
	}
	g.TextSrc = sdl.Rect{0, 0, int32(w), int32(h)}
}

func (g *Game) CalcPosition() (int32, int32) {
	c := g.Games.application.config.Columns
	margin := g.Games.application.config.Margin
	row := int32(math.Ceil(float64(g.Index)/float64(c)+0.1)) - 1
	col := g.Index - (row * c)
	x := (col * (g.Games.ItemSize.W + margin)) + margin
	y := int32(g.Index/c+1)*(g.Games.ItemSize.H+margin) + margin - (g.Games.ItemSize.H + margin)

	return x, y
}

func (g *Game) Destroy() {
	g.Text.Destroy()
}

func (g *Game) Reset() {
	g.CreateText(g.Title)
}

func (g *Game) Exec() {
	if *g.executed {
		return
	}
	*g.executed = true
	execute <- *g
}
