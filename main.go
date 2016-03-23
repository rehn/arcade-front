package main

import (
	"flag"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"runtime"
)

var (
	//app          Application
	config       Configuration
	db           Database
	targetY      int32
	currentIndex int32
	execute      chan Game
)

func main() {
	runtime.LockOSThread()

	execute = make(chan Game, 1)
	importgames := flag.String("import-games", "", "Import all games for selected emulator.")
	dropgames := flag.String("drop-games", "", "Remove all games from database for selected emulator")
	flag.Parse()
	LogSetup()

	f, err := PathToABS("assets/config.toml")
	LogFatal(err)
	err = config.LoadFromFile(f)
	LogFatal(err)

	db.Setup(config.Database)

	if *importgames != "" {
		fmt.Println("Import games for emulator: " + *importgames)
		ImportGames(config, db, *importgames)
		return
	}
	if *dropgames != "" {
		fmt.Println("Drop all games for emulator: " + *dropgames)
		db.DropGamesByEmulator(*dropgames)
		return
	}
	SyncGames(config, db)

	res := true
	go NewWebApi(config)
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
	res = StartApplication()
	for res {
		sdl.Init(sdl.INIT_EVERYTHING)
		ttf.Init()
		res = StartApplication()
		ttf.Quit()
		sdl.Quit()
	}

	LogDefer()
}

func StartApplication() bool {
	app := NewApplication(&config, db)
	if targetY > 0 && currentIndex > 0 {
		*app.games.TargetY = targetY
		*app.games.CurrentIndex = currentIndex

		targetY = 0
		currentIndex = 0
	}
	return app.Run()
}
