package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os/exec"
	"strings"
)

func NewApplication(c *Configuration, db Database) Application {

	sdl.ShowCursor(sdl.DISABLE)
	a := Application{}
	a.config = c
	a.running = new(bool)
	a.executed = new(bool)
	a.focused = new(string)
	a.db = db
	a.fonts.init()
	a.input = NewInput(c)
	a.Setup()
	return a
}

type Application struct {
	config   *Configuration
	window   *sdl.Window
	renderer *sdl.Renderer
	input    Input
	focused  *string
	games    *Games
	texture  *sdl.Texture
	fonts    Fonts
	db       Database
	running  *bool
	executed *bool
}

func (a *Application) initWindow() error {
	var err error
	var flags uint32 = sdl.WINDOW_FULLSCREEN
	if !a.config.Fullscreen {
		flags = sdl.WINDOW_RESIZABLE | sdl.WINDOW_SHOWN
	}

	w, h := a.getSize()
	a.window, err = sdl.CreateWindow(a.config.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(w), int(h), flags)
	return err
}

func (a *Application) initRenderer() error {
	var err error
	a.renderer, err = sdl.CreateRenderer(a.window, -1, sdl.RENDERER_ACCELERATED)
	return err
}
func (a *Application) loadTexture() error {
	path, err := PathToABS(a.config.Src)
	if err != nil {
		return err
	}
	surface, err := img.Load(path)
	if err != nil {
		return err
	}
	a.texture, err = a.renderer.CreateTextureFromSurface(surface)
	return err
}

func (a *Application) Run() bool {
	*a.running = true
	for *a.running {
		select {
		case g := <-execute:
			a.Execute(g)
		default:
			if a.input.Close() || a.input.Exit.Pressed() {
				a.Destroy()
				return false
			}
			if a.input.Reload.Pressed() {
				a.Destroy()
				return true
			}
			a.renderer.Clear()
			a.games.Update()
			a.renderer.Present()
			sdl.PumpEvents()
			sdl.Delay(5)

		}
	}
	return true
}
func (a *Application) getSize() (int32, int32) {
	w := a.config.WindowSize.W
	h := a.config.WindowSize.H
	return w, h
}

func (a *Application) Execute(game Game) {
	cmd := a.Command(game)
	if cmd == nil {
		return
	}
	LogMsg("Execute game: " + game.Title)
	db.UpdateExec(game.Name)
	targetY = *a.games.TargetY
	currentIndex = *a.games.CurrentIndex
	a.Destroy()
	err := cmd.Start()
	if LogWarning(err) {
		a.Setup()
		return
	}
	cmd.Wait()
	*a.running = false
	*game.executed = false

}

func (a *Application) Command(g Game) *exec.Cmd {
	emulator := a.config.Emulator(g.Type)
	if emulator == nil {
		LogMsg("Missing Emulator")
		return nil
	}
	attr := emulator.Attributes

	fl := g.Name
	if emulator.SkipExtension {
		fl = strings.TrimSuffix(fl, emulator.RomExtension)
	}
	attr = append(attr, fl)
	return exec.Command(emulator.Command, attr...)
}

func (a *Application) Destroy() {
	a.games.Destroy()
	a.texture.Destroy()
	a.renderer.Destroy()
	a.window.Destroy()
}

func (a *Application) Setup() {
	LogFatal(a.initWindow())
	LogFatal(a.initRenderer())
	LogFatal(a.loadTexture())
	a.games = NewGames(a)
}

func (a *Application) Pause() {
	a.Destroy()

}

func (a *Application) Continue() {
	targetY := *a.games.TargetY
	currentIndex := *a.games.CurrentIndex
	LogFatal(a.initRenderer())
	LogFatal(a.initWindow())

	*a.games.TargetY = targetY
	*a.games.CurrentIndex = currentIndex
}
