package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func NewGames(app *Application) *Games {
	var games Games
	games.application = app
	games.CurrentIndex = new(int32)
	games.TargetY = new(int32)
	w, h := app.window.GetSize()
	margin := app.config.Margin
	c := app.config.Columns
	games.Rectangle = sdl.Rect{0, 0, int32(w), int32(h)}
	games.ItemSize.W = int32((games.Rectangle.W - (margin * (c + 1))) / c)
	games.ItemSize.H = int32(float64(games.ItemSize.W) * 0.667)
	app.db.GetGames(app.config.Query, &games)
	return &games
}

type Games struct {
	Items        []Game
	application  *Application
	CurrentIndex *int32
	TargetY      *int32
	ItemSize     Size
	Rectangle    sdl.Rect
}

func (g *Games) Update() {
	g.HandleInput()
	_, h := g.application.window.GetSize()
	margin := g.application.config.Margin
	c := g.application.config.Columns
	isize := (g.ItemSize.H + margin)
	ref := int32(*g.CurrentIndex/c+1) * isize
	ref -= ((g.ItemSize.H + margin) * 2)
	ref = int32(math.Max(float64(ref), 0))
	maxRef := int32(int32(len(g.Items)-1)/c)*isize - (int32(h) - (int32(h) - (isize * 2)))
	ref = int32(math.Min(float64(ref), float64(maxRef)))

	if *g.TargetY < ref {
		*g.TargetY += g.application.config.AnimationSpeed
	}
	if *g.TargetY > ref {
		*g.TargetY -= g.application.config.AnimationSpeed
	}
	clr := g.application.config.Background
	g.application.renderer.SetDrawColor(clr.R, clr.G, clr.B, clr.A)
	g.application.renderer.FillRect(&g.Rectangle)
	if len(g.Items) == 0 {
		g.DrawNoGames()
	} else {
		for _, item := range g.Items {
			item.Update()
		}
	}
}

func (g *Games) DrawNoGames() {
	str := "Kunde inte hitta några spel i databasen."
	surface, err := g.application.fonts.Font20.RenderUTF8_Solid(str, g.application.config.Error.Color)
	if LogWarning(err) {
		return
	}
	defer surface.Free()
	w, h, err := g.application.fonts.Font20.SizeUTF8(str)
	if LogWarning(err) {
		return
	}
	screenW, _ := g.application.window.GetSize()
	src := sdl.Rect{0, 0, int32(w), int32(h)}
	dst := sdl.Rect{int32((screenW - w) / 2), 100, int32(w), int32(h)}
	texture, err := g.application.renderer.CreateTextureFromSurface(surface)
	g.application.renderer.Copy(texture, &src, &dst)
	LogWarning(err)
}
func (g *Games) HandleInput() {
	c := g.application.config.Columns
	if *g.CurrentIndex >= c && g.application.input.Up.Pressed() {
		*g.CurrentIndex -= c
		return
	}
	if *g.CurrentIndex < int32(int32(len(g.Items))-c) && g.application.input.Down.Pressed() {
		*g.CurrentIndex += c
		return
	}
	if *g.CurrentIndex > 0 && g.application.input.Left.Pressed() {
		*g.CurrentIndex -= 1
		return
	}
	if *g.CurrentIndex < int32(len(g.Items)-1) && g.application.input.Right.Pressed() {
		*g.CurrentIndex += 1
		return
	}
	if g.application.input.Select.Pressed() {
		game := g.Items[*g.CurrentIndex]
		LogMsg("Execute game (" + game.Name + ")")
		game.Exec()
	}
}

func (g *Games) Destroy() {
	for _, game := range g.Items {
		game.Destroy()
	}
}

func (g *Games) Reset() {
	for _, game := range g.Items {
		game.Reset()
	}
}
