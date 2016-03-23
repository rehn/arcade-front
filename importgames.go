package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func SyncGames(c Configuration, db Database) {
	for _, emu := range c.Emulators {
		if emu.RomPath != "" {
			path, err := PathToABS(emu.RomPath)
			if LogError(err) {
				return
			}

			path = strings.TrimSuffix(path, "/")
			path = path + "/*" + emu.RomExtension
			files, err := filepath.Glob(path)
			if LogError(err) {
				return
			}
			for _, file := range files {
				name := filepath.Base(file)
				title := filepath.Base(file)
				title = strings.TrimSuffix(title, emu.RomExtension)

				n1 := strings.ToUpper(title[0:1])
				n2 := title[1:]
				title = n1 + n2
				db.InsertMissingGame(name, title, "", emu.Name, 0)
			}
		}

	}
}

func ImportGames(c Configuration, db Database, name string) {
	emu := c.Emulator(name)
	fmt.Println(emu)

	if emu != nil {
		if emu.RomPath != "" {
			path, err := PathToABS(emu.RomPath)
			if LogError(err) {
				return
			}

			path = strings.TrimSuffix(path, "/")
			path = path + "/*" + emu.RomExtension
			files, err := filepath.Glob(path)
			if LogError(err) {
				return
			}
			for _, file := range files {
				name := filepath.Base(file)
				title := filepath.Base(file)
				title = strings.TrimSuffix(title, emu.RomExtension)

				n1 := strings.ToUpper(title[0:1])
				n2 := title[1:]
				title = n1 + n2
				db.InsertGame(name, title, "", emu.Name, 0)
			}

		}
	}
}
