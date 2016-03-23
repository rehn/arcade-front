package main

import (
	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Title               string
	Fullscreen          bool
	WindowSize          Size
	Database            string
	Columns             int32
	Margin              int32
	Query               string
	AnimationSpeed      int32
	HttpPort            int
	Src                 string
	FanartFallback      Rect
	GameTitle           Color
	GameTitleBackground Color
	GameTitleAjustment  Size
	Background          Color
	Error               Color
	Left                []string
	Right               []string
	Up                  []string
	Down                []string
	Select              []string
	Reload              []string
	Exit                []string
	Emulators           []Emulator
}

type Emulator struct {
	Name           string
	Command        string
	Attributes     []string
	RomPath        string
	RomExtension   string
	SkipExtension  bool
	ScreenshotPath string
	Icon           string
}

func (c *Configuration) LoadFromFile(file string) error {
	file, err := PathToABS(file)
	_, err = toml.DecodeFile(file, c)
	return err
}

func (c *Configuration) Emulator(name string) *Emulator {
	for _, e := range c.Emulators {
		if e.Name == name {
			return &e
		}
	}
	LogMsg("Can not find Emulator " + name)
	return nil
}

func (c *Configuration) Save() {

}
