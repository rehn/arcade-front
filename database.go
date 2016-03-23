package main

import (
	"github.com/mxk/go-sqlite/sqlite3"
	"strconv"
	"sync"
	"time"
)

type Database struct {
	file  string
	conn  *sqlite3.Conn
	inuse *sync.Mutex
}

func (db *Database) Close() {
	db.inuse.Unlock()
}

func (db *Database) Setup(file string) {
	db.inuse = &sync.Mutex{}
	db.inuse.Lock()
	defer db.inuse.Unlock()
	d, err := PathToABS(file)
	LogFatal(err)
	db.file = d
	db.conn, err = sqlite3.Open(d)
	LogFatal(err)
	err = db.conn.Exec(`CREATE TABLE IF NOT EXISTS items (name UNIQUE,
                                                     title TEXT,
                                                     image TEXT,
                                                     type TEXT,
                                                     last TEXT,
                                                     count INTEGER,
                                                     enabled INTEGER)`)
	LogFatal(err)
}

func (db *Database) DropGamesByEmulator(emuName string) {
	db.inuse.Lock()
	defer db.inuse.Unlock()
	err := db.conn.Exec("DELETE FROM items WHERE type='" + emuName + "'")
	LogError(err)
}

func (db *Database) InsertGame(name string, title string, image string, gameType string, enabled int) {
	db.inuse.Lock()
	defer db.inuse.Unlock()

	args := sqlite3.NamedArgs{"$name": name, "$title": title, "$image": image, "$type": gameType, "$enabled": enabled}
	err := db.conn.Exec("INSERT INTO items VALUES($name,$title,$image,$type,'',0,$enabled)", args)
	LogError(err)
}

func (db *Database) InsertMissingGame(name string, title string, image string, gameType string, enabled int) {
	db.inuse.Lock()
	defer db.inuse.Unlock()

	args := sqlite3.NamedArgs{"$name": name, "$title": title, "$image": image, "$type": gameType, "$enabled": enabled}
	db.conn.Exec("INSERT INTO items VALUES($name,$title,$image,$type,'',0,$enabled)", args)
}

func (db *Database) UpdateExec(name string) {
	db.inuse.Lock()
	defer db.inuse.Unlock()
	d := time.Now()
	err := db.conn.Exec("UPDATE items SET last='" + d.String() + "', count = count + 1 WHERE name='" + name + "'")
	LogError(err)
}

func (db *Database) UpdateEnabled(name string, value bool) {
	db.inuse.Lock()
	defer db.inuse.Unlock()

	enable := "0"
	if value {
		enable = "1"
	}
	err := db.conn.Exec("UPDATE items SET enabled='" + enable + "' WHERE name='" + name + "'")
	LogError(err)
}

func (db *Database) GetGames(query string, g *Games) {
	db.inuse.Lock()
	defer db.inuse.Unlock()

	i := int32(0)
	s, err := db.conn.Query(query)
	if LogWarning(err) {
		return
	}
	for err == nil {
		game := NewGame(g, i)
		s.Scan(&game.Name, &game.Title, &game.Image, &game.Type, &game.Last, &game.Count, &game.Enabled)
		game.Emulator = g.application.config.Emulator(game.Type)
		game.CreateText(strconv.Itoa(int(game.Index+1)) + ". " + game.Title)
		game.DrawFanart()
		g.Items = append(g.Items, game)
		i++
		err = s.Next()
	}
	s.Close()

}

func (db *Database) GetGamesData(query string) []Game {
	db.inuse.Lock()
	defer db.inuse.Unlock()
	var games []Game

	idx := 1
	s, err := db.conn.Query(query)
	if LogWarning(err) {
		return games
	}
	for err == nil {
		game := Game{}
		s.Scan(&game.Name, &game.Title, &game.Image, &game.Type, &game.Last, &game.Count, &game.Enabled)
		game.Index = int32(idx)

		games = append(games, game)
		idx++
		err = s.Next()
	}
	s.Close()
	return games

}

func (db *Database) GetGameData(name string) Game {
	db.inuse.Lock()
	defer db.inuse.Unlock()
	rows, err := db.conn.Query("SELECT * FROM items WHERE name='" + name + "' LIMIT 1")
	game := Game{}
	if err == nil {
		rows.Scan(&game.Name, &game.Title, &game.Image, &game.Type, &game.Last, &game.Count, &game.Enabled)
		return game
	}
	rows.Close()
	return game
}
