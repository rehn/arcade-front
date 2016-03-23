package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func NewInput(c *Configuration) Input {
	var ipt Input
	ipt.Left = Key{c.Left, time.Now()}
	ipt.Right = Key{c.Right, time.Now()}
	ipt.Up = Key{c.Up, time.Now()}
	ipt.Down = Key{c.Down, time.Now()}
	ipt.Select = Key{c.Select, time.Now()}
	ipt.Reload = Key{c.Reload, time.Now()}
	ipt.Exit = Key{c.Exit, time.Now()}
	return ipt
}

type Input struct {
	Left   Key
	Right  Key
	Up     Key
	Down   Key
	Reload Key
	Select Key
	Exit   Key
}

func (ipt *Input) Close() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return true
		}
	}
	return false
}

type Key struct {
	Name []string
	last time.Time
}

func (k *Key) Pressed() bool {
	if time.Since(k.last) < (150 * time.Millisecond) {
		return false
	}
	// Escape, Button3, Axis1+, Axis1-
	k.last = time.Now()
	for _, key := range k.Name {

		if strings.HasPrefix(key, "Button") {
			key = strings.TrimPrefix(key, "Button")
			if k.checkGamepadButton(key) {
				return true
			}
		} else if strings.HasPrefix(key, "Axis") {
			key = strings.TrimPrefix(key, "Axis")
			if k.checkGamepadAxis(key) {
				return true
			}
		} else {
			if k.checkKeyboardKey(key) {
				return true
			}
		}
	}
	return false
}

func (k *Key) checkKeyboardKey(key string) bool {
	states := sdl.GetKeyboardState()
	state := sdl.GetScancodeFromName(key)
	return states[state] == 1
}

func (k *Key) checkGamepadButton(key string) bool {
	btn, err := strconv.Atoi(key)
	if err != nil {
		log.Print(err)
		return false
	}
	for i := 0; i < sdl.NumJoysticks(); i++ {
		j := sdl.JoystickOpen(i)
		d := j.GetButton(btn)
		if d == 1 {
			return true
		}
	}
	return false
}

func (k *Key) checkGamepadAxis(key string) bool {
	axis, err := strconv.Atoi(key[:1])
	if err != nil {
		log.Print(err)
		return false
	}
	for i := 0; i < sdl.NumJoysticks(); i++ {
		j := sdl.JoystickOpen(i)
		d := j.GetAxis(axis)
		switch key[1:2] {
		case "+":
			if d > 10000 {
				return true
			}
		case "-":
			if d < -10000 {
				return true
			}
		}
	}
	return false
}
